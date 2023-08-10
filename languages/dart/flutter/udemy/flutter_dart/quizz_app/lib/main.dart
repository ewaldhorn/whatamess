import 'package:flutter/material.dart';
import 'package:quizz_app/constants/colours.dart';
import 'package:quizz_app/screens/screens.dart';

void main() {
  runApp(const MaterialApp(
    home: MyAppPage(),
  ));
}

class MyAppPage extends StatelessWidget {
  const MyAppPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColour,
      body: MainScreen(),
    );
  }
}
