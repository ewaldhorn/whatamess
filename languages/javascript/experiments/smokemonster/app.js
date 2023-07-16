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
        this.radius = (Math.random() * 20) + 30;
        this.speedX = Math.random() - 0.5;
        this.speedY = Math.random() - 0.5;
        this.angle = 0;
        this.va = Math.random() * 0.1 - 0.05;
        this.range = Math.random() * 5;
    }

    update() {
        if (this.x < this.radius || this.x > (canvas.width-this.radius)) {
            this.speedX *= -1;
        }
        if (this.y < this.radius || this.y > (canvas.height-this.radius)) {
            this.speedY *= -1;
        }

        if (this.angle > 1500) {
            this.angle = 0;
        }

        this.angle += this.va;
        this.x += this.speedX * Math.cos(this.angle); 
        this.y += this.speedY * Math.sin(this.angle);
    }

    draw(context) {
        context.fillStyle = "grey";
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
effect.init(20);

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