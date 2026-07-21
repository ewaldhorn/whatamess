package za.co.nofuss.hello;

import org.teavm.jso.browser.Window;
import org.teavm.jso.canvas.CanvasRenderingContext2D;
import org.teavm.jso.dom.html.HTMLCanvasElement;
import org.teavm.jso.dom.html.HTMLDocument;

public class Client {
    private static final int DUCK_COUNT = 20;

    private static HTMLCanvasElement canvas;
    private static CanvasRenderingContext2D ctx;
    private static Duck[] ducks;

    // pond ellipse, in fractions of canvas size
    private static double pondCx;
    private static double pondCy;
    private static double pondRx;
    private static double pondRy;

    public static void main(String[] args) {
        HTMLDocument document = HTMLDocument.current();
        canvas = (HTMLCanvasElement) document.getElementById("pond");
        ctx = (CanvasRenderingContext2D) canvas.getContext("2d");

        resize();
        Window.current().addEventListener("resize", evt -> resize());

        ducks = new Duck[DUCK_COUNT];
        for (int i = 0; i < DUCK_COUNT; i++) {
            ducks[i] = new Duck();
        }

        Window.requestAnimationFrame(new AnimationLoop());
    }

    private static void resize() {
        int w = (int) (Window.current().getInnerWidth() * 0.9);
        int h = (int) (Window.current().getInnerHeight() * 0.9);
        canvas.setWidth(w);
        canvas.setHeight(h);

        pondCx = w * 0.5;
        pondCy = h * 0.58;
        pondRx = w * 0.38;
        pondRy = h * 0.32;
    }

    private static class AnimationLoop implements org.teavm.jso.browser.AnimationFrameCallback {
        @Override
        public void onAnimationFrame(double timestamp) {
            update(timestamp);
            draw(timestamp);
            Window.requestAnimationFrame(this);
        }
    }

    private static void update(double timestamp) {
        for (Duck duck : ducks) {
            duck.update();
        }
    }

    private static void draw(double timestamp) {
        int w = canvas.getWidth();
        int h = canvas.getHeight();

        // sky
        ctx.setFillStyle("#8ecae6");
        ctx.fillRect(0, 0, w, h);

        // grass
        ctx.setFillStyle("#5a9c3a");
        ctx.fillRect(0, h * 0.35, w, h * 0.65);

        // trees
        drawTree(w * 0.08, h * 0.38, 0.9);
        drawTree(w * 0.20, h * 0.30, 0.6);
        drawTree(w * 0.90, h * 0.36, 1.0);
        drawTree(w * 0.80, h * 0.28, 0.55);

        // pond
        ctx.beginPath();
        ctx.setFillStyle("#3a7ca8");
        ctx.ellipse(pondCx, pondCy, pondRx, pondRy, 0, 0, Math.PI * 2);
        ctx.fill();

        // pond shore ring
        ctx.beginPath();
        ctx.setLineWidth(6);
        ctx.setStrokeStyle("#4a8c2a");
        ctx.ellipse(pondCx, pondCy, pondRx + 3, pondRy + 3, 0, 0, Math.PI * 2);
        ctx.stroke();

        // ducks
        for (Duck duck : ducks) {
            drawDuck(duck);
        }
    }

    private static void drawTree(double x, double y, double scale) {
        double trunkW = 12 * scale;
        double trunkH = 40 * scale;

        ctx.setFillStyle("#6b4423");
        ctx.fillRect(x - trunkW / 2, y - trunkH, trunkW, trunkH);

        ctx.beginPath();
        ctx.setFillStyle("#2f6b2f");
        ctx.arc(x, y - trunkH - 10 * scale, 30 * scale, 0, Math.PI * 2);
        ctx.fill();

        ctx.beginPath();
        ctx.setFillStyle("#3a7d3a");
        ctx.arc(x - 18 * scale, y - trunkH, 22 * scale, 0, Math.PI * 2);
        ctx.fill();

        ctx.beginPath();
        ctx.setFillStyle("#3a7d3a");
        ctx.arc(x + 18 * scale, y - trunkH, 22 * scale, 0, Math.PI * 2);
        ctx.fill();
    }

    private static void drawDuck(Duck duck) {
        double bob = Math.sin(duck.bobPhase) * 2;

        ctx.save();
        ctx.translate(duck.x, duck.y + bob);
        ctx.rotate(duck.angle);

        // ripple
        ctx.beginPath();
        ctx.setStrokeStyle("rgba(255,255,255,0.4)");
        ctx.setLineWidth(1);
        ctx.ellipse(0, 0, 18, 8, 0, 0, Math.PI * 2);
        ctx.stroke();

        // tail, pointed up like a real duck's tail feathers
        ctx.beginPath();
        ctx.setFillStyle("#f2c200");
        ctx.moveTo(-11, 2);
        ctx.quadraticCurveTo(-19, -1, -15, -8);
        ctx.quadraticCurveTo(-13, -2, -9, -3);
        ctx.closePath();
        ctx.fill();

        // body, plump oval
        ctx.beginPath();
        ctx.setFillStyle("#ffd400");
        ctx.ellipse(-1, 1, 13, 9, 0, 0, Math.PI * 2);
        ctx.fill();

        // folded wing, curved feather crease
        ctx.beginPath();
        ctx.setStrokeStyle("#d99a00");
        ctx.setLineWidth(1.5);
        ctx.arc(-3, 1, 7, -2.3, -0.4);
        ctx.stroke();
        ctx.beginPath();
        ctx.setFillStyle("#f5c800");
        ctx.ellipse(-2, 3, 8, 5, 0.2, 0, Math.PI * 2);
        ctx.fill();

        // neck/head, sitting high and close to the body
        ctx.beginPath();
        ctx.setFillStyle("#ffd400");
        ctx.arc(11, -6, 6.5, 0, Math.PI * 2);
        ctx.fill();

        // flat, wide duck bill
        ctx.beginPath();
        ctx.setFillStyle("#ff9900");
        ctx.ellipse(19, -6, 6.5, 3.2, 0, 0, Math.PI * 2);
        ctx.fill();
        ctx.beginPath();
        ctx.setStrokeStyle("#cc7a00");
        ctx.setLineWidth(0.8);
        ctx.moveTo(14, -6);
        ctx.lineTo(25, -6);
        ctx.stroke();

        // nostril
        ctx.beginPath();
        ctx.setFillStyle("#a35a00");
        ctx.arc(17, -6.8, 0.5, 0, Math.PI * 2);
        ctx.fill();

        // eye
        ctx.beginPath();
        ctx.setFillStyle("#000000");
        ctx.arc(12.5, -8, 1.1, 0, Math.PI * 2);
        ctx.fill();

        ctx.restore();
    }

    private static class Duck {
        double x;
        double y;
        double angle;
        double speed;
        double bobPhase;

        Duck() {
            double t = Math.random() * Math.PI * 2;
            double r = Math.random();
            x = pondCx + Math.cos(t) * pondRx * 0.6 * r;
            y = pondCy + Math.sin(t) * pondRy * 0.6 * r;
            angle = Math.random() * Math.PI * 2;
            speed = 0.3 + Math.random() * 0.4;
            bobPhase = Math.random() * Math.PI * 2;
        }

        void update() {
            angle += (Math.random() - 0.5) * 0.08;
            bobPhase += 0.05;

            double nx = x + Math.cos(angle) * speed;
            double ny = y + Math.sin(angle) * speed;

            double dx = (nx - pondCx) / pondRx;
            double dy = (ny - pondCy) / pondRy;
            if (dx * dx + dy * dy > 0.85) {
                angle += Math.PI + (Math.random() - 0.5) * 0.5;
            } else {
                x = nx;
                y = ny;
            }
        }
    }
}
