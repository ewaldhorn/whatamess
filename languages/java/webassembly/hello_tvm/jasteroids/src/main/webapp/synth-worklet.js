// ------------------------------------------------------------------------------------------------
// synth-worklet.js — AudioWorkletProcessor: a spacey ambient generator for Destroids.
//
// Runs on the browser's dedicated audio thread (not the game/render thread), so producing the
// background music can never stall the animation loop — this is the "safe" background-music
// pattern borrowed from OdinDOM's retrowave synth. It is pure procedural Web Audio: a slowly
// gliding 5-voice minor-9 pad, gentle LFO filter/tremolo movement, and random bell "twinkles"
// fed through a long feedback echo for a deep-space tail. No audio files, nothing to fetch.
//
// The main thread controls it with a single message: { type: 'play', value: <bool> }. When not
// playing, a master gain glides to zero so start/stop never clicks, and process() wraps
// everything in try/catch to emit silence instead of throwing.
// ------------------------------------------------------------------------------------------------

// Lush, wide minor-9 / add9 voicings (MIDI notes, low→high). Upper voices overlap between chords
// so the per-voice portamento glide slides smoothly from one to the next. Progression is slow.
const CHORDS = [
  [45, 52, 55, 59, 64], // Am add9   (A2 E3 G3 B3 E4)
  [41, 48, 52, 57, 60], // Fmaj9     (F2 C3 E3 A3 C4)
  [48, 55, 59, 62, 64], // Cmaj9     (C3 G3 B3 D4 E4)
  [43, 50, 57, 59, 64], // G6/9      (G2 D3 A3 B3 E4)
];

const NVOICES = 5;
const MAXBELL = 8;
const TWO_PI = Math.PI * 2;

function midiToFreq(m) {
  return 440.0 * Math.pow(2.0, (m - 69.0) / 12.0);
}

function softClip(x) {
  return x / (1.0 + Math.abs(x));
}

class Ambient {
  constructor(sr) {
    this.sr = sr;

    // Master play envelope: glides toward playTarget (0 or 1) so toggling never clicks.
    this.playGain = 0.0;
    this.playTarget = 0.0;

    // Slow chord sequencer (~7.5s per chord).
    this.chordDur = Math.floor(sr * 7.5);
    this.chordTimer = 0;
    this.chordIdx = 0;

    // Pad voices — each glides from its current frequency toward the chord target.
    this.glide = 5.0 / sr; // portamento coefficient per sample
    this.phase = new Float32Array(NVOICES);
    this.freq = new Float32Array(NVOICES);
    this.tfreq = new Float32Array(NVOICES);
    for (let v = 0; v < NVOICES; v++) {
      const f = midiToFreq(CHORDS[0][v]);
      this.freq[v] = f;
      this.tfreq[v] = f;
    }

    this.lp = 0.0;          // 1-pole lowpass state for the pad
    this.lfo = 0.0;         // very slow movement LFO
    this.lfoInc = (TWO_PI * 0.045) / sr; // ~0.045 Hz

    // Bell twinkles — a fixed pool, no per-sample allocation.
    this.bellTimer = 0;
    this.bellEvery = Math.floor(sr * 0.7);
    this.bPhase = new Float32Array(MAXBELL);
    this.bFreq = new Float32Array(MAXBELL);
    this.bEnv = new Float32Array(MAXBELL);
    this.bDec = new Float32Array(MAXBELL);

    // Feedback delay for the spacey echo tail (~0.42s).
    this.delay = new Float32Array(Math.floor(sr * 0.42));
    this.dw = 0;

    this.rng = 0x1a2b3c4d;
  }

  rand() {
    let x = this.rng;
    x ^= x << 13; x ^= x >>> 17; x ^= x << 5;
    this.rng = x >>> 0;
    return this.rng / 4294967296;
  }

  setChord(idx) {
    for (let v = 0; v < NVOICES; v++) {
      this.tfreq[v] = midiToFreq(CHORDS[idx][v]);
    }
  }

  triggerBell() {
    const notes = CHORDS[this.chordIdx];
    let note = notes[2 + ((this.rand() * 3) | 0)]; // an upper chord tone
    if (this.rand() < 0.45) note += 12;            // sometimes an octave up
    for (let k = 0; k < MAXBELL; k++) {
      if (this.bEnv[k] <= 0.001) {
        this.bPhase[k] = 0.0;
        this.bFreq[k] = midiToFreq(note);
        this.bEnv[k] = 1.0;
        this.bDec[k] = Math.exp(-1.0 / (this.sr * 1.3)); // ~1.3s decay
        return;
      }
    }
  }

  nextSample() {
    // -- Chord sequencer --
    if (this.chordTimer <= 0) {
      this.setChord(this.chordIdx);
      this.chordIdx = (this.chordIdx + 1) % CHORDS.length;
      this.chordTimer = this.chordDur;
    }
    this.chordTimer--;

    // -- Pad: 5 gliding sine voices --
    let pad = 0.0;
    for (let v = 0; v < NVOICES; v++) {
      this.freq[v] += (this.tfreq[v] - this.freq[v]) * this.glide;
      this.phase[v] += this.freq[v] / this.sr;
      if (this.phase[v] >= 1.0) this.phase[v] -= 1.0;
      pad += Math.sin(TWO_PI * this.phase[v]);
    }
    pad *= 0.13;

    // Slow movement: gentle lowpass whose cutoff and level breathe with the LFO.
    this.lfo += this.lfoInc;
    if (this.lfo > 1e6) this.lfo -= 1e6;
    const s = Math.sin(this.lfo);
    const cutoff = 0.06 + 0.03 * s;
    this.lp += (pad - this.lp) * cutoff;
    const padOut = this.lp * (0.8 + 0.2 * Math.sin(this.lfo * 0.7));

    // -- Bell twinkles --
    this.bellTimer--;
    if (this.bellTimer <= 0) {
      this.bellTimer = this.bellEvery + ((this.rand() * this.bellEvery) | 0);
      this.triggerBell();
    }
    let bell = 0.0;
    for (let k = 0; k < MAXBELL; k++) {
      if (this.bEnv[k] > 0.001) {
        this.bPhase[k] += this.bFreq[k] / this.sr;
        if (this.bPhase[k] >= 1.0) this.bPhase[k] -= 1.0;
        bell += Math.sin(TWO_PI * this.bPhase[k]) * this.bEnv[k];
        this.bEnv[k] *= this.bDec[k];
      }
    }
    bell *= 0.09;

    // -- Feedback echo (bells only, for a long trailing shimmer) --
    const echoed = this.delay[this.dw];
    this.delay[this.dw] = bell + echoed * 0.5;
    this.dw = (this.dw + 1) % this.delay.length;

    // -- Mix + smooth play envelope --
    let mix = padOut * 0.9 + bell * 0.5 + echoed * 0.8;
    this.playGain += (this.playTarget - this.playGain) * 0.00002; // ~1s glide
    mix *= this.playGain;

    return softClip(mix * 0.9);
  }
}

class SpaceyAmbientProcessor extends AudioWorkletProcessor {
  constructor() {
    super();
    this.eng = new Ambient(sampleRate);
    this.port.onmessage = (e) => {
      const msg = e.data;
      if (!msg || typeof msg !== "object") return;
      if (msg.type === "play") {
        this.eng.playTarget = msg.value ? 1.0 : 0.0;
      }
    };
  }

  process(inputs, outputs) {
    const output = outputs[0];
    if (!output || output.length === 0) return true;
    const ch0 = output[0];
    if (!ch0) return true;
    const n = ch0.length;

    // Fully idle: nothing playing and the envelope has faded out — emit silence cheaply.
    if (this.eng.playTarget === 0.0 && this.eng.playGain < 1e-5) {
      for (let c = 0; c < output.length; c++) output[c].fill(0);
      return true;
    }

    try {
      for (let i = 0; i < n; i++) {
        const v = this.eng.nextSample();
        for (let c = 0; c < output.length; c++) output[c][i] = v;
      }
    } catch (err) {
      for (let c = 0; c < output.length; c++) output[c].fill(0);
    }
    return true;
  }
}

registerProcessor("spacey-ambient", SpaceyAmbientProcessor);
