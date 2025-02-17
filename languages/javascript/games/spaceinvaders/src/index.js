import { clearGameArea, clearScreen, drawCanvasBorder } from "./canvasutils.js";
import { Player } from "./player.js";
import { Bullet } from "./bullet.js";
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
var bullets = [];

// ----------------------------------------------------------------------------
//                                                                    GAME LOOP
const gameLoop = () => {
  isPlaying = true;
  clearGameArea(ctx);

  for (let i = 0; i < bullets.length; i++) {
    bullets[i].draw(ctx);
    bullets[i].update();
  }

  bullets = bullets.filter((element) => element.isAlive());

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
    case "Space": {
      let b = new Bullet({
        position: p.getBulletPosition(),
        velocity: { x: 0, y: -5 },
      });
      bullets.push(b);
      event.preventDefault();
    }
    default: {
      console.log(event.code);
    }
  }
});

// ----------------------------------------------------------------------------
//                                                                  ENTRY POINT
prepGameScreen();
gameLoop();
