import 'dart:async';

import 'package:fit_fighter/constants/globals.dart';
import 'package:fit_fighter/games/fit_fighter_game.dart';
import 'package:flame/components.dart';

class BackgroundComponent extends SpriteComponent
    with HasGameRef<FitFighterGame> {
  @override
  Future<void> onLoad() async {
    super.onLoad();
    sprite = await gameRef.loadSprite(Globals.backgroundSprite);
    size = gameRef.size;
  }
}
