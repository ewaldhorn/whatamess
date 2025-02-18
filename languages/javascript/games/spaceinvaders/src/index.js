import {
  clearGameArea,
  clearScreen,
  drawCanvasBorder,
} from "./canvas_utils.js";
import { Enemy } from "./enemy.js";
import { Player } from "./player.js";
import { Bullet } from "./bullet.js";
import { Star } from "./star.js";
import { showPauseScreen } from "./pause_screen.js";
// ----------------------------------------------------------------------------
//                                                                      GLOBALS
/** @type {HTMLCanvasElement} */
const canvas = document.getElementById("canvas");

/** @type {CanvasRenderingContext2D} */
const ctx = canvas.getContext("2d");

let isPlaying = false;
let isPaused = false;
let isOver = false;
let score = 0;
let fps = 0;
let frames = 0;
let lastTime = performance.now();

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

/** @type {Star[]} */
var stars = [];

// ----------------------------------------------------------------------------
const createInitialStars = () => {
  for (let i = 0; i < 100; i++) {
    let s = new Star({
      position: {
        x: 50 + Math.floor(Math.random() * (canvas.width - 60)),
        y: 3 + Math.floor(Math.random() * (canvas.height - 60)),
      },
      velocity: { x: 0, y: 1 + Math.floor(Math.random() * 3) },
      colour: "#ffffff",
    });

    s.colour = s.getRandomColour();

    stars.push(s);
  }
};

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
    } else if (isOver) {
      clearGameArea(ctx);
      showScore(ctx);
      showIsOverScreen(ctx);
    } else {
      clearGameArea(ctx);
      showScore(ctx);
      showFPS(ctx);

      for (let i = 0; i < stars.length; i++) {
        stars[i].draw(ctx);
        stars[i].update();
      }

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
          score += 1;
        }
      }

      // check if bullets intercepted anything
      for (let i = 0; i < bullets.length; i++) {
        // check if a downward moving bullet hit a player
        if (bullets[i].velocity.y > 0 && bullets[i].isAlive()) {
          if (
            bullets[i].position.x >= p.position.x &&
            bullets[i].position.x <= p.position.x + p.width &&
            bullets[i].position.y >= p.position.y - p.height &&
            bullets[i].position.y <= p.position.y
          ) {
            p.health -= 4;
            bullets[i].position.x = 0;
            bullets[i].position.y = 0;
            bullets[i].velocity.y = 0;
            score -= 8;
          }
        }

        // check if an upward moving bullet hit an enemy
        if (bullets[i].velocity.y < 0 && bullets[i].isAlive()) {
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
              score += 11;
            }
          }
        }
      }

      bullets = bullets.filter((element) => element.isAlive());
      enemies = enemies.filter((element) => element.isAlive());

      p.update();
      p.draw(ctx);

      // check if any of the end conditions are true
      if (p.health <= 0 || enemies.length == 0) {
        isOver = true;
      }
    }
  }

  frames++;
  const currentTime = performance.now();
  const timeDiff = currentTime - lastTime;

  if (timeDiff >= 1000) {
    fps = Math.trunc(frames);
    frames = 0;
    lastTime = currentTime;
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
        score -= 1;
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
/**
 * @param {CanvasRenderingContext2D=required} ctx - 2D rendering context
 */
const showScore = (ctx) => {
  ctx.font = "18px Arial";
  ctx.textAlign = "left";
  ctx.textBaseline = "middle";
  ctx.fillStyle = "green";

  const text = `Score: ${score}`;
  // const metrics = ctx.measureText(text);
  // const textWidth = metrics.width;

  ctx.fillText(text, 10, 20);
};

// ----------------------------------------------------------------------------
/**
 * @param {CanvasRenderingContext2D=required} ctx - 2D rendering context
 */
const showFPS = (ctx) => {
  ctx.font = "18px Arial";
  ctx.textAlign = "left";
  ctx.textBaseline = "middle";
  ctx.fillStyle = "green";

  const text = `FPS: ${fps}`;
  const metrics = ctx.measureText(text);
  const textWidth = metrics.width;

  ctx.fillText(text, canvas.width - textWidth - 5, 20);
};

// ----------------------------------------------------------------------------
const showIsOverScreen = (ctx) => {
  ctx.font = "36px Arial";
  ctx.textAlign = "left";
  ctx.textBaseline = "middle";
  ctx.fillStyle = "yellow";

  let text = "GAME OVER";
  let metrics = ctx.measureText(text);
  let textWidth = metrics.width;

  ctx.fillText(text, (canvas.width - textWidth) / 2, canvas.height / 2 - 50);

  if (p.isAlive()) {
    ctx.fillStyle = "green";
    text = "YOU WON";
  } else {
    ctx.fillStyle = "red";
    text = "YOU LOST";
  }

  metrics = ctx.measureText(text);
  textWidth = metrics.width;

  ctx.fillText(text, (canvas.width - textWidth) / 2, canvas.height / 2);
};

// ----------------------------------------------------------------------------
//                                                                  ENTRY POINT
prepGameScreen();
createFirstEnemies();
createInitialStars();
isPlaying = true;
gameLoop();
