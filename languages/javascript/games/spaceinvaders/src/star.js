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

    const boundaryX = 4;
    const boundaryY = 3;
    const resetZoneY = canvas.height - boundaryY;

    this.position.x = Math.max(
      boundaryX,
      Math.min(this.position.x, canvas.width - 10),
    );
    this.position.y = Math.max(boundaryY, this.position.y);

    if (this.position.y > resetZoneY) {
      this.position.x = 50 + Math.floor(Math.random() * (canvas.width - 80));
      this.position.y = boundaryY;
      this.velocity.y = 1 + Math.floor(Math.random() * 3);
      this.colour = this.getRandomColour();
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
