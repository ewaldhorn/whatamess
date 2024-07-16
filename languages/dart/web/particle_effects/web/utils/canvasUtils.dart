import 'dart:js_interop';

import 'package:web/web.dart' as web;
import 'stringUtils.dart';
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

// ------------------------------------------------------------------- drawText
void drawText(
  web.CanvasRenderingContext2D ctx,
  String text,
  Position position, {
  String alignment = 'center',
  String baseline = 'middle',
  num lineWidth = 1,
}) {
  ctx.fillStyle = 'white'.toJS;
  ctx.strokeStyle = 'green'.toJS;
  ctx.textAlign = alignment;
  ctx.textBaseline = baseline;
  ctx.lineWidth = lineWidth;
  ctx.font = '50px Georgia';

  ctx.fillText(text, position.x, position.y);
  ctx.strokeText(text, position.x, position.y);
}

// ------------------------------------------------------------ drawWrappedText
void drawWrappedText(
  web.CanvasRenderingContext2D ctx,
  String text,
  Position position, {
  String alignment = 'center',
  String baseline = 'middle',
  num lineWidth = 1,
  int maxCharacters = 20,
  int gap = 45,
}) {
  List<String> strings = lineSplitter(text, maxLen: maxCharacters);

  final int startY = position.y - ((strings.length ~/ 2) * gap);

  for (int i = 0; i < strings.length; i++) {
    drawText(ctx, strings[i], (x: position.x, y: startY + (i * gap)),
        alignment: alignment);
  }
}

// ------------------------------------------------------------ drawCenterLines
void drawCenterLines(web.CanvasRenderingContext2D ctx, Size size) {
  final midX = size.w ~/ 2;
  final midY = size.h ~/ 2;

  ctx.strokeStyle = 'grey'.toJS;
  ctx.lineWidth = 1;
  ctx.setLineDash([1.toJS, 0.toJS, 0.toJS, 1.toJS, 0.toJS].toJS);

  ctx.beginPath();
  ctx.moveTo(0, midY);
  ctx.lineTo(size.w, midY);
  ctx.closePath();
  ctx.stroke();

  ctx.beginPath();
  ctx.moveTo(midX, 0);
  ctx.lineTo(midX, size.h);
  ctx.closePath();
  ctx.stroke();
}
