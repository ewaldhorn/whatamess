import 'package:web/web.dart' as web;

late web.HTMLCanvasElement mainCanvas;
late web.CanvasRenderingContext2D context2d;

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
  (int, int) startXY,
  (int, int) sizeWH,
) {
  final (x, y) = startXY;
  final (w, h) = sizeWH;

  ctx.beginPath();
  ctx
    ..moveTo(x, y)
    ..lineTo(x + w, y)
    ..lineTo(x + w, y + h)
    ..lineTo(x, y + h)
    ..lineTo(x, y)
    ..stroke();
}

// ---- fillCanvasWithPyramid
void fillCanvasWithPyramid(
    web.HTMLCanvasElement canvas, web.CanvasRenderingContext2D ctx) {
  final midX = canvas.width ~/ 2;
  final midY = canvas.height ~/ 2;

  for (int i = 1; i < midX; i += 3) {
    drawSquare(ctx, (midX - i, midY - i), (i * 2, i * 2));
  }
}

// ----------------------------------------------------------------------- main
void main() {
  setHeading();
  insertHRBefore('output');
  setupVariables();
  fillCanvasWithPyramid(mainCanvas, context2d);
}
