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
export const initGame = () => {
  setGameInformation();
  initCanvas();

  if (ctx) {
    initTrees(3);
  } else {
    alert("Could not load 2D rendering context - unable to continue");
  }
};
