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

    const boundary = 3;
    this.position.x = Math.max(
      boundary,
      Math.min(this.position.x, canvas.width - boundary),
    );
    this.position.y = Math.max(
      boundary,
      Math.min(this.position.y, canvas.height - boundary),
    );
  }

  // --------------------------------------------------------------------------
  isAlive() {
    const { x, y } = this.position;
    const boundary = 3;
    return (
      x > boundary &&
      x < canvas.width - boundary &&
      y > boundary &&
      y < canvas.height - boundary
    );
  }
}
