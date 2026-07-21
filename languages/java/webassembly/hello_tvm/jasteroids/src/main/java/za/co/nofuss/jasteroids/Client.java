package za.co.nofuss.jasteroids;

import org.teavm.jso.browser.AnimationFrameCallback;
import org.teavm.jso.browser.Performance;
import org.teavm.jso.browser.Window;
import org.teavm.jso.canvas.CanvasRenderingContext2D;
import org.teavm.jso.dom.events.KeyboardEvent;
import org.teavm.jso.dom.html.HTMLCanvasElement;
import org.teavm.jso.dom.html.HTMLDocument;
import org.teavm.jso.dom.html.HTMLElement;

/**
 * Client.java — Destroids bootstrap: canvas setup, the per-frame animation callback, keyboard
 * input, and the FPS readout, ported from the Odin original's main.odin. The game itself lives
 * in {@link Game} (simulation) and {@link Draw} (rendering).
 */
public class Client {
    private static HTMLCanvasElement canvas;
    private static CanvasRenderingContext2D ctx;
    private static HTMLElement fpsEl;

    private static double lastMs;
    private static int fpsFrames;
    private static double fpsAccum;

    public static void main(String[] args) {
        HTMLDocument document = HTMLDocument.current();
        canvas = (HTMLCanvasElement) document.getElementById("game");
        canvas.setWidth(Game.CANVAS_W);
        canvas.setHeight(Game.CANVAS_H);
        ctx = (CanvasRenderingContext2D) canvas.getContext("2d");
        fpsEl = document.getElementById("fps");

        Game.init();
        bindKeys();

        lastMs = Performance.now();
        Window.requestAnimationFrame(new AnimationLoop());
    }

    private static class AnimationLoop implements AnimationFrameCallback {
        @Override
        public void onAnimationFrame(double timestamp) {
            double now = Performance.now();
            double dt = (now - lastMs) / 1000.0;
            lastMs = now;
            if (dt > 0.05) dt = 0.05; // clamp after tab-out / long stalls
            if (dt < 0) dt = 0;

            Game.update(dt);
            Draw.draw(ctx);

            // FPS: average instantaneous rate over ~0.4s, then push to the footer.
            if (dt > 0) {
                fpsAccum += 1.0 / dt;
                fpsFrames += 1;
                if (fpsFrames >= 24) {
                    fpsEl.setInnerText("FPS: " + Math.round(fpsAccum / fpsFrames));
                    fpsAccum = 0;
                    fpsFrames = 0;
                }
            }

            Window.requestAnimationFrame(this);
        }
    }

    // Maps browser KeyboardEvent.code strings onto Game.KEY_* slots — the Java-side equivalent
    // of the KEYMAP object + destroids_key export in the Odin original's index.html.
    private static int mapKey(String code) {
        switch (code) {
            case "ArrowLeft":
            case "KeyA":
                return Game.KEY_LEFT;
            case "ArrowRight":
            case "KeyD":
                return Game.KEY_RIGHT;
            case "ArrowUp":
            case "KeyW":
                return Game.KEY_THRUST;
            case "Space":
                return Game.KEY_FIRE;
            case "Enter":
                return Game.KEY_START;
            case "ShiftLeft":
            case "ShiftRight":
                return Game.KEY_HYPER;
            case "KeyP":
                return Game.KEY_PAUSE;
            default:
                return -1;
        }
    }

    private static void bindKeys() {
        Window.current().onKeyDown((KeyboardEvent e) -> {
            int slot = mapKey(e.getCode());
            if (slot < 0) return;
            e.preventDefault();
            if (!e.isRepeat()) {
                Game.keyEvent(slot, true);
            }
        });
        Window.current().onKeyUp((KeyboardEvent e) -> {
            int slot = mapKey(e.getCode());
            if (slot < 0) return;
            e.preventDefault();
            Game.keyEvent(slot, false);
        });
    }
}
