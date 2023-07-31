import 'package:flutter/material.dart';
import 'package:progress_chart/data/scores.dart';
import 'package:progress_chart/painters/chart_painter.dart';

class ProgressChart extends StatefulWidget {
  final List<Score> scores;

  const ProgressChart(this.scores, {super.key});

  @override
  ProgressChartState createState() => ProgressChartState();
}

const weekDays = ["", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];

class ProgressChartState extends State<ProgressChart> {
  late double _min, _max;

  late List<double> _Y;
  late List<String> _X;

  @override
  void initState() {
    super.initState();
    var min = double.maxFinite;
    var max = -double.maxFinite;

    for (var element in widget.scores) {
      min = min > element.value ? element.value : min;
      max = max < element.value ? element.value : max;
    }

    setState(() {
      _min = min;
      _max = max;
      _Y = widget.scores.map((e) => e.value).toList();
      _X = widget.scores
          .map((e) => "${weekDays[e.time.weekday]}\n${e.time.day}")
          .toList();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      child: CustomPaint(
        painter: ChartPainter(_X, _Y, _min, _max),
        child: Container(),
      ),
    );
  }
}
