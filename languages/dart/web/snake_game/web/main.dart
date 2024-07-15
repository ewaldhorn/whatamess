import 'package:web/web.dart' as web;

late web.HTMLCanvasElement canvas;
late web.CanvasRenderingContext2D ctx;

// ------------------------------------------------------------------- initGame
void initGame() {
  canvas = web.document.querySelector('#canvas') as web.HTMLCanvasElement;
  ctx = canvas.getContext('2d') as web.CanvasRenderingContext2D;
}

// ----------------------------------------------------------------------- main
void main() {
  initGame();
}
