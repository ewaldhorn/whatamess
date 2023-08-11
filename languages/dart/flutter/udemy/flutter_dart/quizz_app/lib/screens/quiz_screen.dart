import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:quiz_app/widgets/answer_button.dart';
import 'package:quiz_app/questions/questions.dart';

class QuizScreen extends StatefulWidget {
  const QuizScreen({super.key, required this.endGameCallback});

  final void Function() endGameCallback;

  @override
  State<QuizScreen> createState() => _QuizScreenState();
}

class _QuizScreenState extends State<QuizScreen> {
  final questionCount = questions.length;
  var currentQuestionNumber = 0;

  void answerQuestion() {
    if (currentQuestionNumber < (questionCount - 1)) {
      setState(() {
        currentQuestionNumber++;
      });
    } else {
      widget.endGameCallback();
    }
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: Container(
        margin: const EdgeInsets.symmetric(horizontal: 120, vertical: 0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          mainAxisSize: MainAxisSize.max,
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            Text(
              questions[currentQuestionNumber].question,
              textAlign: TextAlign.center,
              style: GoogleFonts.lato(
                  color: Colors.white,
                  fontSize: 28,
                  fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 20),
            ...questions[currentQuestionNumber].getShuffledAnswers().map((a) {
              return AnswerButton(
                  text: a,
                  onPressed: () {
                    answerQuestion();
                  });
            }),
            const SizedBox(height: 20),
            Text(
              "${currentQuestionNumber + 1} / $questionCount",
              textAlign: TextAlign.center,
              style: const TextStyle(color: Colors.grey),
            )
          ],
        ),
      ),
    );
  }
}
