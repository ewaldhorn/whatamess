package za.co.nofuss.jasteroids;

import org.teavm.jso.JSBody;

/**
 * Bridge to the audio synth living in index.html — the Web Audio / AudioWorklet APIs aren't
 * fully exposed through TeaVM's JSO bindings, so sound stays JS-side and is driven from here,
 * mirroring the {@code foreign import app_env} bridge in the Odin original.
 */
final class Env {
    private Env() {
    }

    @JSBody(params = "id", script = "playSound(id);")
    static native void playSound(int id);

    @JSBody(params = "on", script = "setThrust(on);")
    static native void setThrust(boolean on);
}
