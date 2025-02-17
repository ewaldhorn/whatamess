import { clearGameArea, clearScreen, drawCanvasBorder } from "./canvasutils.js";
import { Player } from "./player.js";
// ----------------------------------------------------------------------------
//                                                                      GLOBALS
const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");

// ----------------------------------------------------------------------------
//                                                             CONFIGURE CANVAS
canvas.width = innerWidth - 50;
canvas.height = innerHeight - 100;

// ----------------------------------------------------------------------------
const prepGameScreen = () => {
  clearScreen(ctx);
  drawCanvasBorder(ctx);
  clearGameArea(ctx);
};

// ----------------------------------------------------------------------------
// todo : move into a class
const p = new Player();

// ----------------------------------------------------------------------------
//                                                                    GAME LOOP
const gameLoop = () => {
  clearGameArea(ctx);
  p.draw(ctx);
  requestAnimationFrame(gameLoop);
};

// ----------------------------------------------------------------------------
//                                                                  ENTRY POINT
prepGameScreen();
gameLoop();
