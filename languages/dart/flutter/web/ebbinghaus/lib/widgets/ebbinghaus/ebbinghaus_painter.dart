import 'package:flutter/material.dart';

class EbbinghausPainter extends CustomPainter {
  @override
  void paint(Canvas canvas, Size size) {
    final paint = Paint()
      ..color = Colors.blue
      ..style = PaintingStyle.fill;

    const rect = Rect.fromLTWH(0, 0, 100, 100);
    canvas.drawRect(rect, paint);

    final textPainter = TextPainter(
      text: TextSpan(
        text: 'Hello, world!',
        style: TextStyle(color: Colors.black, fontSize: 30),
      ),
      textDirection: TextDirection.ltr,
    );

    textPainter.layout();
    textPainter.paint(canvas, Offset(50, 50));
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return false;
  }
}
