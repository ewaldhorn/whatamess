export class Entity {
    constructor(x, y) {
        this.x = x;
        this.y = y;
        this.speedX = Math.random() * 4 - 2;
        this.speedY = Math.random() * 4 - 2;
        this.maxSize = Math.random() * 7 + 5;
        this.size = Math.random() * 1 + 2;
    }

    /**
     * 
     * @param {CanvasRenderingContext2D} context 
     */
    update(context) {
        this.x += this.speedX;
        this.y += this.speedY;
        this.size += 0.1;
        if (this.size < this.maxSize) {
            context.beginPath();
            context.arc(this.x, this.y, this.size, 0, 2 * Math.PI);
            context.fillStyle = 'hsl(140,100%,50%)';
            context.fill();
            context.stroke();
        }
    }
}