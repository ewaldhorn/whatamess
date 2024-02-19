"use strict";

let canvas, context, width, height, effect;

const setup = () => {
    canvas = document.getElementById("backgroundCanvas");
    context = canvas.getContext("2d");
    resized(context);
};

const resized = (context) => {
    width = canvas.width = window.innerWidth;
    height = canvas.height = window.innerHeight;
    drawBackground(context);
    effect = new Effect(width, height);
    animate();
}

/**
 * 
 * @param {CanvasRenderingContext2D} context 
 */
const drawBackground = (context) => {
    context.fillStyle = "blue";
    context.fillRect(10, 10, 10, 10);
    context.fillRect(width - 20, 10, 10, 10);
    context.fillRect(10, height - 20, 10, 10);
    context.fillRect(width - 20, height - 20, 10, 10);
}

const animate = () => {
    context.fillStyle = "rgba(0, 0, 0, 0.05)";
    context.fillRect(0, 0, width, height);

    context.font = effect.fontSize + "px monospace";
    effect.symbols.forEach(symbol => symbol.draw(context));

    requestAnimationFrame(animate);
}

class Symbol {
    constructor(x, y, fontSize, canvasHeight) {
        this.characters = '0123456789ABCDEF';
        this.x = x;
        this.y = y;
        this.fontSize = fontSize;
        this.text = '';
        this.canvasHeight = canvasHeight;
    }

    /**
     * 
     * @param {CanvasRenderingContext2D} context 
     */
    draw(context) {
        this.text = this.characters.charAt(Math.floor(Math.random() * this.characters.length));
        context.fillStyle = "#093009";
        context.fillText(this.text, this.x * this.fontSize, this.y * this.fontSize);
        if (this.y * this.fontSize > this.canvasHeight && Math.random() > 0.97) {
            this.y = 0;
        } else {
            this.y += 1;
        }
    }
}

class Effect {
    constructor(width, height) {
        this.canvasWidth = width;
        this.canvasHeight = height;
        this.fontSize = 25;
        this.columns = this.canvasWidth / this.fontSize;
        this.symbols = [];

        this.#initialize();
    }
    #initialize() {
        for (let i = 0; i < this.columns; i++) {
            this.symbols[i] = new Symbol(i, 0, this.fontSize, this.canvasHeight);
        }
    }
}

window.addEventListener("DOMContentLoaded", setup);
window.addEventListener("resize", resized(context));