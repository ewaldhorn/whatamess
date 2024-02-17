import { Entity } from "./entity.js";

let canvas, context, width, height;
let isDrawing = false;

const setup = () => {
    canvas = document.getElementById("canvas");
    context = canvas.getContext("2d");

    resized();
    render();
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
    const entity = new Entity(e.x, e.y);
    entity.update(context);
}

window.addEventListener("DOMContentLoaded", setup);
window.addEventListener("resize", resized);
window.onmousedown = (e) => { isDrawing = true; };
window.onmouseup = (e) => { isDrawing = false; };
window.onmousemove = (e) => { if (isDrawing) { handleMouseMoveEvent(e); } }
// window.onclick = (e) => { handleInteractionEvent(e); };
