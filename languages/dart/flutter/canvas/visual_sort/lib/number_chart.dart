import 'dart:math';

import 'package:flutter/material.dart';

class NumberChart extends StatefulWidget {
  const NumberChart({super.key});

  @override
  State<NumberChart> createState() => _NumberChartState();
}

class _NumberChartState extends State<NumberChart> {
  final _r = Random();
  late List<int> _numbers;

  @override
  void initState() {
    super.initState();

    // generate a list of 50 random numbers
    final numbers = List.generate(50, (_) {
      return 2 + _r.nextInt(100);
    });

    setState(() {
      _numbers = numbers;
    });
  }

  @override
  Widget build(BuildContext context) {
    return CustomPaint(
        painter: NumberChartPainter(_numbers), child: Container());
  }
}

class NumberChartPainter extends CustomPainter {
  final List<int> numbers;

  NumberChartPainter(this.numbers);

  static double borderSize = 10.0;

  @override
  void paint(Canvas canvas, Size size) {
    final clipRect = Rect.fromLTWH(0, 0, size.width, size.height);
    canvas.clipRect(clipRect);
    canvas.drawPaint(Paint()..color = Colors.grey);

    // calculate usable height
    final drawableHeight = size.height - (2 * borderSize);
    final drawableWidth = size.width - (2 * borderSize);

    // final hd = drawableHeight / 100;
    final wd = drawableWidth / numbers.length;
    final colWidth = wd - 1;

    _drawOutline(canvas, borderSize, drawableWidth, drawableHeight);

    final paint = Paint()..color = Colors.deepOrange;

    for (int i = 0; i < numbers.length; i++) {
      final startOffset =
          Offset(borderSize + (wd * i), drawableHeight + borderSize);
      final endOffset = Offset((borderSize + colWidth) + (wd * i),
          borderSize + (drawableHeight * (numbers[i] / 100)));
      final rect = Rect.fromPoints(startOffset, endOffset);
      canvas.drawRect(rect, paint);
    }
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true;
  }

  void _drawOutline(
      Canvas canvas, double borderSize, double width, double height) {
    final center = Offset(borderSize + (width / 2), borderSize + (height / 2));
    final rect = Rect.fromCenter(
        center: center, width: width + borderSize, height: height + borderSize);
    canvas.drawRect(
        rect,
        Paint()
          ..color = const Color.fromARGB(255, 122, 121, 121)
          ..style = PaintingStyle.stroke
          ..strokeWidth = 2.0);
  }
}
