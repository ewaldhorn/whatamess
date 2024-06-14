import 'package:flutter/material.dart';

void main() {
  runApp(buildApp());
}

MaterialApp buildApp() {
  return MaterialApp(
    home: Scaffold(
      body: Container(
        decoration: const BoxDecoration(gradient: DarkGreenGradient()),
        child: const Center(
          child: Text('Yolo!',
              style: TextStyle(color: Colors.white, fontSize: 28)),
        ),
      ),
    ),
  );
}

class DarkGreenGradient extends LinearGradient {
  const DarkGreenGradient()
      : super(begin: Alignment.topLeft, colors: const [
          Color.fromARGB(255, 33, 95, 33),
          Color.fromARGB(255, 29, 87, 29),
          Color.fromARGB(255, 23, 73, 23),
          Color.fromARGB(255, 23, 72, 24),
          Color.fromARGB(255, 24, 58, 25),
          Color.fromARGB(255, 18, 54, 20),
        ]);
}
