// ------------------------------------------------------------------------------------------------
export const drawBalloon = (ctx: CanvasRenderingContext2D) => {
  // basket
  ctx.fillStyle = "#DB504A";
  ctx.fillRect(-30, -40, 60, 10);
  ctx.fillStyle = "#EA9E8D";
  ctx.fillRect(-30, -30, 60, 30);

  // cables
  ctx.strokeStyle = "#D62828";
  ctx.lineWidth = 2;
  ctx.beginPath();
  ctx.moveTo(-24, -40);
  ctx.lineTo(-24, -60);
  ctx.moveTo(24, -40);
  ctx.lineTo(24, -60);
  ctx.stroke();

  // balloon
  ctx.fillStyle = "D62828";
  ctx.beginPath();
  ctx.moveTo(-30, -60);
  ctx.quadraticCurveTo(-80, -120, -80, -160);
  ctx.arc(0, -160, 80, Math.PI, 0, false);
  ctx.quadraticCurveTo(80, -120, 30, -60);
  ctx.closePath();
  ctx.fill();
};

// ------------------------------------------------------------------------------------------------
export const drawTree = (ctx: CanvasRenderingContext2D) => {};

// ------------------------------------------------------------------------------------------------
/**
 * Draws a circle on the canvas context.
 * @param ctx - The 2D rendering context of the canvas.
 * @param cx - The x-coordinate of the circle's center.
 * @param cy - The y-coordinate of the circle's center.
 * @param radius - The radius of the circle.
 */
const drawCircle = (ctx: CanvasRenderingContext2D, cx: number, cy: number, radius: number) => {
  ctx.beginPath();
  ctx.arc(cx, cy, radius, 0, 2 * Math.PI);
  ctx.fill();
};
