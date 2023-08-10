import 'package:flutter/material.dart';
import 'package:quizz_app/constants/colours.dart';

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
      appBar: AppBar(title: const Text("My Quiz App")),
      body: Center(
        child: Column(
          mainAxisSize: MainAxisSize.max,
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Image.asset(
              "images/quiz-logo.png",
              width: 200,
            ),
            const SizedBox(
              height: 50,
            ),
            const Text(
              "Learn Flutter the fun way!",
              style: TextStyle(fontSize: 20, color: Colors.white),
            ),
            const SizedBox(
              height: 50,
            ),
            OutlinedButton(
                onPressed: () {},
                child: const Text(
                  'Start Quiz',
                  style: TextStyle(color: Colors.white),
                ))
          ],
        ),
      ),
    );
  }
}
