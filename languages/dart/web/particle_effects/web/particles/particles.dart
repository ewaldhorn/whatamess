import 'package:web/web.dart' as web;
import 'dart:js_interop';

import '../utils/canvasUtils.dart';

// ------------------------------------------------------------- class Particle
class Particle {
  Particle();

  draw(web.CanvasRenderingContext2D ctx) {}

  update() {}
}

// --------------------------------------------------------------- class Effect
class Effect {
  web.CanvasRenderingContext2D ctx;
  int canvasWidth, canvasHeight;
  late web.CanvasGradient gradient;

  // ------------------------------------------------------------------- Effect
  Effect(this.ctx, this.canvasWidth, this.canvasHeight) {
    gradient = createTextGradient(ctx, (w: canvasWidth, h: canvasHeight));
  }

  // ----------------------------------------------------------------- wrapText
  wrapText(String text) {
    drawWrappedText(ctx, text, (x: canvasWidth ~/ 2, y: canvasHeight ~/ 2),
        fillStyle: gradient);
  }

  // ------------------------------------------------------- convertToParticles
  convertToParticles() {}

  // ------------------------------------------------------------------- render
  render() {}
}
