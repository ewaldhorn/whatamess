import { drawBalloon, drawTree } from "./drawing.ts";
import globals from "./globals.ts";
import { Tree } from "./trees.ts";

// ------------------------------------------------------------------------------------------------
let ctx: CanvasRenderingContext2D | null;
let mainAreaWidth: number, mainAreaHeight: number;
let horizontalPadding: number, verticalPadding: number;
let trees: Array<Tree>;

// ------------------------------------------------------------------------------------------------
const setGameInformation = () => {
  document.title = `${globals.GAME_NAME} v${globals.VERSION}`;
};

// ------------------------------------------------------------------------------------------------
const initCanvas = () => {
  const canvas = document.getElementById("canvas") as HTMLCanvasElement;
  canvas.width = Math.floor(window.innerWidth * globals.CANVAS_SCALING_FACTOR);
  canvas.height = Math.floor(window.innerHeight * globals.CANVAS_SCALING_FACTOR);

  mainAreaHeight = canvas.height / globals.MAIN_AREA_RATIO;
  mainAreaWidth = canvas.width / globals.MAIN_AREA_RATIO;

  horizontalPadding = (window.innerWidth - mainAreaWidth) / 2;
  verticalPadding = (window.innerHeight - mainAreaHeight) / 2;

  canvas.style.backgroundColor = "#111133";
  ctx = canvas.getContext("2d");
};

// ------------------------------------------------------------------------------------------------
const initTrees = (count: number) => {
  trees = [];
  for (let i = 0; i < count; i++) {
    const tmpTree = new Tree(
      globals.MIN_TREE_SIZE + Math.random() * globals.MAX_TREE_SIZE,
      globals.TREE_COLOURS[Math.floor(Math.random() * globals.TREE_COLOURS.length)],
      globals.MIN_RANDOM_VALUE + Math.random() * globals.MAX_RANDOM_VALUE,
      globals.MIN_RANDOM_VALUE + Math.random() * globals.MAX_RANDOM_VALUE,
      globals.MIN_RANDOM_VALUE + Math.random() * globals.MAX_RANDOM_VALUE,
      globals.MIN_RANDOM_VALUE + Math.random() * globals.MAX_RANDOM_VALUE,
      globals.MIN_RANDOM_VALUE + Math.random() * globals.MAX_RANDOM_VALUE,
      globals.MIN_RANDOM_VALUE + Math.random() * globals.MAX_RANDOM_VALUE,
      globals.MIN_RANDOM_VALUE + Math.random() * globals.MAX_RANDOM_VALUE,
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
const render = (ctx: CanvasRenderingContext2D) => {
  ctx.clearRect(0, 0, window.innerWidth, window.innerHeight);

  ctx.save();
  ctx.translate(horizontalPadding, verticalPadding + mainAreaHeight);
  drawBalloon(ctx);

  renderTrees(ctx);
  ctx.restore();
};

// ------------------------------------------------------------------------------------------------
export const initGame = () => {
  setGameInformation();
  initCanvas();

  if (ctx) {
    initTrees(3);
    render(ctx);
  } else {
    alert("Could not load 2D rendering context - unable to continue.");
  }
};
