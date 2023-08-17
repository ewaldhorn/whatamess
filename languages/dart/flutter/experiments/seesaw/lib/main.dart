import 'package:flutter/material.dart';
import 'package:charts_flutter/flutter.dart' as charts;
import 'package:seesaw/data.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Data Vis',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const VisApp(),
    );
  }
}

class VisApp extends StatelessWidget {
  const VisApp({super.key});

  @override
  Widget build(BuildContext context) {
    {
      return Scaffold(
        appBar: AppBar(
          title: const Text('Data Vis App'),
          centerTitle: true,
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              SizedBox(
                height: 250,
                child: charts.BarChart(
                  SampleData.createChartData(),
                  animate: true,
                  behaviors: [
                    charts.ChartTitle('Company Size vs Number of Companies'),
                    charts.ChartTitle('Number of Companies',
                        behaviorPosition: charts.BehaviorPosition.start),
                    charts.ChartTitle('Company Size',
                        behaviorPosition: charts.BehaviorPosition.bottom)
                  ],
                ),
              )
            ],
          ),
        ),
      );
    }
  }
}
