import {
  clearGameArea,
  clearScreen,
  drawCanvasBorder,
} from "./canvas_utils.js";
import { Enemy } from "./enemy.js";
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
// todo : move game state into a class
const p = new Player();

/** @type {Bullet[]} */
var bullets = [];

/** @type {Enemy[]} */
var enemies = [];

// ----------------------------------------------------------------------------
const createFirstEnemies = () => {
  let enemy = new Enemy();
  enemies.push(enemy);

  enemy = new Enemy();
  enemy.position.y += enemy.height + 5;
  enemy.position.x = 5 + Math.floor((Math.random() * canvas.width) / 2);
  if (Math.random() < 0.5) {
    enemy.reverseHorizontalDirection();
  }
  enemy.colour = "#cc1156";
  enemies.push(enemy);

  enemy = new Enemy();
  enemy.position.y += enemy.height * 2 + 10;
  enemy.position.x =
    canvas.width / 2 + Math.floor((Math.random() * canvas.width) / 4);
  enemy.reverseHorizontalDirection();
  enemy.colour = "#44cc56";
  enemies.push(enemy);
};

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

      for (let i = 0; i < enemies.length; i++) {
        enemies[i].update();
        enemies[i].draw(ctx);
        if (Math.random() < 0.25 && enemies[i].canShoot()) {
          let b = new Bullet({
            position: enemies[i].getBulletPosition(),
            velocity: { x: Math.floor(0.05 * p.velocity.x), y: 5 },
          });
          bullets.push(b);
        }
      }

      // check if bullets intercepted anything
      for (let i = 0; i < bullets.length; i++) {
        // check if a downward moving bullet hit a player
        if (bullets[i].velocity.y < 0 && bullets[i].isAlive()) {
          if (
            bullets[i].position.x >= p.position.x &&
            bullets[i].position.x <= p.position.x + p.width &&
            bullets[i].position.y >= p.position.y &&
            bullets[i].position.y <= ep.position.y + p.height
          ) {
            p.health -= 4;
            bullets[i].position.x = 0;
            bullets[i].position.y = 0;
            bullets[i].velocity.y = 0;
          }
        }

        // check if an upward moving bullet hit an enemy
        if (bullets[i].velocity.y > 0 && bullets[i].isAlive()) {
          for (let j = 0; j < enemies.length; j++) {
            if (
              bullets[i].position.x >= enemies[j].position.x &&
              bullets[i].position.x <=
                enemies[j].position.x + enemies[j].width &&
              bullets[i].position.y >= enemies[j].position.y &&
              bullets[i].position.y <= enemies[j].position.y + enemies[j].height
            ) {
              enemies[j].health -= 4;
              bullets[i].position.x = 0;
              bullets[i].position.y = 0;
              bullets[i].velocity.y = 0;
            }
          }
        }
      }

      bullets = bullets.filter((element) => element.isAlive());
      enemies = enemies.filter((element) => element.isAlive());

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
      if (!isPaused) {
        let b = new Bullet({
          position: p.getBulletPosition(),
          velocity: { x: Math.floor(0.05 * p.velocity.x), y: -7 },
        });
        bullets.push(b);
      }
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
createFirstEnemies();
isPlaying = true;
gameLoop();
