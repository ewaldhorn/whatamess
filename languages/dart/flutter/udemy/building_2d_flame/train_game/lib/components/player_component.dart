import 'dart:async';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame_audio/flame_audio.dart';
import 'package:train_game/components/virus_component.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class PlayerComponent extends SpriteComponent
    with HasGameRef<FitFighterGame>, CollisionCallbacks {
  final double _spriteHeight = 100;
  final double _speed = 500;
  bool _virusAttacked = false;
  final Timer _feverTimer = Timer(3);

  late double _leftEdge, _rightEdge, _topEdge, _bottomEdge;
  late Sprite playerSkinny, playerFever, playerFit, playerMuscular;

  JoystickComponent joystick;

  PlayerComponent({required this.joystick});

  void _freezePlayer() {
    if (!_virusAttacked) {
      FlameAudio.play(Globals.virusSound);
      _virusAttacked = true;
      gameRef.score -= 2;
      sprite = playerFever;
      _feverTimer.start();
    }
  }

  void _unfreezePlayer() {
    _virusAttacked = false;
    sprite = playerSkinny;
  }

  @override
  void update(double dt) {
    super.update(dt);

    if (_virusAttacked) {
      _feverTimer.update(dt);
      if (_feverTimer.finished) {
        _unfreezePlayer();
      }
    } else {
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
  }

  @override
  Future<void> onLoad() async {
    await super.onLoad();

    playerFever = await gameRef.loadSprite(Globals.playerFeverSprite);
    playerFit = await gameRef.loadSprite(Globals.playerFitSprite);
    playerMuscular = await gameRef.loadSprite(Globals.playerMuscularSprite);
    playerSkinny = await gameRef.loadSprite(Globals.playerSkinnySprite);

    sprite = playerSkinny;
    position = gameRef.size / 2;
    height = width = _spriteHeight;
    anchor = Anchor.center;

    _rightEdge = gameRef.size.x - 60;
    _leftEdge = 60;
    _topEdge = 60;
    _bottomEdge = gameRef.size.y - 60;

    add(RectangleHitbox());
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);
    if (other is VirusComponent) {
      _freezePlayer();
    }
  }
}
