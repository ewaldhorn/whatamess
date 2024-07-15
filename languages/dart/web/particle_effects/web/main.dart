import 'package:web/web.dart' as web;

late web.HTMLCanvasElement mainCanvas;
late web.CanvasRenderingContext2D context2d;

// ---------------------------------------------------------------------- Types
typedef Position = ({int x, int y});
typedef Size = ({int w, int h});

// ----------------------------------------------------------------- setHeading
void setHeading() {
  final headingElement =
      web.document.getElementById('heading') as web.HTMLHeadingElement;
  headingElement.text = 'Dart-to-JS Experiments';
}

// ------------------------------------------------------------- insertHRBefore
void insertHRBefore(String element) {
  final hrElement = web.HTMLHRElement();
  final outputElement =
      web.document.getElementById(element) as web.HTMLDivElement;
  web.document.body?.insertBefore(hrElement, outputElement);
}

// ------------------------------------------------------------- setupVariables
void setupVariables() {
  mainCanvas =
      web.document.getElementById('mainCanvas') as web.HTMLCanvasElement;
  context2d = mainCanvas.context2D;
}

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

// ---- fillCanvasWithPyramid
void fillCanvasWithPyramid(
    web.HTMLCanvasElement canvas, web.CanvasRenderingContext2D ctx) {
  final midX = canvas.width ~/ 2;
  final midY = canvas.height ~/ 2;

  for (int i = 1; i < midX; i += 3) {
    drawSquare(ctx, (x: midX - i, y: midY - i), (w: i * 2, h: i * 2));
  }
}

// ----------------------------------------------------------------------- main
void main() {
  setHeading();
  insertHRBefore('output');
  setupVariables();
  fillCanvasWithPyramid(mainCanvas, context2d);
}
