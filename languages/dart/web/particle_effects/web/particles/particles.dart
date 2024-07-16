import 'dart:math';

import 'package:web/web.dart' as web;
import 'dart:js_interop';

import '../utils/canvasUtils.dart';

final Random r = Random();

// ------------------------------------------------------------- class Particle
class Particle {
  Effect effect;
  int x, y, originX, originY, size;
  JSString color;
  int dx = 0, dy = 0, vx = 0, vy = 0, force = 0, angle = 0, distance = 0;
  double friction = r.nextDouble() * 0.6 + 0.15;
  double ease = r.nextDouble() * 0.1 + 0.005;

  Particle(this.effect, this.x, this.y, this.color, this.size)
      : originX = x,
        originY = y;

  draw() {
    effect.ctx.fillStyle = color;
    effect.ctx.fillRect(x, y, size, size);
  }

  update() {}
}

// --------------------------------------------------------------- class Effect
class Effect {
  web.CanvasRenderingContext2D ctx;
  int canvasWidth, canvasHeight;
  late web.CanvasGradient gradient;
  late ({int x, int y}) mousePosition;
  List<Particle> particles = [];

  final int gap = 3;

  // ------------------------------------------------------------------- Effect
  Effect(this.ctx, this.canvasWidth, this.canvasHeight) {
    gradient = createTextGradient(ctx, (w: canvasWidth, h: canvasHeight));
    mousePosition = (x: 0, y: 0);
    web.window.addEventListener('mousemove', mouseMovementListener.toJS);
  }

  // ---------------------------------------------------- mouseMovementListener
  mouseMovementListener(web.MouseEvent something) {
    mousePosition =
        (x: something.offsetX.toInt(), y: something.offsetY.toInt());
  }

  // ----------------------------------------------------------------- wrapText
  wrapText(String text) {
    drawWrappedText(ctx, text, (x: canvasWidth ~/ 2, y: canvasHeight ~/ 2),
        fillStyle: gradient);
    convertToParticles();
  }

  // ------------------------------------------------------- convertToParticles
  convertToParticles() {
    particles.clear();
    var imageData =
        ctx.getImageData(50, 50, canvasWidth - 50, canvasHeight - 50);
    final localImageData = imageData.data.toDart;

    for (int y = 0; y < imageData.height; y += gap) {
      for (int x = 0; x < imageData.width; x += gap) {
        final idx = ((y * imageData.width) + x) * 4; // 4 entries per pixel
        final alpha = localImageData[idx + 3];
        if (alpha > 0) {
          final int red = localImageData[idx];
          final int green = localImageData[idx + 1];
          final int blue = localImageData[idx + 2];
          final color = 'rgb($red,$green,$blue)'.toJS;
          particles.add(Particle(this, x, y, color, gap));
        }
      }
    }
  }

  // ------------------------------------------------------------------- render
  render() {
    for (var p in particles) {
      p.update();
      p.draw();
    }
  }
}
