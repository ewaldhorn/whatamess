import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:quiz_app/questions/questions.dart';

class ResultsScreen extends StatelessWidget {
  const ResultsScreen(this.answers, this.callback, {super.key});

  final void Function() callback;
  final List<String> answers;

  int calculateCorrectAnswers() {
    var total = 0;

    for (int i = 0; i < questions.length; i++) {
      if (answers.length >= i) {
        if (questions[i].answers[0] == answers[i]) {
          total += 1;
        }
      }
    }

    return total;
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      alignment: Alignment.center,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisSize: MainAxisSize.max,
        children: [
          Text(
            "You got ${calculateCorrectAnswers()} / ${answers.length} right!",
            style: GoogleFonts.lora(color: Colors.white, fontSize: 20),
            textAlign: TextAlign.center,
          ),
          const SizedBox(height: 30),
          OutlinedButton(
            onPressed: callback,
            child: const Text("Restart Quiz",
                style: TextStyle(color: Colors.white, fontSize: 20)),
          )
        ],
      ),
    );
  }
}
