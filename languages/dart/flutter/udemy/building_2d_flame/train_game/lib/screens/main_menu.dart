import 'package:flutter/material.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/screens/game_play.dart';

class MainMenu extends StatelessWidget {
  static const String id = "MainMenu";
  const MainMenu({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Container(
      decoration: const BoxDecoration(
          image: DecorationImage(
              image: AssetImage('assets/images/${Globals.backgroundSprite}'),
              fit: BoxFit.cover)),
      child: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Padding(
              padding: EdgeInsets.symmetric(vertical: 50),
              child: Text('Fit Fighter', style: TextStyle(fontSize: 50)),
            ),
            const SizedBox(height: 250),
            SizedBox(
              width: 200,
              height: 50,
              child: ElevatedButton(
                  onPressed: () => Navigator.of(context).pushReplacement(
                        MaterialPageRoute(
                          builder: (context) => const GamePlay(),
                        ),
                      ),
                  child: const Text(
                    'Play',
                    style: TextStyle(fontSize: 25),
                  )),
            ),
            const SizedBox(height: 50),
            const Text(
              'Avoid getting sick, build your muscle.\nAll in 30 seconds!',
              textAlign: TextAlign.center,
              style: TextStyle(fontSize: 20),
            ),
          ],
        ),
      ),
    ));
  }
}
