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
        this.radius = (Math.random() * 20) + 10;
        this.x = Math.random() * this.effect.width;
        this.y = -this.radius;
        this.speedX = Math.random() - 0.5;
        this.speedY = Math.random() * 1.5 + 0.05;
        this.gravity = Math.random() * 0.005 + 0.05;
        this.vy = 0;
    }

    update() {
        if (this.x < this.radius || this.x > (canvas.width-this.radius)) {
            this.speedX *= -1;
        }
        if (this.y > (canvas.height + this.radius)) {
            this.y = -this.radius;
            this.vy = 0;
        }

        this.vy += this.gravity;
        this.x += this.speedX;
        this.y += this.speedY + this.vy;
    }

    draw(context) {
        context.fillStyle = "green";
        context.beginPath();
        context.arc(this.x, this.y, this.radius, 0, Math.PI * 2.0);
        context.fill();
    }

    reset() {
        this.x = this.effect.width * 0.5;
        this.y = this.effect.height * 0.5;
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

    reset(newWidth, newHeight) {
        this.width = newWidth;
        this.height = newHeight;
        this.metaBallsArray.forEach(m => m.reset());
    }
}

const effect = new MetaBallsEffect(canvas.width, canvas.height);
effect.init(60);

function animate() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    effect.update();
    effect.draw(ctx);
    requestAnimationFrame(animate);
}

animate();

window.addEventListener("resize", function(){
    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;
    effect.reset(canvas.width, canvas.height);
})