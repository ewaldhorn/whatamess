import { Engine } from "https://floaty.dev/engine-v1.js";

var engine = new Engine();
engine.expose();

var sprites = {
  // this is where your sprites go
};

var sounds = {
  // this is where your sounds go
};

class Stars {
  limit = 50;

  init(engine) {
    this.engine = engine;
    this.stars = [];

    for (var i = 0; i < this.limit; i++) {
      this.stars.push(
        new Star({
          x: Math.floor(rnd(128)),
          y: Math.floor(rnd(128)),
          type: Math.floor(rnd(3)),
          engine: this.engine,
        }),
      );
    }
  }

  update() {
    this.stars = this.stars.filter((star) => star.y < 128);

    while (this.stars.length < this.limit) {
      this.stars.push(
        new Star({
          x: Math.floor(rnd(128)),
          y: 0,
          type: Math.floor(rnd(3)),
          engine: this.engine,
        }),
      );
    }

    for (var star of this.stars) {
      star.update();
    }
  }

  draw() {
    for (var star of this.stars) {
      star.draw();
    }
  }
}
class Star {
  x = 0;
  y = 0;
  type = 0;

  constructor({ x, y, type, engine }) {
    this.x = x;
    this.y = y;
    this.type = type;
    this.engine = engine;
  }

  update() {
    if (this.type === 0) {
      this.y += 1.5;
    }

    if (this.type === 1) {
      this.y += 1.0;
    }

    if (this.type === 2) {
      this.y += 0.5;
    }
  }

  draw() {
    if (this.type === 0) {
      pset(this.x, this.y, 6);
    }

    if (this.type === 1) {
      pset(this.x, this.y, 13);
    }

    if (this.type === 2) {
      pset(this.x, this.y, 1);
    }
  }
}

var mode = "load";
var restartTimer = null;
var restartDelay = 2500;
var waveTimer = null;
var waveDelay = 2500;
var winTimer = null;
var winDelay = 2500;

var stars = new Stars();

async function init() {
  stars.init(this);

  await this.load();

  mode = "start";
}

function update() {
  stars.update();

  if (mode === "start" && btnp()) {
    mode = "wave";
  } else if (mode === "wave" && !waveTimer) {
    waveTimer = setTimeout(() => {
      waveTimer = null;
      mode = "game";
    }, waveDelay);
  } else if (mode === "over" && !restartTimer) {
    restartTimer = setTimeout(() => {
      restartTimer = null;
      mode = "start";
    }, restartDelay);
  } else if (mode === "win" && !winTimer) {
    winTimer = setTimeout(() => {
      winTimer = null;
      mode = "start";
    }, winDelay);
  } else if (mode === "game") {
  }
}

function draw() {
  cls();

  stars.draw();

  if (mode === "load") {
    text("Loading...", 45, 56);
  } else if (mode === "start") {
    text("Press any key to start", 15, 56);
    text("Dev by assertchris", 15, 96, 5);
    text("Font by kenney", 15, 104, 5);
  } else if (mode === "wave") {
    text(`Wave`, 45, 56);
  } else if (mode === "over") {
    text("Game over", 45, 56);
  } else if (mode === "win") {
    text("You win!", 45, 56);
  } else if (mode === "game") {
  }

  // DEBUG
  if (btnp("i")) {
    mode = "win";
  }

  // DEBUG
  if (btnp("o")) {
    mode = "over";
  }
}

engine.start({
  sprites,
  sounds,
  init,
  update,
  draw,
  target: document.querySelector(".game"),
});
