export class Player {
  // --------------------------------------------------------------------------
  constructor() {
    this.width = 50;
    this.height = 50;
    this.position = {
      x: canvas.width / 2 - this.width / 2,
      y: canvas.height - 1,
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
  draw(ctx) {
    ctx.fillStyle = "red";

    ctx.beginPath();
    ctx.moveTo(this.position.x, this.position.y);
    ctx.lineTo(this.position.x + this.width, this.position.y); // left bottom to right
    ctx.lineTo(this.position.x + this.width / 2, this.position.y - this.height); // to middle
    ctx.fill();

    // TODO: Consider drawing an image here instead
    // if (this.image) {
    // ctx.drawImage(this.image,this.position.x,this.position.y);
    // }
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
}
