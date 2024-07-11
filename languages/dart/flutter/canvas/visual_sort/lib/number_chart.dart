import 'dart:math';

import 'package:flutter/material.dart';

import 'number_chart_painter.dart';

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
