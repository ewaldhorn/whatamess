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
    this.lastFired = 0; // controls how often an enemy can shoot
    this.health = 20;
    this.maxHealth = 20;
    this.colour = "orange";
  }

  // --------------------------------------------------------------------------
  /**
   * @param {CanvasRenderingContext2D=required} ctx - 2D rendering context
   */
  draw(ctx) {
    ctx.fillStyle = this.colour;
    ctx.fillRect(this.position.x, this.position.y, this.width, this.height);

    // health
    ctx.fillStyle = "green";
    ctx.fillRect(
      this.position.x + this.width / 2 - this.health / 2,
      this.position.y + this.height + 5,
      this.health,
      5,
    );
  }

  // --------------------------------------------------------------------------
  update() {
    if (this.velocity.x == 0) {
      this.velocity.x = 3;
    } else if (this.velocity.x > 9) {
      this.reverseHorizontalDirection();
    }

    this.position.x += this.velocity.x;

    if (this.position.x > canvas.width - this.width - 1) {
      this.position.x = canvas.width - this.width - 1;
      this.velocity.x = Math.floor(-0.5 * this.velocity.x);
      this.gainHealth();
    }

    if (this.position.x < 1) {
      this.position.x = 1;
      this.velocity.x = Math.floor(-0.75 * this.velocity.x);
      this.gainHealth();
    }

    if (this.lastFired >= 200) {
      this.reverseHorizontalDirection();
      this.velocity.x = Math.floor(this.velocity.x * 1.2);
    }

    this.lastFired += 1;
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
  gainHealth() {
    this.health = Math.min(this.health + 3, this.maxHealth);
  }

  // --------------------------------------------------------------------------
  canShoot() {
    return this.lastFired > 150;
  }

  // --------------------------------------------------------------------------
  isAlive() {
    return this.health > 0;
  }

  // --------------------------------------------------------------------------
  getBulletPosition() {
    this.lastFired = 0;
    return {
      x: this.position.x + this.width / 2,
      y: this.position.y + this.height,
    };
  }
}
