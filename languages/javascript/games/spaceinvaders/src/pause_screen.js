// ----------------------------------------------------------------------------
/**
 * @param {CanvasRenderingContext2D=required} ctx - 2D rendering context
 */
export const showPauseScreen = (ctx) => {
  ctx.font = "36px Arial";
  ctx.textAlign = "left";
  ctx.textBaseline = "middle";

  const text = "- - -  (P)AUSED  - - -";
  const metrics = ctx.measureText(text);
  const textWidth = metrics.width;

  ctx.fillText(text, (canvas.width - textWidth) / 2, canvas.height / 2);
};
