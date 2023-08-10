import 'package:flutter/material.dart';

class MainScreen extends StatelessWidget {
  const MainScreen({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Center(
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
    );
  }
}
