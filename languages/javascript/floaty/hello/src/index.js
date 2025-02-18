import { Engine } from "https://floaty.dev/engine-v1.js";

var engine = new Engine();
engine.expose();

var sprites = {
  "player-center": preload(
    "https://assets.floaty.dev/assertchris/shmup/player-center.spr.json",
  ),
  "player-left": preload(
    "https://assets.floaty.dev/assertchris/shmup/player-left.spr.json",
  ),
  "player-right": preload(
    "https://assets.floaty.dev/assertchris/shmup/player-right.spr.json",
  ),
  "heart-empty": preload(
    "https://assets.floaty.dev/assertchris/shmup/heart-empty.spr.json",
  ),
  "heart-full": preload(
    "https://assets.floaty.dev/assertchris/shmup/heart-full.spr.json",
  ),
  "thrust-0": preload(
    "https://assets.floaty.dev/assertchris/shmup/thrust-0.spr.json",
  ),
  "thrust-1": preload(
    "https://assets.floaty.dev/assertchris/shmup/thrust-1.spr.json",
  ),
  "thrust-2": preload(
    "https://assets.floaty.dev/assertchris/shmup/thrust-2.spr.json",
  ),
  "enemy-1-0": preload(
    "https://assets.floaty.dev/assertchris/shmup/enemy-1-0.spr.json",
  ),
  "enemy-1-1": preload(
    "https://assets.floaty.dev/assertchris/shmup/enemy-1-1.spr.json",
  ),
  "enemy-1-2": preload(
    "https://assets.floaty.dev/assertchris/shmup/enemy-1-2.spr.json",
  ),
  "enemy-1-3": preload(
    "https://assets.floaty.dev/assertchris/shmup/enemy-1-3.spr.json",
  ),
  "bullet-1": preload(
    "https://assets.floaty.dev/assertchris/shmup/bullet-1.spr.json",
  ),
};

var sounds = {
  shoot: preload("https://assets.floaty.dev/assertchris/shmup/shoot.sfx.json"),
  explode: preload(
    "https://assets.floaty.dev/assertchris/shmup/explode.sfx.json",
  ),
};

class Stars {
  limit = 50;

  init(engine) {
    this.engine = engine;
    this.stars = [];

    for (var i = 0; i < this.limit; i++) {
      this.stars.push(
        new Star({
          x: Math.floor(rnd(128)),
          y: Math.floor(rnd(128)),
          type: Math.floor(rnd(3)),
          engine: this.engine,
        }),
      );
    }
  }

  update() {
    this.stars = this.stars.filter((star) => star.y < 128);

    while (this.stars.length < this.limit) {
      this.stars.push(
        new Star({
          x: Math.floor(rnd(128)),
          y: 0,
          type: Math.floor(rnd(3)),
          engine: this.engine,
        }),
      );
    }

    for (var star of this.stars) {
      star.update();
    }
  }

  draw() {
    for (var star of this.stars) {
      star.draw();
    }
  }
}
class Star {
  x = 0;
  y = 0;
  type = 0;

  constructor({ x, y, type, engine }) {
    this.x = x;
    this.y = y;
    this.type = type;
    this.engine = engine;
  }

  update() {
    if (this.type === 0) {
      this.y += 1.5;
    }

    if (this.type === 1) {
      this.y += 1.0;
    }

    if (this.type === 2) {
      this.y += 0.5;
    }
  }

  draw() {
    if (this.type === 0) {
      pset(this.x, this.y, 6);
    }

    if (this.type === 1) {
      pset(this.x, this.y, 13);
    }

    if (this.type === 2) {
      pset(this.x, this.y, 1);
    }
  }
}
class GameUi {
  init(engine) {
    this.engine = engine;
    this.score = 0;
  }

  update() {}

  draw() {
    text(`move: wasd | fire: j | debug: i o p`, 0, 118, 5);
    text(`score: ${this.score}`, 0, 0, 5);

    for (var i = 0; i < player.maxHealth; i++) {
      if (i >= player.health) {
        spr("heart-empty", i * 8, 8);
      } else {
        spr("heart-full", i * 8, 8);
      }
    }
  }
}
class Player {
  speed = 1;
  hurtDelay = 1000;
  attackDelay = 250;
  maxHealth = 5;
  invisibleDelay = 100;

  init(engine) {
    this.engine = engine;
    this.x = 64;
    this.y = 64;
    this.moving = false;
    this.sprite = 2;
    this.thrustSprite = null;
    this.muzzleSize = 0;
    this.hurtTimer = null;
    this.attackTimer = null;
    this.health = this.maxHealth;
    this.invisible = false;
    this.invisibleTimer = null;
  }

  update() {
    this.moving = false;

    if (btn("ArrowLeft") || btn("a")) {
      this.x = Math.max(0, this.x - this.speed);
      this.sprite = "player-left";
      this.moving = true;
    }

    if (btn("ArrowRight") || btn("d")) {
      this.x = Math.min(this.x + this.speed, 120);
      this.sprite = "player-right";
      this.moving = true;
    }

    if (!btn("ArrowLeft") && !btn("a") && !btn("ArrowRight") && !btn("d")) {
      this.sprite = "player-center";
    }

    if (btn("ArrowUp") || btn("w")) {
      this.y = Math.max(0, this.y - this.speed);
      this.moving = true;
    }

    if (btn("ArrowDown") || btn("s")) {
      this.y = Math.min(this.y + this.speed, 120);
      this.moving = true;
    }

    if (this.moving) {
      if (this.thrustSprite == null) {
        this.thrustSprite = 0;
      } else {
        // see draw to understand why we use 0.2 when changing sprite numbers
        this.thrustSprite += 0.2;
      }

      if (this.thrustSprite > 2) {
        this.thrustSprite = 2;
      }
    } else {
      if (this.thrustSprite) {
        this.thrustSprite -= 0.2;
      }

      if (this.thrustSprite < 0) {
        this.thrustSprite = null;
      }
    }

    if (this.muzzleSize > 0) {
      // similar to thrustSprite but this time an integer radius instead of sprite index
      this.muzzleSize -= 0.75;
    }
  }

  draw() {
    if (!this.invisible) {
      spr(this.sprite, this.x, this.y);

      if (this.thrustSprite != null) {
        // rounding is one way to pick a whole sprite number but still allow gradual sprite changes
        // we adjust the sprite by 0.2 in update, so the whole number will only change every 5 frames
        spr(`thrust-${Math.round(this.thrustSprite)}`, this.x, this.y + 8);
      }

      if (this.muzzleSize > 0) {
        circFill(this.x + 3, this.y - 1, Math.round(this.muzzleSize), 7);
      }
    }
  }

  hurt() {
    if (!this.hurtTimer) {
      this.hurtTimer = setTimeout(() => {
        this.hurtTimer = null;
        clearTimeout(this.invisibleTimer);
        this.invisible = false;
      }, this.hurtDelay);

      this.invisibleTimer = setInterval(
        () => (this.invisible = !this.invisible),
        this.invisibleDelay,
      );

      this.health--;
      explode(30, this.x + 4, this.y + 4);
      sfx("explode");

      if (this.health < 1) {
        mode = "over";
      }
    }
  }

  canAttack() {
    return !this.attackTimer;
  }

  attack() {
    if (!this.canAttack()) {
      return;
    }

    this.attackTimer = setTimeout(
      () => (this.attackTimer = null),
      this.attackDelay,
    );
    player.muzzleSize = 7;
    sfx("shoot");
  }

  getBox() {
    return [this.x, this.y, 7, 7];
  }
}
class Bullets {
  init(engine) {
    this.engine = engine;
    this.bullets = [];
  }

  update() {
    this.bullets = this.bullets.filter(
      (bullet) => bullet.y > 0 && bullet.health > 0,
    );

    for (var bullet of this.bullets) {
      bullet.update();
    }

    if (btn("j") && player.canAttack()) {
      this.bullets.push(
        new Bullet({
          x: player.x,
          y: player.y,
          engine: this.engine,
        }),
      );

      player.attack();
    }
  }

  draw() {
    for (var bullet of this.bullets) {
      bullet.draw();
    }
  }
}
class Bullet {
  x = 0;
  y = 0;
  health = 1;
  speed = 2;
  hurtTimer = null;
  hurtDelay = 500;

  constructor({ x, y, engine }) {
    this.x = x;
    this.y = y;
    this.engine = engine;
  }

  update() {
    this.y -= this.speed;
  }

  draw() {
    spr("bullet-1", this.x, this.y);
  }

  hurt() {
    if (!this.hurtTimer) {
      this.hurtTimer = setTimeout(
        () => (this.hurtTimer = null),
        this.hurtDelay,
      );
      this.health--;
    }
  }

  getBox() {
    return [this.x, this.y, 7, 7];
  }
}
class Enemies {
  init(engine) {
    this.engine = engine;
    this.enemies = [];
    this.wave = 1;
  }

  spawn() {
    if (this.wave === 1) {
      this.enemies.push(
        new Enemy({
          x: 20,
          y: 20,
          engine: this.engine,
        }),
      );
    }
  }

  update() {
    this.enemies = this.enemies.filter((enemy) => enemy.health > 0);

    var playerBox = player.getBox();

    for (var enemy of this.enemies) {
      enemy.update();

      var enemyBox = enemy.getBox();

      if (this.engine.boxesCollide(playerBox, enemyBox)) {
        player.hurt();
      }

      for (var bullet of bullets.bullets) {
        var bulletBox = bullet.getBox();

        if (this.engine.boxesCollide(enemyBox, bulletBox)) {
          enemy.hurt();
          bullet.hurt();
        }
      }
    }
  }

  draw() {
    for (var enemy of this.enemies) {
      enemy.draw();
    }
  }
}
class Enemy {
  x = 0;
  y = 0;
  type = 1;
  frame = 0;
  health = 1;
  hurtTimer = null;
  hurtDelay = 500;

  constructor({ x, y, engine }) {
    this.x = x;
    this.y = y;
    this.engine = engine;
  }

  update() {
    this.frame += 0.2;

    if (this.frame > 3) {
      this.frame = 0;
    }
  }

  draw() {
    if (this.type === 1) {
      spr(`enemy-${this.type}-${Math.round(this.frame)}`, this.x, this.y);
    }
  }

  hurt() {
    if (!this.hurtTimer) {
      this.hurtTimer = setTimeout(
        () => (this.hurtTimer = null),
        this.hurtDelay,
      );
      this.health--;
      sfx("explode");
    }

    if (this.health < 1) {
      explode(30, this.x + 4, this.y + 4);
    }
  }

  getBox() {
    return [this.x, this.y, 7, 7];
  }
}
class Particles {
  init(engine) {
    this.engine = engine;
    this.particles = [];
  }

  update() {
    this.particles = this.particles.filter(
      (particle) => particle.age < particle.maxAge,
    );

    for (var particle of this.particles) {
      particle.update();
    }
  }

  draw() {
    for (var particle of this.particles) {
      particle.draw();
    }
  }

  spawn(number, x, y, options) {
    for (var i = 0; i < number; i++) {
      this.particles.push(
        new Particle({
          x,
          y,
          ...options,
          engine: this.engine,
        }),
      );
    }
  }
}
class Particle {
  x = 0;
  y = 0;
  age = 0;
  maxAge = 30;

  sx = 0;
  sy = 0;
  onInit = null;
  onUpdate = null;
  onDraw = null;

  constructor({ x, y, onInit, onUpdate, onDraw, engine }) {
    this.x = x;
    this.y = y;
    this.engine = engine;

    if (onUpdate) {
      this.onUpdate = onUpdate;
    }

    if (onDraw) {
      this.onDraw = onDraw;
    }

    if (onInit) {
      this.onInit = onInit;
      this.onInit.apply(this);
    }
  }

  update() {
    this.x += this.sx;
    this.y += this.sy;
    this.age += 1;

    if (this.onUpdate) {
      this.onUpdate.apply(this);
    }
  }

  draw() {
    if (this.onDraw) {
      this.onDraw.apply(this);
    }
  }
}

function explode(amount, x, y) {
  particles.spawn(amount, x, y, {
    onInit() {
      this.sx = this.engine.randomFloatBetween(-0.5, 0.5);
      this.sy = this.engine.randomFloatBetween(-0.5, 0.5);
      this.maxAge = this.engine.randomIntegerBetween(20, 40);
      this.radius = this.engine.randomIntegerBetween(1, 4);
      this.color = 7;
    },
    onUpdate() {
      if (this.age >= (this.maxAge * 1) / 6) {
        this.color = 10;
      }
      if (this.age >= (this.maxAge * 2) / 6) {
        this.color = 9;
      }
      if (this.age >= (this.maxAge * 3) / 6) {
        this.color = 8;
      }
      if (this.age >= (this.maxAge * 4) / 6) {
        this.color = 2;
      }
      if (this.age >= (this.maxAge * 5) / 6) {
        this.color = 5;
      }

      this.sx *= 0.99;
      this.sy *= 0.99;
    },
    onDraw() {
      circFill(this.x, this.y, this.radius, this.color);
    },
  });
}

var mode = "load";
var restartTimer = null;
var restartDelay = 2500;
var waveTimer = null;
var waveDelay = 2500;
var winTimer = null;
var winDelay = 2500;

var stars = new Stars();
var gameUi = new GameUi();
var player = new Player();
var bullets = new Bullets();
var enemies = new Enemies();
var particles = new Particles();

async function init() {
  stars.init(this);
  particles.init(this);

  await this.load();

  mode = "start";
}

function update() {
  stars.update();
  particles.update();

  if (mode === "start" && btnp()) {
    gameUi.init(this);
    player.init(this);
    bullets.init(this);
    enemies.init(this);
    mode = "wave";
  } else if (mode === "wave" && !waveTimer) {
    waveTimer = setTimeout(() => {
      waveTimer = null;
      mode = "game";
    }, waveDelay);
  } else if (mode === "over" && !restartTimer) {
    restartTimer = setTimeout(() => {
      restartTimer = null;
      mode = "start";
    }, restartDelay);
  } else if (mode === "win" && !winTimer) {
    winTimer = setTimeout(() => {
      winTimer = null;
      mode = "start";
    }, winDelay);
  } else if (mode === "game") {
    gameUi.update();
    player.update();
    bullets.update();
    enemies.update();
  }
}

function draw() {
  cls();

  stars.draw();
  particles.draw();

  if (mode === "load") {
    text("Loading...", 45, 56);
  } else if (mode === "start") {
    text("Press any key to start", 15, 56);
    text("Dev by assertchris", 15, 96, 5);
    text("Font by kenney", 15, 104, 5);
  } else if (mode === "wave") {
    text(`Wave ${enemies.wave}`, 45, 56);
  } else if (mode === "over") {
    text("Game over", 45, 56);
  } else if (mode === "win") {
    text("You win!", 45, 56);
  } else if (mode === "game") {
    gameUi.draw();
    player.draw();
    bullets.draw();
    enemies.draw();
  }

  // DEBUG
  if (btnp("i")) {
    mode = "win";
  }

  // DEBUG
  if (btnp("o")) {
    mode = "over";
  }

  // DEBUG
  if (btnp("p")) {
    enemies.spawn();
  }
}

engine.start({
  sprites,
  sounds,
  init,
  update,
  draw,
  target: document.querySelector(".game"),
});
