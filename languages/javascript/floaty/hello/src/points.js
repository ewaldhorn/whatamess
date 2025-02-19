import { isRandomTrue } from "./utils";

// ----------------------------------------------------------------------------
export class Points {
  limit = 22;

  // --------------------------------------------------------------------------
  init() {
    this.points = [];

    for (var i = 0; i < this.limit; i++) {
      this.points.push(
        new Point({
          x: 3 + Math.floor(rnd(120)),
          y: 3 + Math.floor(rnd(120)),
        }),
      );
    }
  }

  // --------------------------------------------------------------------------
  update() {
    for (var point of this.points) {
      point.update();
    }
  }

  // --------------------------------------------------------------------------
  draw() {
    for (let i = 0; i < this.points.length - 1; i++) {
      line(
        this.points[i].x,
        this.points[i].y,
        this.points[i + 1].x,
        this.points[i + 1].y,
        i % 14,
      );
    }

    line(
      this.points[0].x,
      this.points[0].y,
      this.points[this.points.length - 1].x,
      this.points[this.points.length - 1].y,
      15,
    );
  }
}

// ----------------------------------------------------------------------------
//                                                                        POINT
class Point {
  // --------------------------------------------------------------------------
  constructor({ x, y }) {
    this.x = x;
    this.y = y;

    this.xv = 1 + Math.floor(Math.random() * 2);
    this.yv = 1 + Math.floor(Math.random() * 2);

    if (isRandomTrue()) {
      this.xv *= -1;
    }

    if (isRandomTrue()) {
      this.yv *= -1;
    }
  }

  // --------------------------------------------------------------------------
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
}
