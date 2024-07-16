import 'package:web/web.dart' as web;
import 'dart:js_interop' as js;

import 'utils/canvasUtils.dart';
import 'utils/domUtils.dart';
import 'utils/headerManagement.dart';
import 'particles/particles.dart';

late web.HTMLCanvasElement mainCanvas;
late web.CanvasRenderingContext2D context2d;
late web.HTMLInputElement inputField;

int score = 0;
int gameFrame = 0;

Effect? effect;

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
  inputField.addEventListener('keyup', inputListener.toJS);
}

// -------------------------------------------------------------- inputListener
void inputListener(js.JSObject something) {
  renderData();
}

// ----------------------------------------------------------------- renderData
void renderData() {
  context2d.clearRect(0, 0, mainCanvas.width, mainCanvas.height);
  effect?.wrapText(inputField.value);
}

// -------------------------------------------------------------------- animate

@js.JS('Function')
void animate(num value) {
  web.window.requestAnimationFrame(animate.toJS);
}

// ----------------------------------------------------------------------- main
void main() {
  setHeading();
  insertHRBefore('output');
  setupVariables();
  // fillCanvasWithPyramid(mainCanvas, context2d);
  // drawCenterLines(context2d, (w: mainCanvas.width, h: mainCanvas.height));
  drawWrappedText(context2d, 'Enter your message',
      (x: mainCanvas.width ~/ 2, y: mainCanvas.height ~/ 2));

  effect = Effect(context2d, mainCanvas.width, mainCanvas.height);
  effect?.wrapText(inputField.value);
  web.window.requestAnimationFrame(animate.toJS);
}
