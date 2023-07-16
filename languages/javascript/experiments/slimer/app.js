"use strict";

// get a handle to the main canvas and 2d context
const canvas = document.getElementById("mainCanvas");
const ctx = canvas.getContext("2d");

// fill the canvas
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

class Ball {
    constructor(effect) {
        this.effect = effect;
        this.x = this.effect.width * 0.5;
        this.y = this.effect.height * 0.5;
        this.radius = (Math.random() * 40) + 10;
        this.speedX = Math.random() - 0.5;
        this.speedY = Math.random() - 0.5;
    }

    update() {
        if (this.x < this.radius || this.x > (canvas.width-this.radius)) {
            this.speedX *= -1;
        }
        if (this.y < this.radius || this.y > (canvas.height-this.radius)) {
            this.speedY *= -1;
        }

        this.x += this.speedX;
        this.y += this.speedY;
    }

    draw(context) {
        context.fillStyle = "white";
        context.beginPath();
        context.arc(this.x, this.y, this.radius, 0, Math.PI * 2.0);
        context.fill();
    }
}

class MetaBallsEffect {
    constructor(width, height) {
        this.width = width;
        this.height = height;
        this.metaBallsArray = [];
    }

    init(numberOfBalls) {
        for (let i = 0; i < numberOfBalls; i++) {
            this.metaBallsArray.push(new Ball(this));
        }
    }

    update() {
        this.metaBallsArray.forEach(m => m.update());
    }

    draw(context) {
        this.metaBallsArray.forEach(m => m.draw(context));
    }
}

const effect = new MetaBallsEffect(canvas.width, canvas.height);
effect.init(20);

function animate() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    effect.update();
    effect.draw(ctx);
    requestAnimationFrame(animate);
}

animate();