import 'dart:js_interop';

import 'package:web/web.dart' as web;

// ----------------------------------------------------------- Global Constants
const int cellSize = 10;

// ----------------------------------------------------------- Global Variables
late web.HTMLCanvasElement canvas;
late web.CanvasRenderingContext2D ctx;

// ------------------------------------------------------------------- initGame
void initGame() {
  canvas = web.document.querySelector('#canvas') as web.HTMLCanvasElement;
  ctx = canvas.getContext('2d') as web.CanvasRenderingContext2D;
}

// ------------------------------------------------------------------- drawCell
void drawCell(({int x, int y}) coords, String color) {
  ctx
    ..fillStyle = color.toJS
    ..strokeStyle = 'white'.toJS;

  final num x = coords.x * cellSize;
  final num y = coords.y * cellSize;

  ctx
    ..fillRect(x, y, cellSize, cellSize)
    ..strokeRect(x, y, cellSize, cellSize);
}

// ----------------------------------------------------------------------- main
void main() {
  initGame();

  drawCell((x: 10, y: 10), "salmon");
}
