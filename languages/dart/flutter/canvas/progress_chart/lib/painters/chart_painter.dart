import 'package:flutter/material.dart';

class ChartPainter extends CustomPainter {
  ChartPainter(List<String> x, List<double> y, double min, double max);

  @override
  void paint(Canvas canvas, Size size) {
    final c = Offset(size.width / 2, size.height / 2);
    final paint = Paint()..color = Colors.black;
    canvas.drawCircle(c, size.width / 12, paint);
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true;
  }
}
