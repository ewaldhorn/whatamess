import 'package:dice_app/widgets/gradient_container.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp(
      home: Scaffold(
          backgroundColor: Colors.yellow[800],
          body: const GradientContainer(
            Center(
              child: Text('Yolo',
                  style: TextStyle(
                      color: Colors.white,
                      fontWeight: FontWeight.bold,
                      fontSize: 30.0,
                      shadows: [
                        Shadow(color: Colors.black, offset: Offset(2.0, 2.0))
                      ])),
            ),
          )),
    ),
  );
}
