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

class GradientContainer extends StatelessWidget {
  final Widget child;

  const GradientContainer(this.child, {super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        gradient: LinearGradient(colors: [
          Color.fromARGB(255, 243, 136, 61),
          Color.fromARGB(255, 220, 110, 40),
          Color.fromARGB(255, 181, 88, 31),
        ], begin: Alignment.centerLeft, end: Alignment.bottomRight),
      ),
      child: child,
    );
  }
}
