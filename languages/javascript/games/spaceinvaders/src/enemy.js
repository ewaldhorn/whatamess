export class Enemy {
  // --------------------------------------------------------------------------
  /**
   * @typedef {Object} Point2D Point coordinate (x,y)
   * @property {number} x
   * @property {number} y
   */

  /**
   * Enemy class represents the computer controlled enemy
   * @class
   * @property {number} width Width of player in pixels
   * @property {number} height Height of player in pixels
   * @property {Point2D} position Position coordinates
   * @property {Point2D} velocity Velocity vector
   */
  constructor() {
    this.width = 50;
    this.height = 50;
    this.position = {
      x: canvas.width / 2 - this.width / 2,
      y: 2,
    };
    this.velocity = { x: 0, y: 0 };

    // const image = new Image()
    // image.src = "./images/player.png"
    // image.onload = () => {
    // this.image = image
    // this.width = image.width * 0.5; //scale down by 50% for example
    // this.height = image.height;
    // }
  }

  // --------------------------------------------------------------------------
  /**
   * @param {CanvasRenderingContext2D=required} ctx - 2D rendering context
   */
  draw(ctx) {
    ctx.fillStyle = "orange";
    ctx.fillRect(this.position.x, this.position.y, this.width, this.height);
  }

  // --------------------------------------------------------------------------
  update() {
    this.position.x += this.velocity.x;

    if (this.position.x > canvas.width - this.width - 1) {
      this.position.x = canvas.width - this.width - 1;
      this.velocity.x = Math.floor(-0.5 * this.velocity.x);
    }

    if (this.position.x < 1) {
      this.position.x = 1;
      this.velocity.x = Math.floor(-0.5 * this.velocity.x);
    }
  }

  // --------------------------------------------------------------------------
  goLeft() {
    this.velocity.x -= 4;
  }

  // --------------------------------------------------------------------------
  goRight() {
    this.velocity.x += 4;
  }

  // --------------------------------------------------------------------------
  reverseHorizontalDirection() {
    this.velocity.x *= -1;
  }

  // --------------------------------------------------------------------------
  getBulletPosition() {
    return {
      x: this.position.x + this.width / 2,
      y: this.position.y - this.height,
    };
  }
}
