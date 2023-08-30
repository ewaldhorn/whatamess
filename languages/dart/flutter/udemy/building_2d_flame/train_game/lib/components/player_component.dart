import 'dart:async';

import 'package:flame/components.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class PlayerComponent extends SpriteComponent with HasGameRef<FitFighterGame> {
  final double _spriteHeight = 100;
  final double _speed = 500;

  late double _leftEdge, _rightEdge, _topEdge, _bottomEdge;

  JoystickComponent joystick;

  PlayerComponent({required this.joystick});

  @override
  void update(double dt) {
    super.update(dt);

    if (joystick.direction == JoystickDirection.idle) {
      return;
    }

    if (x >= _rightEdge) {
      x = _rightEdge;
    }

    if (x <= _leftEdge) {
      x = _leftEdge;
    }

    if (y <= _topEdge) {
      y = _topEdge;
    }

    if (y >= _bottomEdge) {
      y = _bottomEdge;
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

    _rightEdge = gameRef.size.x - 60;
    _leftEdge = 60;
    _topEdge = 60;
    _bottomEdge = gameRef.size.y - 60;
  }
}
