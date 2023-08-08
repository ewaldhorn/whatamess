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

  const GradientContainer.purple(this.child, {super.key})
      : colourRange = const [
          Color.fromARGB(255, 128, 64, 128),
          Color.fromARGB(255, 255, 128, 255)
        ];

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
