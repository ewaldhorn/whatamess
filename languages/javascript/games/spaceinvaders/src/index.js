import { clearGameArea, clearScreen, drawCanvasBorder } from "./canvasutils.js";
import { Player } from "./player.js";
import { Bullet } from "./bullet.js";
import { showPauseScreen } from "./pause_screen.js";
// ----------------------------------------------------------------------------
//                                                                      GLOBALS
/** @type {HTMLCanvasElement} */
const canvas = document.getElementById("canvas");

/** @type {CanvasRenderingContext2D} */
const ctx = canvas.getContext("2d");

let isPlaying = false;
let isPaused = false;

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
var bullets = [];

// ----------------------------------------------------------------------------
//                                                                    GAME LOOP
const gameLoop = () => {
  if (isPlaying) {
    if (isPaused) {
      showPauseScreen(ctx);
    } else {
      clearGameArea(ctx);

      for (let i = 0; i < bullets.length; i++) {
        bullets[i].draw(ctx);
        bullets[i].update();
      }

      bullets = bullets.filter((element) => element.isAlive());

      p.update();
      p.draw(ctx);
    }
  }
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
    case "Space": {
      let b = new Bullet({
        position: p.getBulletPosition(),
        velocity: { x: Math.floor(0.05 * p.velocity.x), y: -7 },
      });
      bullets.push(b);
      event.preventDefault();
      break;
    }
    case "KeyP": {
      isPaused = !isPaused;
      event.preventDefault();
      break;
    }
    // default: {
    //   console.log(event.code);
    // }
  }
});

// ----------------------------------------------------------------------------
//                                                                  ENTRY POINT
prepGameScreen();
isPlaying = true;
gameLoop();
