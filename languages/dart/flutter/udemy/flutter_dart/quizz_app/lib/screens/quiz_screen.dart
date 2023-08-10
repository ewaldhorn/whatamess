import 'package:flutter/material.dart';

class QuizScreen extends StatefulWidget {
  const QuizScreen({super.key});

  @override
  State<QuizScreen> createState() => _QuizScreenState();
}

class _QuizScreenState extends State<QuizScreen> {
  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        mainAxisSize: MainAxisSize.max,
        children: [
          Text(
            "Question",
            style: TextStyle(color: Colors.white, fontSize: 24),
          ),
          const SizedBox(height: 20),
          ElevatedButton(onPressed: () {}, child: Text("Answer One")),
          SizedBox(height: 10),
          ElevatedButton(onPressed: () {}, child: Text("Answer Two")),
          SizedBox(height: 10),
          ElevatedButton(onPressed: () {}, child: Text("Answer Three")),
          SizedBox(height: 10),
          ElevatedButton(onPressed: () {}, child: Text("Answer Four"))
        ],
      ),
    );
  }
}
