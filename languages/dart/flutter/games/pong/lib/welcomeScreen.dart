import 'package:flutter/material.dart';

class Welcome extends StatelessWidget {

  final bool gameStarted;
  Welcome(this.gameStarted);

  @override
  Widget build(BuildContext context) {
    return Container(
        alignment: const Alignment(0, -0.2),
        child: Text(
          gameStarted ? "": "T A P   T O   P L A Y",
          style: const TextStyle(color: Colors.white),
        ));
  }
}