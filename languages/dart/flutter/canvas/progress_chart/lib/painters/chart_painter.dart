import 'package:flutter/material.dart';

class ChartPainter extends CustomPainter {
  final List<String> x;
  final List<double> y;
  final double min, max;

  ChartPainter(this.x, this.y, this.min, this.max);
  final linePaint = Paint()
    ..color = Colors.grey.shade700
    ..style = PaintingStyle.stroke
    ..strokeWidth = 0.5;
  static double border = 10.0;

  @override
  void paint(Canvas canvas, Size size) {
    // draw background
    final clipRect = Rect.fromLTWH(0, 0, size.width, size.height);
    canvas.clipRect(clipRect);
    canvas.drawPaint(Paint()..color = Colors.black);

    // set up some drawing dimensions
    final drawableHeight = size.height - 2.0 * border;
    final drawableWidth = size.width - 2.0 * border;
    final hd = drawableHeight / 5.0;
    final wd = drawableWidth / x.length.toDouble();
    final height = hd * 3.0;
    final width = drawableWidth;

    // check validity of data against our drawable area
    if (height <= 0.0 || width <= 0.0) return;
    if (max - min < 1.0e-6) return;

    final hr = height / (max - min); // height per unit value
    final left = border;
    final top = border;
    final c = Offset(left + wd / 2.0, top + height / 2.0);
    _drawOutline(canvas, c, wd, height);

    // now compute the data point locations
    final points = _computePoints(c, wd, height, hr);
    final path = _computePath(points);

    // draw the path, then overlay the points
    canvas.drawPath(path, linePaint);
    for (var p in points) {
      canvas.drawCircle(p, 5.0, Paint()..color = Colors.white);
    }
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true;
  }

  final outlinePaint = Paint()
    ..strokeWidth = 1
    ..style = PaintingStyle.stroke
    ..color = Colors.white;

  void _drawOutline(Canvas canvas, Offset c, double width, double height) {
    for (var yElement in y) {
      final rect = Rect.fromCenter(center: c, width: width, height: height);
      canvas.drawRect(rect, outlinePaint);
      c += Offset(width, 0);
    }
  }

  List<Offset> _computePoints(
      Offset c, double width, double height, double hr) {
    List<Offset> points = [];
    for (var element in y) {
      final yy = height - (element - min) * hr;
      final dp = Offset(c.dx, c.dy - height / 2.0 + yy);
      points.add(dp);
      c += Offset(width, 0);
    }
    return points;
  }

  Path _computePath(List<Offset> points) {
    Path result = Path();

    for (var i = 0; i < points.length; i++) {
      final p = points[i];
      if (i == 0) {
        result.moveTo(p.dx, p.dy);
      } else {
        result.lineTo(p.dx, p.dy);
      }
    }
    return result;
  }
}
