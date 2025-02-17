// ----------------------------------------------------------------------------
export const clearScreen = (ctx) => {
  ctx.fillStyle = "#000000";
  ctx.fillRect(0, 0, canvas.width, canvas.height);
};

// ----------------------------------------------------------------------------
export const drawCanvasBorder = (ctx) => {
  ctx.strokeStyle = "#444444";
  ctx.lineWidth = 1;
  ctx.strokeRect(0, 0, canvas.width, canvas.height);
};
