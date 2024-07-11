import 'dart:math';

import 'package:flutter/material.dart';

import 'number_chart.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Visual Sorter',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepOrange),
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'Flutter Visual Sorter'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  bool isBusy = false;
  int _position = -1;
  final _r = Random();
  late List<int> _numbers;

  @override
  void initState() {
    super.initState();

    // generate a list of 50 random numbers
    initNumbersList();
  }

  // --------------------------------------------------------------------------
  // reset the numbers list, set the sorting position back to -1
  void initNumbersList() {
    final numbers = List.generate(50, (_) {
      return 1 + _r.nextInt(100);
    });

    setState(() {
      _numbers = numbers;
      _position = -1;
    });
  }

  // --------------------------------------------------------------------------
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          children: <Widget>[
            SizedBox(
              height: 300,
              child: NumberChart(
                numbers: _numbers,
                position: _position,
              ),
            ),
            const SizedBox(height: 16),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                TextButton(
                  onPressed: isBusy
                      ? null
                      : () {
                          setState(() {
                            isBusy = true;
                          });
                        },
                  child: const Text('Start',
                      style: TextStyle(fontWeight: FontWeight.bold)),
                ),
                TextButton(
                  onPressed: isBusy
                      ? () {
                          initNumbersList();
                          setState(() {
                            isBusy = false;
                          });
                        }
                      : null,
                  child: const Text('Reset',
                      style: TextStyle(fontWeight: FontWeight.bold)),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
