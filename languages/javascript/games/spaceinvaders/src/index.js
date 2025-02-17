// ----------------------------------------------------------------------------
const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");
canvas.width = innerWidth - 50;
canvas.height = innerHeight - 100;

ctx.fillStyle = "#000000";
ctx.fillRect(0, 0, canvas.width, canvas.height);
