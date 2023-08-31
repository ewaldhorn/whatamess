import 'dart:async';

import 'package:flame/collisions.dart';
import 'package:flame/game.dart';
import 'package:flame_audio/flame_audio.dart';
import 'package:train_game/components/background_component.dart';
import 'package:train_game/components/dumbbell_component.dart';
import 'package:train_game/components/player_component.dart';
import 'package:train_game/components/virus_component.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/inputs/joystick.dart';

class FitFighterGame extends FlameGame with HasCollisionDetection {
  @override
  Future<void> onLoad() async {
    super.onLoad();
    FlameAudio.audioCache.loadAll([
      Globals.dumbbellSound,
      Globals.proteinSound,
      Globals.vaccineSound,
      Globals.virusSound
    ]);

    add(BackgroundComponent());
    add(VirusComponent(startPosition: Vector2(100, 150)));
    add(VirusComponent(startPosition: Vector2(size.x - 50, size.y - 200)));
    add(PlayerComponent(joystick: joystick));
    add(DumbbellComponent());
    add(joystick);

    add(ScreenHitbox());
  }
}
