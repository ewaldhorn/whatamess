import 'package:flutter/material.dart';

class ResultsScreen extends StatelessWidget {
  const ResultsScreen(this.answers, {super.key});

  final List<String> answers;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [Text("There are ${answers.length}.")],
    );
  }
}
