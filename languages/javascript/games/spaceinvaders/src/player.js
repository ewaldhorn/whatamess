export class Player {
  // --------------------------------------------------------------------------
  constructor() {
    this.position = { x: canvas.width / 2 - 50, y: canvas.height - 110 };
    this.velocity = { x: 0, y: 0 };
    this.width = 100;
    this.height = 100;

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
    ctx.fillRect(this.position.x, this.position.y, this.width, this.height);
    // TODO: Draw an image here instead
    // if (this.image) {
    // ctx.drawImage(this.image,this.position.x,this.position.y);
    // }
  }

  // --------------------------------------------------------------------------
  update() {
    if (
      this.position.x + this.velocity.x < canvas.width - 100 &&
      this.position.x + this.velocity.x > 1
    ) {
      this.position.x += this.velocity.x;
    } else {
      this.velocity.x = Math.floor(-0.5 * this.velocity.x);
    }
  }

  // --------------------------------------------------------------------------
  goLeft() {
    this.velocity.x -= 1;
  }

  // --------------------------------------------------------------------------
  goRight() {
    this.velocity.x += 1;
  }
}
