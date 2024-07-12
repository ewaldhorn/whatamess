import 'dart:async';
import 'dart:math';

import 'package:flutter/material.dart';

import '../number_chart/number_chart.dart';

class BubbleSorter extends StatefulWidget {
  const BubbleSorter({super.key});

  @override
  State<BubbleSorter> createState() => _BubbleSorterState();
}

class _BubbleSorterState extends State<BubbleSorter> {
  bool _canSort = false;
  bool _canReset = false;
  bool _isSorting = false;
  bool _haveSwapped = false;
  Timer? _sortTimer;
  int _position = -1;
  int _comparePosition = 0;
  final _r = Random();
  late List<int> _numbers;

  @override
  void initState() {
    super.initState();

    initNumbersList();
  }

  @override
  void dispose() {
    _sortTimer?.cancel();
    super.dispose();
  }

  // --------------------------------------------------------------------------
  // reset the numbers list, set the sorting position back to -1
  void initNumbersList() {
    final numbers = List.generate(100, (_) {
      return 1 + _r.nextInt(100);
    });

    setState(() {
      _numbers = numbers;
      _position = -1;
      _comparePosition = 0;
      _canSort = true;
      _canReset = true;
      _isSorting = false;
      _haveSwapped = false;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: <Widget>[
        SizedBox(
          height: 300,
          child: NumberChart(
            numbers: _numbers,
            position: (_position >= 0) ? _numbers.length - _position : -1,
            comparePosition: _comparePosition,
          ),
        ),
        const SizedBox(height: 16),
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextButton(
              onPressed: _canSort ? _startButtonPressed : null,
              child: const Text('Start',
                  style: TextStyle(fontWeight: FontWeight.bold)),
            ),
            TextButton(
              onPressed: _canReset ? _resetButtonPressed : null,
              child: const Text('Reset',
                  style: TextStyle(fontWeight: FontWeight.bold)),
            ),
          ],
        ),
      ],
    );
  }

  // --------------------------------------------------------------------------
  void _startButtonPressed() {
    setState(() {
      _canSort = false;
      _canReset = false;
    });
    startSorting();
  }

  // --------------------------------------------------------------------------
  void _resetButtonPressed() {
    initNumbersList();
  }

  // --------------------------------------------------------------------------
  void startSorting() {
    setState(() {
      _position = 0;
      _comparePosition = 0;
      _isSorting = true;
      _haveSwapped = false;
      _sortTimer = Timer.periodic(const Duration(milliseconds: 10), (t) {
        if (_isSorting) {
          _haveSwapped = true;
          walkPosition();
        } else {
          _sortTimer?.cancel();
          _canReset = true;
        }
      });
    });
  }

  // --------------------------------------------------------------------------
  void walkPosition() {
    if (_position < _numbers.length && _haveSwapped) {
      setState(() {
        _haveSwapped = false;

        if (_numbers[_comparePosition] > _numbers[_comparePosition + 1]) {
          _haveSwapped = true;
          int tmp = _numbers[_comparePosition];
          _numbers[_comparePosition] = _numbers[_comparePosition + 1];
          _numbers[_comparePosition + 1] = tmp;
        }

        if (_comparePosition > _numbers.length - _position - 3) {
          _position += 1;
          _comparePosition = 0;
        } else {
          _comparePosition += 1;
        }
      });
    } else {
      setState(() {
        _position = -1;
        _canReset = true;
        _canSort = true;
        _haveSwapped = false;
        _comparePosition = 0;
        _position = -1;
        _sortTimer?.cancel();
      });
    }
  }
}
