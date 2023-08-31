import 'dart:async';
import 'dart:math';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame_audio/flame_audio.dart';
import 'package:train_game/components/player_component.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class VaccineComponent extends SpriteComponent
    with HasGameRef<FitFighterGame>, CollisionCallbacks {
  final double _spriteHeight = 60;
  final Random _random = Random();

  Vector2 _getRandomPosition() {
    double x = 20 + _random.nextInt(gameRef.size.x.toInt() - 50).toDouble();
    double y = 20 + _random.nextInt(gameRef.size.y.toInt() - 50).toDouble();
    return Vector2(x, y);
  }

  @override
  Future<void> onLoad() async {
    super.onLoad();
    sprite = await gameRef.loadSprite(Globals.vaccineSprite);
    height = width = _spriteHeight;
    position = _getRandomPosition();
    anchor = Anchor.center;
    add(RectangleHitbox());
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);
    if (other is PlayerComponent) {
      FlameAudio.play(Globals.vaccineSound);
      removeFromParent();
    }
  }
}
