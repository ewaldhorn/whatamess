import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:train_game/game/fit_fighter_game.dart';
import 'package:train_game/screens/game_over_menu.dart';

final FitFighterGame _fitFighterGame = FitFighterGame();

class GamePlay extends StatelessWidget {
  const GamePlay({super.key});

  @override
  Widget build(BuildContext context) {
    return GameWidget(game: _fitFighterGame, overlayBuilderMap: {
      GameOverMenu.id: (BuildContext context, FitFighterGame gameRef) =>
          GameOverMenu(game: gameRef),
    });
  }
}
