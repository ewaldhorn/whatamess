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
      return 1 + _r.nextInt(100);
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
  NumberChartPainter(List<int> numbers);

  static double borderSize = 10.0;

  @override
  void paint(Canvas canvas, Size size) {
    final clipRect = Rect.fromLTWH(0, 0, size.width, size.height);
    canvas.clipRect(clipRect);
    canvas.drawPaint(Paint()..color = Colors.grey);

    // calculate usable height
    final drawableHeight = size.height - (2 * borderSize);
    final drawableWidth = size.width - (2 * borderSize);

    final paint = Paint()..color = Colors.deepOrange;
    canvas.drawCircle(Offset(size.width / 2, size.height / 2), 50, paint);
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true;
  }
}
