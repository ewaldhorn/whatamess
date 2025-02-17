import { clearGameArea, clearScreen, drawCanvasBorder } from "./canvasutils.js";
import { Player } from "./player.js";
// ----------------------------------------------------------------------------
//                                                                      GLOBALS
const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");
let isPlaying = false;

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
  isPlaying = true;
  clearGameArea(ctx);
  p.update();
  p.draw(ctx);
  requestAnimationFrame(gameLoop);
};

// ----------------------------------------------------------------------------
//                                                              EVENT LISTENERS
addEventListener("keydown", (event) => {
  switch (event.code) {
    case "ArrowLeft": {
      p.goLeft();
      event.preventDefault();
      break;
    }
    case "ArrowRight": {
      p.goRight();
      event.preventDefault();
      break;
    }
  }
});

// ----------------------------------------------------------------------------
//                                                                  ENTRY POINT
prepGameScreen();
gameLoop();
