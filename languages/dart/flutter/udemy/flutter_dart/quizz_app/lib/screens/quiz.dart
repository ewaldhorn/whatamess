import 'package:flutter/material.dart';
import 'package:quizz_app/constants/colours.dart';
import 'package:quizz_app/screens/quiz_screen.dart';
import 'package:quizz_app/screens/screens.dart';

class MainQuizWidget extends StatefulWidget {
  const MainQuizWidget({super.key});

  @override
  State<MainQuizWidget> createState() => _MainQuizWidgetState();
}

class _MainQuizWidgetState extends State<MainQuizWidget> {
  bool isPlaying = false;

  void startGame() {
    setState(() {
      isPlaying = true;
    });
  }

  Widget getCurrentScreen() {
    if (isPlaying) {
      return const QuizScreen();
    } else {
      return const MainScreen(startGame);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: backgroundColour, body: getCurrentScreen());
  }
}
