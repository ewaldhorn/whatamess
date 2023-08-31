import 'package:flutter/material.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class GameOverMenu extends StatelessWidget {
  static const String id = "GameOverMenu";
  final FitFighterGame game;

  const GameOverMenu({super.key, required this.game});

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
            const Text("Game Over", style: TextStyle(fontSize: 50)),
            const SizedBox(height: 200),
            SizedBox(
              width: 200,
              height: 50,
              child: ElevatedButton(
                onPressed: () {},
                child:
                    const Text('Play Again?', style: TextStyle(fontSize: 25)),
              ),
            ),
            const SizedBox(height: 20),
            SizedBox(
              width: 200,
              height: 50,
              child: ElevatedButton(
                  onPressed: () {},
                  child:
                      const Text('Main Menu', style: TextStyle(fontSize: 25))),
            )
          ],
        ),
      ),
    ));
  }
}
