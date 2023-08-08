import 'package:flutter/material.dart';

const List<Color> defaultColours = [
  Color.fromARGB(255, 243, 136, 61),
  Color.fromARGB(255, 220, 110, 40),
  Color.fromARGB(255, 181, 88, 31),
];

class GradientContainer extends StatelessWidget {
  final Widget child;
  final List<Color> colourRange;

  const GradientContainer(this.child,
      {super.key, this.colourRange = defaultColours});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        gradient: LinearGradient(
          colors: colourRange,
          begin: Alignment.centerLeft,
          end: Alignment.bottomRight,
        ),
      ),
      child: child,
    );
  }
}
