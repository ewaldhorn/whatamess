export class Player {
  // --------------------------------------------------------------------------
  constructor() {
    this.position = { x: canvas.width / 2, y: canvas.height - 110 };
    this.velocity = { x: 0, y: 0 };
    this.width = 100;
    this.height = 100;

    // const image = new Image()
    // image.src = "./images/player.png"
    // this.image = image
  }

  // --------------------------------------------------------------------------
  draw(ctx) {
    ctx.fillStyle = "red";
    ctx.fillRect(this.position.x, this.position.y, this.width, this.height);
    // TODO: Draw an image here instead
    // ctx.drawImage(this.image,this.position.x,this.position.y);
  }
}
