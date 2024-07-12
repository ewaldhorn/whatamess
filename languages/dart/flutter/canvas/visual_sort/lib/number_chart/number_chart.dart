import 'package:flutter/material.dart';

import 'number_chart_painter.dart';

class NumberChart extends StatefulWidget {
  final int position, comparePosition;
  final List<int> numbers;

  const NumberChart(
      {super.key,
      required this.numbers,
      this.position = -1,
      this.comparePosition = 0});

  @override
  State<NumberChart> createState() => _NumberChartState();
}

class _NumberChartState extends State<NumberChart> {
  @override
  Widget build(BuildContext context) {
    return CustomPaint(
        painter: NumberChartPainter(widget.numbers,
            position: widget.position, comparePosition: widget.comparePosition),
        child: Container());
  }
}
