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
        this.radius = 50;
        this.speedX = 1;
        this.speedY = 1;
    }

    update() {
        this.x += this.speedX;
        this.y += this.speedY;
    }

    draw(context) {
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

    init() {}
    update() {}
    draw(context) {}
}

function animate() {

}