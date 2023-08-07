import 'package:flutter/material.dart';

class ShadowedText extends StatelessWidget {
  const ShadowedText(this.msg, {super.key});

  final String msg;

  @override
  Widget build(BuildContext context) {
    return Text(msg,
        style: const TextStyle(
            color: Colors.white,
            fontWeight: FontWeight.bold,
            fontSize: 30.0,
            shadows: [Shadow(color: Colors.black, offset: Offset(2.0, 2.0))]));
  }
}
