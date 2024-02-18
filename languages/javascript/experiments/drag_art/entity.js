export class Entity {
    constructor(x, y, context) {
        this.x = x;
        this.y = y;
        this.speedX = Math.random() * 4 - 2;
        this.speedY = Math.random() * 4 - 2;
        this.maxSize = Math.random() * 7 + 5;
        this.size = Math.random() * 1 + 2;
        this.vs = Math.random() * 0.2 + 0.05;
        this.angleX = Math.random() * 6.2;
        this.vaX = Math.random() * 0.6 - 0.3;
        this.angleY = Math.random() * 6.2;
        this.vaY = Math.random() * 0.6 - 0.3;
        this.lightness = 10;
        this.context = context;
    }

    /**
     * 
     * @param {CanvasRenderingContext2D} context 
     */
    update() {
        this.x += this.speedX + Math.sin(this.angleX);
        this.y += this.speedY + Math.sin(this.angleY);
        this.size += this.vs;
        this.angleX += this.vaX;
        this.angleY += this.vaY;

        if (this.lightness < 70) { this.lightness += 0.25; };

        if (this.size < this.maxSize) {
            this.context.beginPath();
            this.context.arc(this.x, this.y, this.size, 0, 2 * Math.PI);
            this.context.fillStyle = 'hsl(190,100%,' + this.lightness + '%)';
            this.context.fill();
            this.context.stroke();

            requestAnimationFrame(this.update.bind(this));
        }
    }
}