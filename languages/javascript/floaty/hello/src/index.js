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

let mustClearScreen = true;

var stars = new Stars();

async function init() {
  stars.init(this);

  await this.load();

  mode = "start";
}

function update() {
  stars.update();
}

function draw() {
  if (mustClearScreen) {
    cls();
  }

  stars.draw();

  // Clear the screen
  if (btnp("c")) {
    cls();
  }

  // Toggle overwrite or not
  if (btnp("o")) {
    mustClearScreen = !mustClearScreen;
  }
}

// ----------------------------------------------------------------------------
//                                                                  ENTRY POINT
engine.start({
  sprites,
  sounds,
  init,
  update,
  draw,
  target: document.querySelector(".game"),
});
