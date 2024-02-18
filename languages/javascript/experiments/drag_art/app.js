import { Entity } from "./entity.js";

let rng, canvas, context, width, height;
let isDrawing = false;

const setup = () => {
    rng = document.getElementById("rng");
    canvas = document.getElementById("canvas");
    context = canvas.getContext("2d");

    setupRNG();
    resized();
    render();
}

const setupRNG = () => {
    const choices = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

    let tmp = "";

    for (let i = 0; i < 8; i++) {
        const pos = Math.floor(1 + Math.random() * choices.length);
        tmp += choices.charAt(pos);
    }

    rng.innerText = tmp;
}

const resized = () => {
    width = canvas.width = window.innerWidth;
    height = canvas.height = window.innerHeight;
}

const render = () => {
    requestAnimationFrame(render);
}

/**
 * 
 * @param {MouseEvent} e 
 */
const handleMouseMoveEvent = (e) => {
    for (let i = 0; i < 3; i++) {
        const entity = new Entity(e.x, e.y, context);
        entity.update();
    }
}

window.addEventListener("DOMContentLoaded", setup);
window.addEventListener("resize", resized);
window.onmousedown = (e) => { isDrawing = true; };
window.onmouseup = (e) => { isDrawing = false; };
window.onmousemove = (e) => { if (isDrawing) { handleMouseMoveEvent(e); } }
// window.onclick = (e) => { handleInteractionEvent(e); };
