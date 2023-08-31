import 'package:flutter/material.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';
import 'package:train_game/screens/main_menu.dart';

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
            Text(
              'You scored ${game.score}',
              style: const TextStyle(fontSize: 20),
            ),
            const SizedBox(height: 200),
            SizedBox(
              width: 200,
              height: 50,
              child: ElevatedButton(
                onPressed: () {
                  game.overlays.remove(GameOverMenu.id);
                  game.reset();
                  game.resumeEngine();
                },
                child:
                    const Text('Play Again?', style: TextStyle(fontSize: 25)),
              ),
            ),
            const SizedBox(height: 20),
            SizedBox(
              width: 200,
              height: 50,
              child: ElevatedButton(
                  onPressed: () {
                    game.overlays.remove(GameOverMenu.id);
                    game.reset();
                    game.resumeEngine();
                    Navigator.pushReplacement(context,
                        MaterialPageRoute(builder: (ctx) => const MainMenu()));
                  },
                  child:
                      const Text('Main Menu', style: TextStyle(fontSize: 25))),
            )
          ],
        ),
      ),
    ));
  }
}
