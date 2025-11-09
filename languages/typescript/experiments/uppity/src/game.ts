import globals from "./globals.ts";

// ------------------------------------------------------------------------------------------------
let ctx: CanvasRenderingContext2D | null;

// ------------------------------------------------------------------------------------------------
const setGameInformation = () => {
  document.title = `${globals.GAME_NAME} v${globals.VERSION}`;
};

// ------------------------------------------------------------------------------------------------
const initCanvas = () => {
  const canvas = document.getElementById("canvas") as HTMLCanvasElement;
  ctx = canvas.getContext("2d");
};

// ------------------------------------------------------------------------------------------------
export const initGame = () => {
  setGameInformation();
  initCanvas();

  if (ctx) {
  } else {
    alert("Could not load 2D rendering context - unable to continue");
  }
};
