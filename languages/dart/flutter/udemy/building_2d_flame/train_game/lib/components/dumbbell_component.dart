import 'dart:async';

import 'package:flame/components.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class DumbbellComponent extends SpriteComponent
    with HasGameRef<FitFighterGame> {
  final double _spriteHeight = 60;
  @override
  Future<void> onLoad() async {
    super.onLoad();
    sprite = await gameRef.loadSprite(Globals.dumbbellSprite);
    height = width = _spriteHeight;
    position = Vector2(200, 200);
    anchor = Anchor.center;
  }
}
