import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp(
        home: Material(
      color: Colors.amber.shade800,
      child: const Center(
        child: Text('''Hello World!!! 
        Welcome to Flame Crash Course
        Code a 2D Mobile Game in less than 6 hours!!!''',
            textAlign: TextAlign.center),
      ),
    )),
  );
}
