import 'package:flutter/material.dart';

class ProgressChart extends StatefulWidget {
  const ProgressChart({super.key});

  @override
  ProgressChartState createState() => ProgressChartState();
}

class ProgressChartState extends State<ProgressChart> {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: CustomPaint(
        child: Container(),
        painter: ChartPainter(),
      ),
    );
  }
}

class ChartPainter extends CustomPainter {
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
