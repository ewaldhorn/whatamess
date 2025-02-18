export class Star {
  // --------------------------------------------------------------------------
  constructor({ position, velocity, colour }) {
    this.position = position;
    this.velocity = velocity;
    this.radius = 1;
    this.colour = colour;
  }

  // --------------------------------------------------------------------------
  draw(ctx) {
    ctx.beginPath();
    ctx.arc(this.position.x, this.position.y, this.radius, 0, Math.PI * 2);
    ctx.fillStyle = this.colour;
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
      this.position.y = 3;
      this.position.x = 50 + Math.floor(Math.random() * (canvas.width - 80));
      this.velocity.y = 1 + Math.floor(Math.random() * 3);
      this.colour = this.getRandomColour();
    }

    if (this.position.x < 4) {
      this.position.x = 4;
    }

    if (this.position.x > canvas.width - 10) {
      this.position.x = canvas.width - 10;
    }
  }

  // --------------------------------------------------------------------------
  getRandomColour() {
    let tmp = Math.random() * 100;

    if (tmp > 90) {
      return "#222222";
    }
    if (tmp > 80) {
      return "#333333";
    }
    if (tmp > 70) {
      return "#444444";
    }
    if (tmp > 60) {
      return "#555555";
    }
    if (tmp > 50) {
      return "#225522";
    }
    if (tmp > 40) {
      return "#222255";
    }
    if (tmp > 30) {
      return "#552222";
    }

    if (tmp > 20) {
      return "#551111";
    }

    return "#aaaaaa";
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
