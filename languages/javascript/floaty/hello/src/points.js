import { isRandomTrue } from "./utils";

// ----------------------------------------------------------------------------
export class Points {
  limit = 50;

  init(engine) {
    this.engine = engine;
    this.points = [];

    for (var i = 0; i < this.limit; i++) {
      this.points.push(
        new Point({
          x: 3 + Math.floor(rnd(120)),
          y: 3 + Math.floor(rnd(120)),
          engine: this.engine,
        }),
      );
    }
  }

  update() {
    this.points = this.points.filter((point) => point.y < 128);

    while (this.points.length < this.limit) {
      this.points.push(
        new Point({
          x: 3 + Math.floor(rnd(120)),
          y: 3 + Math.floor(rnd(120)),
          engine: this.engine,
        }),
      );
    }

    for (var point of this.points) {
      point.update();
    }
  }

  draw() {
    for (var point of this.points) {
      point.draw();
    }
  }
}

// ----------------------------------------------------------------------------
//                                                                        POINT
class Point {
  // --------------------------------------------------------------------------
  constructor({ x, y, engine }) {
    this.x = x;
    this.y = y;
    this.engine = engine;

    this.xv = 1 + Math.floor(Math.random() * 2);
    this.yv = 1 + Math.floor(Math.random() * 2);

    if (isRandomTrue()) {
      this.xv *= -1;
    }

    if (isRandomTrue()) {
      this.yv *= -1;
    }
  }

  update() {
    this.x += this.xv;
    this.y += this.yv;

    if (this.x < 2 || this.x > 126) {
      this.xv *= -1;
    }
    if (this.y < 2 || this.y > 126) {
      this.yv *= -1;
    }
  }

  draw() {
    pset(this.x, this.y, 12);
  }
}
