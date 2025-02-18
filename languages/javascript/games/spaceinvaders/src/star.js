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
    const tmp = Math.random() * 100;

    const colours = [
      { range: 90, colour: "#222222" },
      { range: 80, colour: "#333333" },
      { range: 70, colour: "#444444" },
      { range: 60, colour: "#555555" },
      { range: 50, colour: "#228822" },
      { range: 40, colour: "#222288" },
      { range: 30, colour: "#882222" },
      { range: 20, colour: "#881111" },
      { range: 10, colour: "#333388" },
      { range: 0, colour: "#999999" },
    ];

    const selectedColour = colours.find((c) => tmp >= c.range);
    return selectedColour ? selectedColour.colour : "#777777";
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
