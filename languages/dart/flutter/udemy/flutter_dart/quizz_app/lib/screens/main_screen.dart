import 'package:flutter/material.dart';

class MainScreen extends StatelessWidget {
  const MainScreen({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
          gradient: LinearGradient(colors: [
        Color.fromARGB(255, 245, 107, 52),
        Color.fromARGB(255, 217, 83, 30),
        Color.fromARGB(255, 139, 64, 18),
        Color.fromARGB(255, 113, 47, 6),
      ])),
      child: Center(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Image.asset(
              "images/quiz-logo.png",
              width: 200,
              color: const Color.fromARGB(175, 255, 255, 255),
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
            OutlinedButton.icon(
              icon: const Icon(Icons.arrow_forward,
                  color: Colors.white, size: 16),
              onPressed: () {},
              style: OutlinedButton.styleFrom(elevation: 2.0),
              label: const Text(
                'Start Quiz',
                style: TextStyle(color: Colors.white),
              ),
            )
          ],
        ),
      ),
    );
  }
}
