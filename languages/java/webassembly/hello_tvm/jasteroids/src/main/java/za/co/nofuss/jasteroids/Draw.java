package za.co.nofuss.jasteroids;

import org.teavm.jso.canvas.CanvasRenderingContext2D;

/**
 * Draw.java — vector rendering for Destroids, ported from the Odin original's draw.odin.
 * The Odin version batches every draw op into one shared buffer and flushes it in a single
 * WASM&lt;-&gt;JS crossing to keep hundreds of ops/frame cheap; here each op is just a direct JSO
 * call onto {@link CanvasRenderingContext2D} — TeaVM's WASM-GC/JSO crossing cost model doesn't
 * carry the same per-call overhead that motivated batching in the Odin/Batchiness original.
 */
final class Draw {
    private Draw() {
    }

    private static double animTime; // drives star twinkle; advanced each drawn frame

    static void draw(CanvasRenderingContext2D ctx) {
        animTime += 1.0 / 60.0;

        // Background.
        ctx.setFillStyle("#05070d");
        ctx.fillRect(0, 0, Game.CANVAS_W, Game.CANVAS_H);

        // The world (everything but the HUD) jitters with screen shake; the HUD stays put so the
        // score/wave/lives readout never becomes unreadable mid-shake.
        ctx.save();
        ctx.translate(Game.shakeX, Game.shakeY);

        drawStars(ctx);
        drawShockwaves(ctx);
        drawParticles(ctx);
        drawDebris(ctx);
        drawAsteroids(ctx);
        drawBullets(ctx);
        if (Game.ship.alive) {
            drawShip(ctx);
        }

        ctx.restore();

        drawHud(ctx);

        if (Game.intro) {
            drawIntro(ctx);
        } else if (Game.paused) {
            drawPaused(ctx);
        } else if (Game.gameOver) {
            drawGameOver(ctx);
        }
    }

    // --------------------------------------------------------------------------------------
    // Hundreds of tiny points — the bulk of the per-frame draw-call count.
    private static void drawStars(CanvasRenderingContext2D ctx) {
        for (Game.Star s : Game.stars) {
            double tw = 0.55 + 0.45 * Math.sin(animTime * 2.0 + s.twinkle);
            double a = s.brightness * tw;
            ctx.setFillStyle(grey(a));
            double sz = s.brightness > 0.7 ? 2 : 1;
            ctx.fillRect(s.x, s.y, sz, sz);
        }
    }

    private static void drawParticles(CanvasRenderingContext2D ctx) {
        ctx.setShadowColor("#ffce7a");
        ctx.setShadowBlur(4);
        ctx.setFillStyle("#ffce7a");
        for (Game.Particle p : Game.particles) {
            if (!p.active) continue;
            double a = p.life / p.maxLife;
            ctx.setGlobalAlpha(a);
            ctx.fillRect(p.x - 1, p.y - 1, 2, 2);
        }
        ctx.setGlobalAlpha(1);
        ctx.setShadowBlur(0);
    }

    // Spinning line-shards flung out of destroyed ships/asteroids.
    private static void drawDebris(CanvasRenderingContext2D ctx) {
        ctx.setStrokeStyle("#e8c79a");
        ctx.setLineWidth(1.4);
        for (Game.Debris d : Game.debris) {
            if (!d.active) continue;
            double a = d.life / d.maxLife;
            ctx.setGlobalAlpha(a);
            double c = Math.cos(d.angle);
            double s = Math.sin(d.angle);
            double hx = c * d.len * 0.5;
            double hy = s * d.len * 0.5;
            ctx.beginPath();
            ctx.moveTo(d.x - hx, d.y - hy);
            ctx.lineTo(d.x + hx, d.y + hy);
            ctx.stroke();
        }
        ctx.setGlobalAlpha(1);
    }

    // The expanding, fading rings spawned at explosion origins.
    private static void drawShockwaves(CanvasRenderingContext2D ctx) {
        ctx.setStrokeStyle("#8fd8ff");
        for (Game.Shockwave s : Game.shockwaves) {
            if (!s.active) continue;
            double a = s.life / s.maxLife;
            ctx.setGlobalAlpha(a * 0.8);
            ctx.setLineWidth(2.0 + 3.0 * (1 - a));
            ctx.beginPath();
            ctx.arc(s.x, s.y, s.radius, 0, Game.TAU);
            ctx.stroke();
        }
        ctx.setGlobalAlpha(1);
    }

    private static void drawAsteroids(CanvasRenderingContext2D ctx) {
        ctx.setStrokeStyle("#c9d4e6");
        ctx.setLineWidth(1.6);
        for (Game.Asteroid a : Game.asteroids) {
            if (!a.active) continue;
            ctx.beginPath();
            for (int i = 0; i < Game.SHAPE_N; i++) {
                double ang = a.angle + Game.TAU * i / Game.SHAPE_N;
                double r = a.radius * a.shape[i];
                double x = a.x + Math.cos(ang) * r;
                double y = a.y + Math.sin(ang) * r;
                if (i == 0) {
                    ctx.moveTo(x, y);
                } else {
                    ctx.lineTo(x, y);
                }
            }
            ctx.closePath();
            ctx.stroke();
        }
    }

    // Each shot is a short glowing streak stretched back along its velocity, instead of a
    // static dot, so motion reads clearly even at high bullet speed.
    private static void drawBullets(CanvasRenderingContext2D ctx) {
        ctx.setShadowColor("#ffffff");
        ctx.setShadowBlur(6);
        ctx.setStrokeStyle("#ffffff");
        ctx.setLineWidth(2.2);
        ctx.setLineCap("round");
        for (Game.Bullet b : Game.bullets) {
            if (!b.active) continue;
            double tailX = b.x - b.vx * 0.02;
            double tailY = b.y - b.vy * 0.02;
            ctx.beginPath();
            ctx.moveTo(tailX, tailY);
            ctx.lineTo(b.x, b.y);
            ctx.stroke();
        }
        ctx.setShadowBlur(0);
    }

    private static void drawShip(CanvasRenderingContext2D ctx) {
        Game.Ship ship = Game.ship;

        // Blink while spawn-protected. Suppressed during intro/pause — invuln doesn't count down
        // while frozen, so the blink phase would otherwise be stuck and could hide the ship.
        if (ship.invuln > 0 && !Game.intro && !Game.paused && ((int) (ship.invuln * 12)) % 2 == 0) {
            return;
        }

        double c = Math.cos(ship.angle);
        double s = Math.sin(ship.angle);
        // Local-space triangle: nose forward (+x), two tail corners.
        double noseX = 16, noseY = 0;
        double tailLX = -10, tailLY = -9;
        double tailRX = -10, tailRY = 9;

        double p0x = ship.x + (noseX * c - noseY * s);
        double p0y = ship.y + (noseX * s + noseY * c);
        double p1x = ship.x + (tailLX * c - tailLY * s);
        double p1y = ship.y + (tailLX * s + tailLY * c);
        double p2x = ship.x + (tailRX * c - tailRY * s);
        double p2y = ship.y + (tailRX * s + tailRY * c);

        ctx.setStrokeStyle("#e8f0ff");
        ctx.setLineWidth(2.0);
        ctx.beginPath();
        ctx.moveTo(p0x, p0y);
        ctx.lineTo(p1x, p1y);
        ctx.lineTo(p2x, p2y);
        ctx.closePath();
        ctx.stroke();

        // Thrust flame — a flickering triangle out the back.
        if (ship.thrust && !Game.paused && !Game.intro && ((int) (animTime * 30)) % 2 == 0) {
            double flLX = -10, flLY = -5;
            double flRX = -10, flRY = 5;
            double tipX = -20, tipY = 0;
            double flx = ship.x + (flLX * c - flLY * s);
            double fly = ship.y + (flLX * s + flLY * c);
            double frx = ship.x + (flRX * c - flRY * s);
            double fry = ship.y + (flRX * s + flRY * c);
            double tpx = ship.x + (tipX * c - tipY * s);
            double tpy = ship.y + (tipX * s + tipY * c);
            ctx.setStrokeStyle("#ff9a3c");
            ctx.beginPath();
            ctx.moveTo(flx, fly);
            ctx.lineTo(tpx, tpy);
            ctx.lineTo(frx, fry);
            ctx.stroke();
        }
    }

    // --------------------------------------------------------------------------------------
    // HUD
    private static void drawHud(CanvasRenderingContext2D ctx) {
        ctx.setFillStyle("#e8f0ff");
        ctx.setFont("20px monospace");
        ctx.setTextAlign("left");
        ctx.setTextBaseline("top");
        ctx.fillText("SCORE " + Game.score, 16, 14);

        ctx.setTextAlign("right");
        ctx.fillText("WAVE " + Game.level, Game.CANVAS_W - 16, 14);

        // Lives as little ship glyphs, top-center-ish.
        ctx.setStrokeStyle("#e8f0ff");
        ctx.setLineWidth(2.0);
        for (int i = 0; i < Game.lives; i++) {
            double x = Game.CANVAS_W / 2.0 - Game.lives * 12 + i * 24;
            drawShipGlyph(ctx, x, 26);
        }
    }

    private static void drawShipGlyph(CanvasRenderingContext2D ctx, double x, double y) {
        ctx.beginPath();
        ctx.moveTo(x, y - 10);
        ctx.lineTo(x - 7, y + 8);
        ctx.lineTo(x + 7, y + 8);
        ctx.closePath();
        ctx.stroke();
    }

    private static void drawGameOver(CanvasRenderingContext2D ctx) {
        ctx.setFillStyle("rgba(3,5,10,0.72)");
        ctx.fillRect(0, 0, Game.CANVAS_W, Game.CANVAS_H);

        ctx.setFillStyle("#ff5a4c");
        ctx.setFont("bold 64px monospace");
        ctx.setTextAlign("center");
        ctx.setTextBaseline("middle");
        ctx.fillText("GAME OVER", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 - 40);

        ctx.setFillStyle("#e8f0ff");
        ctx.setFont("24px monospace");
        ctx.fillText("FINAL SCORE " + Game.score, Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 + 16);

        ctx.setFillStyle("#ffce7a");
        ctx.setFont("18px monospace");
        ctx.fillText("BEST SCORE " + Game.bestScore, Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 + 46);

        ctx.setFillStyle("#9fb2d0");
        ctx.setFont("18px monospace");
        ctx.fillText("PRESS ENTER TO PLAY AGAIN", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 + 78);
    }

    // The one-time splash shown from boot until the first ENTER/SPACE press — this game is
    // keyboard-driven and unplayable on touch, so it says so up front rather than confusing
    // anyone.
    private static void drawIntro(CanvasRenderingContext2D ctx) {
        ctx.setFillStyle("rgba(3,5,10,0.78)");
        ctx.fillRect(0, 0, Game.CANVAS_W, Game.CANVAS_H);

        ctx.setTextAlign("center");
        ctx.setTextBaseline("middle");

        ctx.setFillStyle("#e8f0ff");
        ctx.setFont("bold 40px monospace");
        ctx.fillText("JASTEROIDS", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 - 84);

        ctx.setFillStyle("#ffce7a");
        ctx.setFont("18px monospace");
        ctx.fillText("BUILT FOR DESKTOP — KEYBOARD REQUIRED", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 - 36);

        ctx.setFillStyle("#9fb2d0");
        ctx.setFont("16px monospace");
        ctx.fillText("◀ ▶ TURN   ▲ THRUST   SPACE FIRE   SHIFT HYPERSPACE", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0);
        ctx.fillText("P PAUSE   ENTER RESTART", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 + 26);

        ctx.setFillStyle("#7fe0a0");
        ctx.setFont("bold 20px monospace");
        ctx.fillText("PRESS ENTER OR SPACE TO START", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 + 68);
    }

    // The overlay for a mid-game P toggle — the frozen field stays visible underneath.
    private static void drawPaused(CanvasRenderingContext2D ctx) {
        ctx.setFillStyle("rgba(3,5,10,0.6)");
        ctx.fillRect(0, 0, Game.CANVAS_W, Game.CANVAS_H);

        ctx.setFillStyle("#e8f0ff");
        ctx.setFont("bold 40px monospace");
        ctx.setTextAlign("center");
        ctx.setTextBaseline("middle");
        ctx.fillText("PAUSED", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 - 10);

        ctx.setFillStyle("#9fb2d0");
        ctx.setFont("18px monospace");
        ctx.fillText("PRESS P TO RESUME", Game.CANVAS_W / 2.0, Game.CANVAS_H / 2.0 + 26);
    }

    // --------------------------------------------------------------------------------------
    private static String grey(double v) {
        int iv = (int) (clamp01(v) * 255);
        return "rgb(" + iv + "," + iv + "," + iv + ")";
    }

    private static double clamp01(double v) {
        if (v < 0) return 0;
        if (v > 1) return 1;
        return v;
    }
}
