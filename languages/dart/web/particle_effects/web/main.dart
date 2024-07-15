import 'package:web/web.dart' as web;

import 'canvasUtils.dart';
import 'domUtils.dart';
import 'headerManagement.dart';

late web.HTMLCanvasElement mainCanvas;
late web.CanvasRenderingContext2D context2d;

// ------------------------------------------------------------- setupVariables
void setupVariables() {
  mainCanvas =
      web.document.getElementById('mainCanvas') as web.HTMLCanvasElement;
  context2d = mainCanvas.context2D;
}

// ----------------------------------------------------------------------- main
void main() {
  setHeading();
  insertHRBefore('output');
  setupVariables();
  fillCanvasWithPyramid(mainCanvas, context2d);
}
