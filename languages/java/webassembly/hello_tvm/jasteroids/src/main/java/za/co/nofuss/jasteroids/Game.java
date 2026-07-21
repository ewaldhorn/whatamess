package za.co.nofuss.jasteroids;

import org.teavm.jso.browser.Storage;

/**
 * Game.java — Destroids game state and simulation, ported from the Odin original's game.odin.
 * Entities live in fixed-size arrays with an "active" flag, same as the Odin version (no real
 * need for that here since Java has a GC, but it keeps the port a faithful line-by-line mirror).
 */
final class Game {
    private Game() {
    }

    static final int CANVAS_W = 800;
    static final int CANVAS_H = 600;

    // Key slots — index order matches KEYMAP handling in Client.java.
    static final int KEY_LEFT = 0;
    static final int KEY_RIGHT = 1;
    static final int KEY_THRUST = 2;
    static final int KEY_FIRE = 3;
    static final int KEY_START = 4;
    static final int KEY_HYPER = 5;
    static final int KEY_PAUSE = 6;
    static final int KEY_COUNT = 7;

    static final double TAU = Math.PI * 2;

    static final int STAR_COUNT = 380;
    static final int MAX_ASTEROIDS = 256;
    static final int MAX_BULLETS = 6;
    static final int MAX_PARTICLES = 700;
    static final int MAX_DEBRIS = 160;
    static final int MAX_SHOCKWAVES = 16;

    static final int SHAPE_N = 12;

    static final double SHIP_TURN = 3.6;
    static final double SHIP_THRUST = 260.0;
    static final double SHIP_FRICTION = 0.6;
    static final double SHIP_MAX_SPEED = 420.0;
    static final double SHIP_RADIUS = 11.0;

    static final double BULLET_SPEED = 520.0;
    static final double BULLET_LIFE = 1.15;
    static final double FIRE_COOLDOWN = 0.16;

    static final double INVULN_TIME = 2.5;
    static final double RESPAWN_DELAY = 1.2;

    static final double AST_R_LARGE = 46.0;
    static final double AST_R_MED = 26.0;
    static final double AST_R_SMALL = 14.0;

    // Sound ids — kept in sync with the switch in index.html's playSound().
    static final int SND_FIRE = 0;
    static final int SND_BANG_LARGE = 1;
    static final int SND_BANG_MED = 2;
    static final int SND_BANG_SMALL = 3;
    static final int SND_SHIP_DEATH = 4;
    static final int SND_LEVEL = 5;
    static final int SND_HYPER = 6;
    static final int SND_THUD = 7;

    // ------------------------------------------------------------------------------------------
    // Entities
    static final class Star {
        double x, y;
        double brightness;
        double twinkle;
    }

    static final class Ship {
        double x, y;
        double vx, vy;
        double angle;
        boolean alive;
        boolean thrust;
        double invuln;
        double respawn; // >0 while waiting to respawn after death
        double cooldown;
    }

    static final class Bullet {
        double x, y;
        double vx, vy;
        double life;
        boolean active;
    }

    static final class Asteroid {
        double x, y;
        double vx, vy;
        double angle;
        double spin;
        double radius;
        int size;
        double[] shape = new double[SHAPE_N];
        boolean active;
    }

    static final class Particle {
        double x, y;
        double vx, vy;
        double life;
        double maxLife;
        boolean active;
    }

    static final class Debris {
        double x, y;
        double vx, vy;
        double angle;
        double spin;
        double len;
        double life;
        double maxLife;
        boolean active;
    }

    static final class Shockwave {
        double x, y;
        double radius;
        double maxRadius;
        double life;
        double maxLife;
        boolean active;
    }

    // ------------------------------------------------------------------------------------------
    // State
    static final Star[] stars = new Star[STAR_COUNT];
    static final Ship ship = new Ship();
    static final Bullet[] bullets = new Bullet[MAX_BULLETS];
    static final Asteroid[] asteroids = new Asteroid[MAX_ASTEROIDS];
    static final Particle[] particles = new Particle[MAX_PARTICLES];
    static final Debris[] debris = new Debris[MAX_DEBRIS];
    static final Shockwave[] shockwaves = new Shockwave[MAX_SHOCKWAVES];

    static int score;
    static int bestScore;
    static int lives;
    static int level;
    static boolean gameOver;
    static boolean thrustSndOn;

    // intro is true from boot until the player's first ENTER/SPACE press — the simulation is
    // fully frozen (but still drawn) behind a splash. Never reset by resetGame().
    static boolean intro = true;

    static boolean paused;

    // Screen shake: shakeMag decays each frame; shakeX/shakeY are the per-frame random offset
    // drawn from it, read by Draw to jitter the world transform.
    static double shakeMag;
    static double shakeX;
    static double shakeY;

    static double hyperCooldown;
    static double thudCooldown; // throttles the asteroid-collision thud

    static double thrustParticleTimer;

    static final boolean[] keys = new boolean[KEY_COUNT];
    static final boolean[] justPressed = new boolean[KEY_COUNT];

    static void keyEvent(int code, boolean down) {
        if (code < KEY_COUNT) {
            keys[code] = down;
            if (down) {
                justPressed[code] = true;
            }
        }
    }

    private static boolean consumePressed(int code) {
        if (code < KEY_COUNT && justPressed[code]) {
            justPressed[code] = false;
            return true;
        }
        return false;
    }

    // ------------------------------------------------------------------------------------------
    // Setup
    static void init() {
        for (int i = 0; i < STAR_COUNT; i++) {
            Star s = new Star();
            s.x = Math.random() * CANVAS_W;
            s.y = Math.random() * CANVAS_H;
            s.brightness = frange(0.15, 0.9);
            s.twinkle = frange(0, TAU);
            stars[i] = s;
        }
        for (int i = 0; i < MAX_BULLETS; i++) bullets[i] = new Bullet();
        for (int i = 0; i < MAX_ASTEROIDS; i++) asteroids[i] = new Asteroid();
        for (int i = 0; i < MAX_PARTICLES; i++) particles[i] = new Particle();
        for (int i = 0; i < MAX_DEBRIS; i++) debris[i] = new Debris();
        for (int i = 0; i < MAX_SHOCKWAVES; i++) shockwaves[i] = new Shockwave();

        bestScore = getBestScore();
        resetGame();
    }

    static void resetGame() {
        score = 0;
        lives = 3;
        level = 0;
        gameOver = false;

        for (Bullet b : bullets) b.active = false;
        for (Asteroid a : asteroids) a.active = false;
        for (Particle p : particles) p.active = false;
        for (Debris d : debris) d.active = false;
        for (Shockwave s : shockwaves) s.active = false;

        shakeMag = 0;
        shakeX = 0;
        shakeY = 0;
        hyperCooldown = 0;
        thudCooldown = 0;
        thrustParticleTimer = 0;
        paused = false;

        spawnShip(true);
        nextLevel();
    }

    private static void spawnShip(boolean protectedSpawn) {
        ship.x = CANVAS_W / 2.0;
        ship.y = CANVAS_H / 2.0;
        ship.vx = 0;
        ship.vy = 0;
        ship.angle = -TAU / 4; // pointing up
        ship.alive = true;
        ship.thrust = false;
        ship.invuln = protectedSpawn ? INVULN_TIME : 0;
        ship.respawn = 0;
        ship.cooldown = 0;
    }

    private static void nextLevel() {
        level += 1;
        if (level > 1) Env.playSound(SND_LEVEL); // no jingle for the opening wave
        int count = 4 + level * 2;
        for (int i = 0; i < count; i++) {
            double px, py;
            if (Math.random() < 0.5) {
                px = Math.random() * CANVAS_W;
                py = Math.random() < 0.5 ? 0 : CANVAS_H;
            } else {
                px = Math.random() < 0.5 ? 0 : CANVAS_W;
                py = Math.random() * CANVAS_H;
            }
            spawnAsteroid(px, py, 3);
        }
    }

    private static void spawnAsteroid(double px, double py, int size) {
        for (Asteroid a : asteroids) {
            if (a.active) continue;
            a.active = true;
            a.x = px;
            a.y = py;
            a.size = size;
            switch (size) {
                case 3:
                    a.radius = AST_R_LARGE;
                    break;
                case 2:
                    a.radius = AST_R_MED;
                    break;
                default:
                    a.radius = AST_R_SMALL;
                    break;
            }
            double speed = frange(24, 60) * (1.0 + (3 - size) * 0.35);
            double dir = frange(0, TAU);
            a.vx = Math.cos(dir) * speed;
            a.vy = Math.sin(dir) * speed;
            a.angle = frange(0, TAU);
            a.spin = frange(-1.6, 1.6);
            for (int i = 0; i < SHAPE_N; i++) {
                a.shape[i] = frange(0.72, 1.08);
            }
            return;
        }
    }

    // ------------------------------------------------------------------------------------------
    // Helpers
    private static double wrapX(double x) {
        if (x < 0) x += CANVAS_W;
        if (x >= CANVAS_W) x -= CANVAS_W;
        return x;
    }

    private static double wrapY(double y) {
        if (y < 0) y += CANVAS_H;
        if (y >= CANVAS_H) y -= CANVAS_H;
        return y;
    }

    private static void spawnParticles(double px, double py, int n, double spread) {
        int made = 0;
        for (Particle p : particles) {
            if (made >= n) break;
            if (p.active) continue;
            p.active = true;
            p.x = px;
            p.y = py;
            double dir = frange(0, TAU);
            double speed = frange(30, spread);
            p.vx = Math.cos(dir) * speed;
            p.vy = Math.sin(dir) * speed;
            p.maxLife = frange(0.4, 1.0);
            p.life = p.maxLife;
            made += 1;
        }
    }

    private static void spawnDebris(double px, double py, int n, double spread, double sizeScale) {
        int made = 0;
        for (Debris d : debris) {
            if (made >= n) break;
            if (d.active) continue;
            d.active = true;
            d.x = px;
            d.y = py;
            double dir = frange(0, TAU);
            double speed = frange(20, spread);
            d.vx = Math.cos(dir) * speed;
            d.vy = Math.sin(dir) * speed;
            d.angle = frange(0, TAU);
            d.spin = frange(-6, 6);
            d.len = frange(3, 7) * sizeScale;
            d.maxLife = frange(0.5, 1.1);
            d.life = d.maxLife;
            made += 1;
        }
    }

    private static void spawnShockwave(double px, double py, double maxRadius) {
        for (Shockwave s : shockwaves) {
            if (s.active) continue;
            s.active = true;
            s.x = px;
            s.y = py;
            s.radius = 4;
            s.maxRadius = maxRadius;
            s.maxLife = 0.4;
            s.life = s.maxLife;
            return;
        }
    }

    private static void triggerShake(double amount) {
        if (amount > shakeMag) shakeMag = amount;
    }

    private static void fireBullet() {
        for (Bullet b : bullets) {
            if (b.active) continue;
            b.active = true;
            double dirX = Math.cos(ship.angle);
            double dirY = Math.sin(ship.angle);
            b.x = ship.x + dirX * SHIP_RADIUS;
            b.y = ship.y + dirY * SHIP_RADIUS;
            b.vx = ship.vx + dirX * BULLET_SPEED;
            b.vy = ship.vy + dirY * BULLET_SPEED;
            b.life = BULLET_LIFE;
            ship.vx -= dirX * (BULLET_SPEED * 0.05); // small recoil kick
            ship.vy -= dirY * (BULLET_SPEED * 0.05);
            Env.playSound(SND_FIRE);
            return;
        }
    }

    private static void splitAsteroid(Asteroid a) {
        switch (a.size) {
            case 3:
                score += 20;
                Env.playSound(SND_BANG_LARGE);
                triggerShake(7);
                break;
            case 2:
                score += 50;
                Env.playSound(SND_BANG_MED);
                triggerShake(4);
                break;
            default:
                score += 100;
                Env.playSound(SND_BANG_SMALL);
                triggerShake(2);
                break;
        }
        spawnParticles(a.x, a.y, a.size * 6 + 6, a.radius * 4);
        spawnDebris(a.x, a.y, a.size * 2 + 2, a.radius * 3, a.size);
        spawnShockwave(a.x, a.y, a.radius * 1.8);
        if (a.size > 1) {
            spawnAsteroid(a.x, a.y, a.size - 1);
            spawnAsteroid(a.x, a.y, a.size - 1);
        }
        a.active = false;
    }

    private static int countAsteroids() {
        int n = 0;
        for (Asteroid a : asteroids) {
            if (a.active) n += 1;
        }
        return n;
    }

    // ------------------------------------------------------------------------------------------
    // Simulation
    static void update(double dt) {
        // Drain every one-shot key latch exactly once per frame, regardless of which state we're
        // in below — see main.odin's update_game for why this must be unconditional.
        boolean justStart = consumePressed(KEY_START);
        boolean justFire = consumePressed(KEY_FIRE);
        boolean justPause = consumePressed(KEY_PAUSE);
        boolean justHyper = consumePressed(KEY_HYPER);

        if (intro) {
            if (justStart || justFire) intro = false;
            return;
        }

        if (!gameOver && justPause) paused = !paused;
        if (!gameOver && paused) {
            if (thrustSndOn) {
                thrustSndOn = false;
                Env.setThrust(false);
            }
            return;
        }

        if (thudCooldown > 0) thudCooldown -= dt;

        if (gameOver) {
            if (justStart) {
                resetGame();
            }
            updateParticles(dt);
            updateDebris(dt);
            updateShockwaves(dt);
            updateShake(dt);
            updateThrustSound();
            return;
        }

        updateShip(dt, justHyper);
        updateBullets(dt);
        updateAsteroids(dt);
        updateParticles(dt);
        updateDebris(dt);
        updateShockwaves(dt);
        updateShake(dt);
        handleCollisions();
        updateThrustSound();

        if (countAsteroids() == 0) {
            nextLevel();
        }
    }

    private static void updateThrustSound() {
        boolean want = ship.alive && ship.thrust && !gameOver;
        if (want != thrustSndOn) {
            thrustSndOn = want;
            Env.setThrust(want);
        }
    }

    private static void updateShip(double dt, boolean justHyper) {
        if (ship.respawn > 0) {
            ship.respawn -= dt;
            if (ship.respawn <= 0) {
                spawnShip(true);
            }
            return;
        }
        if (!ship.alive) return;

        if (ship.invuln > 0) ship.invuln -= dt;
        if (ship.cooldown > 0) ship.cooldown -= dt;
        if (hyperCooldown > 0) hyperCooldown -= dt;

        if (justHyper && hyperCooldown <= 0) {
            hyperspaceJump();
            hyperCooldown = 0.5;
        }
        if (!ship.alive) return; // hyperspace malfunction may have just killed the ship

        if (keys[KEY_LEFT]) ship.angle -= SHIP_TURN * dt;
        if (keys[KEY_RIGHT]) ship.angle += SHIP_TURN * dt;

        ship.thrust = keys[KEY_THRUST];
        if (ship.thrust) {
            ship.vx += Math.cos(ship.angle) * SHIP_THRUST * dt;
            ship.vy += Math.sin(ship.angle) * SHIP_THRUST * dt;

            thrustParticleTimer -= dt;
            if (thrustParticleTimer <= 0) {
                double backX = ship.x - Math.cos(ship.angle) * SHIP_RADIUS;
                double backY = ship.y - Math.sin(ship.angle) * SHIP_RADIUS;
                spawnParticles(backX, backY, 2, 40);
                thrustParticleTimer = 0.03;
            }
        }

        // Exponential drag toward zero.
        double drag = Math.pow(SHIP_FRICTION, dt);
        ship.vx *= drag;
        ship.vy *= drag;

        double speed = Math.sqrt(ship.vx * ship.vx + ship.vy * ship.vy);
        if (speed > SHIP_MAX_SPEED) {
            ship.vx *= SHIP_MAX_SPEED / speed;
            ship.vy *= SHIP_MAX_SPEED / speed;
        }

        ship.x = wrapX(ship.x + ship.vx * dt);
        ship.y = wrapY(ship.y + ship.vy * dt);

        if (keys[KEY_FIRE] && ship.cooldown <= 0) {
            fireBullet();
            ship.cooldown = FIRE_COOLDOWN;
        }
    }

    // hyperspaceJump teleports the ship to a random point on the field — classic Asteroids risk:
    // a small chance of a "malfunction" that destroys the ship instead of relocating it.
    private static void hyperspaceJump() {
        spawnParticles(ship.x, ship.y, 14, 220);
        Env.playSound(SND_HYPER);

        if (Math.random() < 0.14) {
            killShip();
            return;
        }

        ship.x = frange(40, CANVAS_W - 40);
        ship.y = frange(40, CANVAS_H - 40);
        ship.vx *= 0.2; // bleed off most velocity on arrival
        ship.vy *= 0.2;
        if (ship.invuln < 0.4) ship.invuln = 0.4;
        spawnParticles(ship.x, ship.y, 14, 220);
    }

    private static void updateBullets(double dt) {
        for (Bullet b : bullets) {
            if (!b.active) continue;
            b.x = wrapX(b.x + b.vx * dt);
            b.y = wrapY(b.y + b.vy * dt);
            b.life -= dt;
            if (b.life <= 0) b.active = false;
        }
    }

    private static void updateAsteroids(double dt) {
        for (Asteroid a : asteroids) {
            if (!a.active) continue;
            a.x = wrapX(a.x + a.vx * dt);
            a.y = wrapY(a.y + a.vy * dt);
            a.angle += a.spin * dt;
        }
    }

    private static void updateParticles(double dt) {
        for (Particle p : particles) {
            if (!p.active) continue;
            p.x += p.vx * dt;
            p.y += p.vy * dt;
            double decay = Math.pow(0.2, dt);
            p.vx *= decay;
            p.vy *= decay;
            p.life -= dt;
            if (p.life <= 0) p.active = false;
        }
    }

    private static void updateDebris(double dt) {
        for (Debris d : debris) {
            if (!d.active) continue;
            d.x += d.vx * dt;
            d.y += d.vy * dt;
            double decay = Math.pow(0.25, dt);
            d.vx *= decay;
            d.vy *= decay;
            d.angle += d.spin * dt;
            d.life -= dt;
            if (d.life <= 0) d.active = false;
        }
    }

    private static void updateShockwaves(double dt) {
        for (Shockwave s : shockwaves) {
            if (!s.active) continue;
            s.life -= dt;
            if (s.life <= 0) {
                s.active = false;
                continue;
            }
            double t = 1 - s.life / s.maxLife;
            s.radius = s.maxRadius * t;
        }
    }

    // updateShake decays the current shake magnitude and redraws the per-frame random offset
    // from it; Draw reads shakeX/shakeY to jitter the world transform.
    private static void updateShake(double dt) {
        if (shakeMag > 0.05) {
            shakeX = frange(-shakeMag, shakeMag);
            shakeY = frange(-shakeMag, shakeMag);
            shakeMag *= Math.pow(0.02, dt);
        } else {
            shakeMag = 0;
            shakeX = 0;
            shakeY = 0;
        }
    }

    private static void handleCollisions() {
        // Bullets vs asteroids.
        for (Bullet b : bullets) {
            if (!b.active) continue;
            for (Asteroid a : asteroids) {
                if (!a.active) continue;
                if (dist2(b.x, b.y, a.x, a.y) <= a.radius * a.radius) {
                    b.active = false;
                    splitAsteroid(a);
                    break;
                }
            }
        }

        // Asteroids vs ship.
        if (ship.alive && ship.invuln <= 0) {
            for (Asteroid a : asteroids) {
                if (!a.active) continue;
                double r = a.radius + SHIP_RADIUS;
                if (dist2(ship.x, ship.y, a.x, a.y) <= r * r) {
                    killShip();
                    break;
                }
            }
        }

        // Asteroids vs asteroids — an elastic bounce so the field feels like real physics rather
        // than a pile of ghosts drifting through each other.
        for (int i = 0; i < MAX_ASTEROIDS; i++) {
            Asteroid a = asteroids[i];
            if (!a.active) continue;
            for (int j = i + 1; j < MAX_ASTEROIDS; j++) {
                Asteroid b = asteroids[j];
                if (!b.active) continue;
                resolveAsteroidPair(a, b);
            }
        }
    }

    // resolveAsteroidPair separates overlapping asteroids and reflects their velocities along
    // the collision normal, with mass proportional to radius^2 (2D area) — a standard elastic
    // collision impulse, not a gameplay mechanic borrowed from the original Asteroids.
    private static void resolveAsteroidPair(Asteroid a, Asteroid b) {
        double dx = a.x - b.x;
        double dy = a.y - b.y;
        double d2 = dx * dx + dy * dy;
        double minDist = a.radius + b.radius;
        if (d2 >= minDist * minDist || d2 < 0.0001) return;

        double dist = Math.sqrt(d2);
        double nx = dx / dist;
        double ny = dy / dist;
        double m1 = a.radius * a.radius;
        double m2 = b.radius * b.radius;
        double total = m1 + m2;

        double overlap = minDist - dist;
        a.x += nx * (overlap * (m2 / total));
        a.y += ny * (overlap * (m2 / total));
        b.x -= nx * (overlap * (m1 / total));
        b.y -= ny * (overlap * (m1 / total));

        double relX = a.vx - b.vx;
        double relY = a.vy - b.vy;
        double vn = relX * nx + relY * ny;
        if (vn >= 0) return; // already separating

        double imp = (2 * vn) / total;
        a.vx -= nx * (imp * m2);
        a.vy -= ny * (imp * m2);
        b.vx += nx * (imp * m1);
        b.vy += ny * (imp * m1);

        if (thudCooldown <= 0) {
            Env.playSound(SND_THUD);
            thudCooldown = 0.12;
        }
    }

    private static void killShip() {
        Env.playSound(SND_SHIP_DEATH);
        spawnParticles(ship.x, ship.y, 28, 260);
        spawnDebris(ship.x, ship.y, 10, 300, 1.4);
        spawnShockwave(ship.x, ship.y, 70);
        triggerShake(14);
        ship.alive = false;
        lives -= 1;
        if (lives <= 0) {
            gameOver = true;
            if (score > bestScore) {
                bestScore = score;
                setBestScore(bestScore);
            }
        } else {
            ship.respawn = RESPAWN_DELAY;
        }
    }

    private static double dist2(double ax, double ay, double bx, double by) {
        double dx = ax - bx;
        double dy = ay - by;
        return dx * dx + dy * dy;
    }

    // ------------------------------------------------------------------------------------------
    // Small helpers
    private static double frange(double lo, double hi) {
        return lo + Math.random() * (hi - lo);
    }

    private static int getBestScore() {
        Storage storage = Storage.getLocalStorage();
        if (storage == null) return 0;
        String v = storage.getItem("jasteroids_best_score");
        if (v == null) return 0;
        try {
            return Integer.parseInt(v);
        } catch (NumberFormatException e) {
            return 0;
        }
    }

    private static void setBestScore(int value) {
        Storage storage = Storage.getLocalStorage();
        if (storage != null) {
            storage.setItem("jasteroids_best_score", Integer.toString(value));
        }
    }
}
