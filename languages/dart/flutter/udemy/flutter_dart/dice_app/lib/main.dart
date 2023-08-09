import 'package:dice_app/widgets/dice_roller.dart';
import 'package:dice_app/widgets/my_widgets.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp(
      home: Scaffold(
        backgroundColor: Colors.yellow[800],
        body: const Center(
          child: GradientContainer.purple(
            Padding(
              padding: EdgeInsets.all(120.0),
              child: DiceRoller(),
            ),
          ),
        ),
      ),
    ),
  );
}
