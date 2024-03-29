import 'dart:async';

import 'package:flame/components.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class BackgroundComponent extends SpriteComponent
    with HasGameRef<FitFighterGame> {
  @override
  Future<void> onLoad() async {
    await super.onLoad();

    sprite = await gameRef.loadSprite(Globals.backgroundSprite);
    opacity = 0.10;
    size = gameRef.size;
  }
}
