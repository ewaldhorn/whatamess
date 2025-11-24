import { drawBalloon, drawTree } from "./drawing.ts";
import globals from "./globals.ts";
import { Tree } from "./trees.ts";

// ------------------------------------------------------------------------------------------------
let ctx: CanvasRenderingContext2D | null;
let trees: Array<Tree>;

// ------------------------------------------------------------------------------------------------
const setGameInformation = () => {
  document.title = `${globals.GAME_NAME} v${globals.VERSION}`;
};

// ------------------------------------------------------------------------------------------------
const initCanvas = () => {
  const canvas = document.getElementById("canvas") as HTMLCanvasElement;
  canvas.width = Math.floor(window.innerWidth * 0.9);
  canvas.height = Math.floor(window.innerHeight * 0.9);
  canvas.style.backgroundColor = "#111133";
  ctx = canvas.getContext("2d");
};

// ------------------------------------------------------------------------------------------------
const initTrees = (count: number) => {
  trees = [];
  for (let i = 0; i < count; i++) {
    const tmpTree = new Tree(
      60 + Math.random() * 80,
      globals.TREE_COLOURS[Math.floor(Math.random() * globals.TREE_COLOURS.length)],
      32 + Math.random() * 16,
      32 + Math.random() * 16,
      32 + Math.random() * 16,
      32 + Math.random() * 16,
      32 + Math.random() * 16,
      32 + Math.random() * 16,
      32 + Math.random() * 16,
    );
    trees.push(tmpTree);
  }
};

// ------------------------------------------------------------------------------------------------
const renderTrees = (ctx: CanvasRenderingContext2D) => {
  console.log("Drawing trees");
  trees.forEach((t) => drawTree(ctx, t));
};

// ------------------------------------------------------------------------------------------------
export const initGame = () => {
  setGameInformation();
  initCanvas();

  if (ctx) {
    initTrees(3);
    ctx.translate(250, 250);
    renderTrees(ctx);
    drawBalloon(ctx);
  } else {
    alert("Could not load 2D rendering context - unable to continue.");
  }
};
