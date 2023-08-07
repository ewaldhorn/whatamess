import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp(
      home: Scaffold(
        backgroundColor: Colors.yellow[800],
        body: Container(
          decoration: const BoxDecoration(
              gradient: LinearGradient(colors: [
            Color.fromARGB(255, 243, 136, 61),
            Color.fromARGB(255, 220, 110, 40),
            Color.fromARGB(255, 181, 88, 31),
          ])),
          child: const Center(
            child: Text('Yolo'),
          ),
        ),
      ),
    ),
  );
}
