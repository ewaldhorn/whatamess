import 'dart:async';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame_audio/flame_audio.dart';
import 'package:train_game/components/player_component.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/game/fit_fighter_game.dart';

class DumbbellComponent extends SpriteComponent
    with HasGameRef<FitFighterGame>, CollisionCallbacks {
  final double _spriteHeight = 60;
  @override
  Future<void> onLoad() async {
    super.onLoad();
    sprite = await gameRef.loadSprite(Globals.dumbbellSprite);
    height = width = _spriteHeight;
    position = Vector2(200, 200);
    anchor = Anchor.center;
    add(RectangleHitbox());
  }

  @override
  void onCollision(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollision(intersectionPoints, other);
    if (other is PlayerComponent) {
      FlameAudio.play(Globals.dumbbellSound);
      removeFromParent();
    }
  }
}
