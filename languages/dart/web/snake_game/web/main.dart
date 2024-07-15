import 'dart:js_interop';

import 'package:web/web.dart' as web;

import 'keyboard.dart';
import 'snake.dart';

// ----------------------------------------------------------- Global Constants
const int cellSize = 10;
const int initialLength = 3;

// ----------------------------------------------------------- Global Variables
late web.HTMLCanvasElement canvas;
late web.CanvasRenderingContext2D ctx;
late Keyboard keyboard;
late Snake snake;

// ------------------------------------------------------------------- initGame
void initGame() {
  canvas = web.document.querySelector('#canvas') as web.HTMLCanvasElement;
  ctx = canvas.getContext('2d') as web.CanvasRenderingContext2D;
  keyboard = Keyboard();
  snake = Snake(initialLength);
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

// ---------------------------------------------------------------- clearScreen
void clearScreen() {
  ctx
    ..fillStyle = 'white'.toJS
    ..fillRect(0, 0, canvas.width, canvas.height);
}

// ----------------------------------------------------------------------- main
void main() {
  initGame();

  snake.reportDirection();
  snake.changeDirection(Snake.down);
  snake.reportDirection();

  drawCell((x: 10, y: 10), "salmon");
}
