import 'package:web/web.dart' as web;

import 'utils/canvasUtils.dart';
import 'utils/domUtils.dart';
import 'utils/headerManagement.dart';

late web.HTMLCanvasElement mainCanvas;
late web.CanvasRenderingContext2D context2d;
late web.HTMLInputElement inputField;

int score = 0;
int gameFrame = 0;

// ------------------------------------------------------------- setupVariables
void setupVariables() {
  mainCanvas =
      web.document.getElementById('mainCanvas') as web.HTMLCanvasElement;

  var availableHeight = web.document.body?.clientHeight;

  if (availableHeight != null) {
    availableHeight = availableHeight - 200;
  }

  mainCanvas.width = web.document.body?.clientWidth ?? 800;
  mainCanvas.height = availableHeight ?? 800;

  context2d = mainCanvas.context2D;
  context2d.font = '50px Georgia';

  inputField = web.document.getElementById('textInput') as web.HTMLInputElement;
  inputField.focus();
}

// ----------------------------------------------------------------------- main
void main() {
  setHeading();
  insertHRBefore('output');
  setupVariables();
  fillCanvasWithPyramid(mainCanvas, context2d);
}
