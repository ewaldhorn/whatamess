import 'dart:async';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame_audio/flame_audio.dart';
import 'package:train_game/components/protein_component.dart';
import 'package:train_game/components/vaccine_component.dart';
import 'package:train_game/components/virus_component.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class PlayerComponent extends SpriteComponent
    with HasGameRef<FitFighterGame>, CollisionCallbacks {
  final double _spriteHeight = 100;
  final double _speed = 500;
  bool _virusAttacked = false;
  bool _vaccinated = false;
  final Timer _feverTimer = Timer(3);
  final Timer _vaccinatedTimer = Timer(5);

  late double _leftEdge, _rightEdge, _topEdge, _bottomEdge;
  late Sprite playerSkinny, playerFever, playerFit, playerMuscular;

  JoystickComponent joystick;

  PlayerComponent({required this.joystick});

  void _selectAppropriatePlayerSprite() {
    if (gameRef.score >= 60) {
      sprite = playerMuscular;
    } else if (gameRef.score >= 30) {
      sprite = playerFit;
    } else {
      sprite = playerSkinny;
    }
  }

  void _vaccinatePlayer() {
    _vaccinated = true;
    _vaccinatedTimer.start();
  }

  void _freezePlayer() {
    if (!_virusAttacked && !_vaccinated) {
      FlameAudio.play(Globals.virusSound);
      _virusAttacked = true;
      gameRef.score -= 2;
      sprite = playerFever;
      _feverTimer.start();
    }
  }

  void _unfreezePlayer() {
    _virusAttacked = false;
    _selectAppropriatePlayerSprite();
  }

  @override
  void update(double dt) {
    super.update(dt);

    if (_vaccinated) {
      _vaccinatedTimer.update(dt);
      if (_vaccinatedTimer.finished) {
        _vaccinated = false;
        scale = Vector2(1.0, 1.0);
      } else {
        scale = Vector2(2.0, 2.0);
      }
    }

    if (_virusAttacked) {
      _feverTimer.update(dt);
      if (_feverTimer.finished) {
        _unfreezePlayer();
      }
    } else {
      _selectAppropriatePlayerSprite();
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

    _selectAppropriatePlayerSprite();
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
    } else if (other is VaccineComponent) {
      if (!_virusAttacked) {
        _vaccinatePlayer();
      }
    } else if (other is ProteinComponent) {
      game.score += 7;
    }
  }
}
