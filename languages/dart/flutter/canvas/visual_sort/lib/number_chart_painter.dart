import 'package:flutter/material.dart';

class NumberChartPainter extends CustomPainter {
  final List<int> numbers;
  final int position;

  NumberChartPainter(this.numbers, {this.position = -1});

  static double borderSize = 10.0;

  // --------------------------------------------------------------------------
  @override
  void paint(Canvas canvas, Size size) {
    final clipRect = Rect.fromLTWH(0, 0, size.width, size.height);
    canvas.clipRect(clipRect);
    canvas.drawPaint(Paint()..color = const Color.fromARGB(255, 232, 205, 205));

    // calculate usable height
    final drawableHeight = size.height - (2 * borderSize);
    final drawableWidth = size.width - (2 * borderSize);

    // final hd = drawableHeight / 100;
    final wd = drawableWidth / numbers.length;
    final colWidth = wd - 1;

    _drawOutline(canvas, borderSize, drawableWidth, drawableHeight);

    final paint = Paint()..color = const Color.fromARGB(255, 255, 123, 83);
    final highlightPaint = Paint()..color = Colors.blue;

    final fontSize = _fontSizeFor(drawableWidth);

    for (int i = 0; i < numbers.length; i++) {
      final startOffset =
          Offset(borderSize + (wd * i), drawableHeight + borderSize);
      final endOffset = Offset(
          (borderSize + colWidth) + (wd * i),
          borderSize +
              (drawableHeight - (drawableHeight * (numbers[i] / 100))));
      final rect = Rect.fromPoints(startOffset, endOffset);
      canvas.drawRect(rect, (i == position) ? highlightPaint : paint);
      _drawTextCentered(canvas, rect.bottomCenter, '${numbers[i]}', fontSize);
    }
  }

  // --------------------------------------------------------------------------
  double _fontSizeFor(double width) {
    double fontSize = 5;

    if (width >= 1600) {
      fontSize = 14;
    } else if (width >= 1400) {
      fontSize = 12;
    } else if (width >= 1200) {
      fontSize = 10;
    } else if (width >= 800) {
      fontSize = 8;
    }

    return fontSize;
  }

  // --------------------------------------------------------------------------
  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true;
  }

  // --------------------------------------------------------------------------
  void _drawOutline(
      Canvas canvas, double borderSize, double width, double height) {
    final center = Offset(borderSize + (width / 2), borderSize + (height / 2));
    final rect = Rect.fromCenter(
        center: center, width: width + borderSize, height: height + borderSize);
    canvas.drawRect(
        rect,
        Paint()
          ..color = const Color.fromARGB(255, 175, 175, 175)
          ..style = PaintingStyle.stroke
          ..strokeWidth = 4.0);
  }

  // --------------------------------------------------------------------------
  TextPainter _measureText(
    String text,
    TextStyle style,
    double maxWidth,
    TextAlign align,
  ) {
    final span = TextSpan(text: text, style: style);
    final tp = TextPainter(
        text: span, textAlign: align, textDirection: TextDirection.ltr);
    tp.layout(minWidth: 0, maxWidth: maxWidth);

    return tp;
  }

  // --------------------------------------------------------------------------
  Size _drawTextCentered(
    Canvas canvas,
    Offset offset,
    String text,
    double fontSize,
  ) {
    final style = TextStyle(
        color: Colors.black, fontSize: fontSize, fontWeight: FontWeight.bold);
    const maxWidth = 30.0;
    final tp = _measureText(text, style, maxWidth, TextAlign.center);
    final pos = offset + Offset(-tp.width / 2, -tp.height);
    tp.paint(canvas, pos);

    return tp.size;
  }
}
