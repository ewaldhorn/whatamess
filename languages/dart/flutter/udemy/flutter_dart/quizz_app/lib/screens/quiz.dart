import 'package:flutter/material.dart';
import 'package:quiz_app/constants/colours.dart';
import 'package:quiz_app/screens/quiz_screen.dart';
import 'package:quiz_app/screens/screens.dart';

class MainQuizWidget extends StatefulWidget {
  const MainQuizWidget({super.key});

  @override
  State<MainQuizWidget> createState() => _MainQuizWidgetState();
}

class _MainQuizWidgetState extends State<MainQuizWidget> {
  late Widget activeScreen;

  void startGame() {
    setState(() {
      activeScreen = const QuizScreen();
    });
  }

  @override
  void initState() {
    super.initState();
    activeScreen = MainScreen(startGame);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(backgroundColor: backgroundColour, body: activeScreen);
  }
}
