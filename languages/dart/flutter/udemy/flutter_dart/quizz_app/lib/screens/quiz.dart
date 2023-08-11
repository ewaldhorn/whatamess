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
  List<String> selectedAnswers = [];

  void startGame() {
    setState(() {
      activeScreen = QuizScreen(
        endGameCallback: endGame,
        answerCallback: chooseAnswer,
      );
    });
  }

  void endGame() {
    setState(() {
      activeScreen = ResultsScreen(selectedAnswers);
    });
  }

  void chooseAnswer(String choice) {
    selectedAnswers.add(choice);
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
