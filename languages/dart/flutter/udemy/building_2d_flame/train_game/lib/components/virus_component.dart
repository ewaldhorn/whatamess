import 'dart:async';
import 'dart:math';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class VirusComponent extends SpriteComponent
    with HasGameRef<FitFighterGame>, CollisionCallbacks {
  final double _spriteHeight = 30;
  final Vector2 startPosition;
  final Random _random = Random();
  late Vector2 _velocity;
  double _speed = 115.0;
  int beenAlive = 0;

  VirusComponent({required this.startPosition});

  Vector2 _moveSprite() {
    final randomAngle = Random().nextDouble() * 2 * pi;
    return Vector2(cos(randomAngle) * _speed, sin(randomAngle) * _speed);
  }

  Vector2 _getRandomPosition() {
    _speed = _random.nextInt(75) + 75.0;
    double x = 25.0 + _random.nextInt(gameRef.size.x.toInt() - 100).toDouble();
    double y = 25.0 + _random.nextInt(gameRef.size.y.toInt() - 100).toDouble();
    return Vector2(x, y);
  }

  @override
  Future<void> onLoad() async {
    await super.onLoad();
    sprite = await gameRef.loadSprite(Globals.virusSprite);
    height = width = _spriteHeight;
    position = startPosition;
    anchor = Anchor.center;
    add(CircleHitbox());
    _velocity = _moveSprite();
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);

    if (other is ScreenHitbox || other is VirusComponent) {
      final Vector2 collisionPoint = intersectionPoints.first;
      if (collisionPoint.x <= 0 || collisionPoint.x >= gameRef.size.x) {
        // hit left wall or right wall
        _velocity.x *= -1;
      }

      if (collisionPoint.y <= 0 || collisionPoint.y >= gameRef.size.y) {
        // hit top or bottom wall
        _velocity.y *= -1;
      }
    }
  }

  @override
  void update(double dt) {
    super.update(dt);
    position += _velocity * dt;
    beenAlive += 1;

    if (beenAlive > 1000) {
      beenAlive = 0;
      position = _getRandomPosition();
    }

    if (position.x < -25 ||
        position.x > gameRef.size.x + 25 ||
        position.y < -25 ||
        position.y > gameRef.size.y + 25) {
      position = _getRandomPosition();
    }
  }
}
