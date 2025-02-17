export class Bullet {
  // --------------------------------------------------------------------------
  constructor({ position, velocity }) {
    this.position = position;
    this.velocity = velocity;
    this.radius = 3;
  }

  // --------------------------------------------------------------------------
  draw(ctx) {
    ctx.beginPath();
    ctx.arc(this.position.x, this.position.y, this.radius, 0, Math.PI * 2);
    ctx.fillStyle = "red";
    ctx.fill();
    ctx.closePath();
  }

  // --------------------------------------------------------------------------
  update() {
    this.position.x += this.velocity.x;
    this.position.y += this.velocity.y;

    if (this.position.y < 3) {
      this.position.y = 3;
    }

    if (this.position.y > canvas.height - 3) {
      this.position.y = canvas.height - 3;
    }

    if (this.position.x < 3) {
      this.position.x = 3;
    }

    if (this.position.x > canvas.width - 3) {
      this.position.x = canvas.width - 3;
    }
  }

  // --------------------------------------------------------------------------
  isAlive() {
    return (
      this.position.x > 3 &&
      this.position.x < canvas.width - 3 &&
      this.position.y > 3 &&
      this.position.y < canvas.height - 3
    );
  }
}
