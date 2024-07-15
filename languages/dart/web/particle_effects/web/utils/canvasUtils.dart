import 'package:web/web.dart' as web;
import '../types.dart';

// ----------------------------------------------------------------- drawSquare
void drawSquare(
  web.CanvasRenderingContext2D ctx,
  Position start,
  Size size,
) {
  ctx.beginPath();
  ctx
    ..moveTo(start.x, start.y)
    ..lineTo(start.x + size.w, start.y)
    ..lineTo(start.x + size.w, start.y + size.h)
    ..lineTo(start.x, start.y + size.h)
    ..lineTo(start.x, start.y)
    ..stroke();
}

// ------------------------------------------------------ fillCanvasWithPyramid
void fillCanvasWithPyramid(
    web.HTMLCanvasElement canvas, web.CanvasRenderingContext2D ctx) {
  final midX = canvas.width ~/ 2;
  final midY = canvas.height ~/ 2;

  for (int i = 1; i < midX; i += 3) {
    drawSquare(ctx, (x: midX - i, y: midY - i), (w: i * 2, h: i * 2));
  }
}
