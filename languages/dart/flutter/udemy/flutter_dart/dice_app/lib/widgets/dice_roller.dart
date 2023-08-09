import 'package:dice_app/widgets/my_widgets.dart';
import 'package:flutter/material.dart';
import 'dart:math';

class DiceRoller extends StatefulWidget {
  const DiceRoller({super.key});

  @override
  State<DiceRoller> createState() => _DiceRollerState();
}

class _DiceRollerState extends State<DiceRoller> {
  final random = Random();
  var activeImageNumber = 1;
  var activeImagePath = 'dice-images/dice-1.png';

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      mainAxisAlignment: MainAxisAlignment.center,
      mainAxisSize: MainAxisSize.min,
      children: [
        const ShadowedText('Try your luck'),
        Image.asset(
          activeImagePath,
          width: 200,
        ),
        const SizedBox(
          height: 20,
        ),
        ElevatedButton(
          onPressed: () {
            activeImageNumber = random.nextInt(6) + 1;
            setState(() {
              activeImagePath = 'dice-images/dice-$activeImageNumber.png';
            });
          },
          child: const Text(
            "Roll it",
            style: TextStyle(fontSize: 25),
          ),
        ),
      ],
    );
  }
}
