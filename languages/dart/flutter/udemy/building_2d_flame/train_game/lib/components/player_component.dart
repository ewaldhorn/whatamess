import 'dart:async';

import 'package:flame/components.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class PlayerComponent extends SpriteComponent with HasGameRef<FitFighterGame> {
  final double _spriteHeight = 100;
  final double _speed = 500;

  JoystickComponent joystick;

  PlayerComponent({required this.joystick});

  @override
  void update(double dt) {
    super.update(dt);

    if (joystick.direction == JoystickDirection.idle) {
      return;
    }

    position.add(joystick.relativeDelta * _speed * dt);
  }

  @override
  Future<void> onLoad() async {
    await super.onLoad();
    sprite = await gameRef.loadSprite(Globals.playerSkinnySprite);
    position = gameRef.size / 2;
    height = width = _spriteHeight;
    anchor = Anchor.center;
  }
}
