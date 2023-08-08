import 'package:dice_app/widgets/my_widgets.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp(
      home: Scaffold(
          backgroundColor: Colors.yellow[800],
          body: GradientContainer(
            ListView(
              children: [
                ShadowedText('Yolo'),
                Image.asset("dice-images/dice-1.png")
              ],
            ),
          )),
    ),
  );
}
