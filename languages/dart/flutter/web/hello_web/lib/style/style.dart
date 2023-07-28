import 'package:flutter/material.dart';
import 'package:hello_web/style/colours.dart';

class PrimaryText extends StatelessWidget {
  final double size;
  final FontWeight fontWeight;
  final Color colour;
  final String text;
  final double height;

  const PrimaryText(
      {required this.text,
      this.fontWeight = FontWeight.w400,
      this.colour = AppColours.primary,
      this.size = 20,
      this.height = 1.3});

  @override
  Widget build(BuildContext context) {
    return Text(text,
        style: TextStyle(
            color: colour,
            height: height,
            fontFamily: 'Poppins',
            fontSize: size,
            fontWeight: fontWeight));
  }
}
