// ----------------------------------------------------------------------------
export const clearScreen = (ctx) => {
  ctx.fillStyle = "#000000";
  ctx.fillRect(0, 0, canvas.width, canvas.height);
};

// ----------------------------------------------------------------------------
export const clearGameArea = (ctx) => {
  ctx.fillStyle = "#000010";
  ctx.fillRect(1, 1, canvas.width - 2, canvas.height - 2);
};

// ----------------------------------------------------------------------------
export const drawCanvasBorder = (ctx) => {
  ctx.strokeStyle = "#444444";
  ctx.lineWidth = 1;
  ctx.strokeRect(0, 0, canvas.width, canvas.height);
};
