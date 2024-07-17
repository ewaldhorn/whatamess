import 'package:web/web.dart' as web;
import 'dart:js_interop' as js;

import 'utils/canvasUtils.dart';
import 'utils/domUtils.dart';
import 'utils/headerManagement.dart';
import 'particles/particles.dart';

late web.HTMLCanvasElement mainCanvas;
late web.CanvasRenderingContext2D context2d;
late web.HTMLInputElement inputField;
late web.CanvasGradient gradient;

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

  gradient = createTextGradient(
      context2d, (w: mainCanvas.width, h: mainCanvas.height));
}

// -------------------------------------------------------------- inputListener
void inputListener(js.JSObject something) {
  renderData();
}

// ----------------------------------------------------------------- renderData
void renderData() {
  effect?.wrapText(inputField.value);
}

// -------------------------------------------------------------------- animate
int xPos = 0, yPos = 0;
int xDir = 1, yDir = 0;
final int rectWidth = 5;
bool hasLooped = false;

@js.JS('Function')
void animate(num value) {
  context2d.clearRect(10, 10, mainCanvas.width - 20, mainCanvas.height - 20);

  if (hasLooped) {
    context2d.fillStyle = 'darkred'.toJS;
  } else {
    context2d.fillStyle = gradient;
  }

  context2d.fillRect(xPos, yPos, rectWidth, rectWidth);

  xPos += xDir;
  yPos += yDir;

  if (xPos >= mainCanvas.width - rectWidth && yPos <= 0) {
    xDir = 0;
    yDir = 1;
  }

  if (yPos >= mainCanvas.height - rectWidth &&
      xPos >= mainCanvas.width - rectWidth) {
    yDir = 0;
    xDir = -1;
  }

  if (xPos <= 0 && yPos >= mainCanvas.height - rectWidth) {
    yDir = -1;
    xDir = 0;
  }

  if (xPos <= 0 && yPos <= 0) {
    xDir = 1;
    yDir = 0;
    hasLooped = !hasLooped;
  }

  effect?.wrapText(inputField.value);
  effect?.render();

  web.window.requestAnimationFrame(animate.toJS);
}

// ----------------------------------------------------------------------- main
void main() {
  setHeading();
  insertHRBefore('output');
  setupVariables();
  // fillCanvasWithPyramid(mainCanvas, context2d);
  // drawCenterLines(context2d, (w: mainCanvas.width, h: mainCanvas.height));

  effect = Effect(context2d, mainCanvas.width, mainCanvas.height);
  effect?.wrapText(inputField.value);
  effect?.render();
  web.window.requestAnimationFrame(animate.toJS);
}
