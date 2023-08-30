import 'dart:async';

import 'package:flame/game.dart';
import 'package:flame_audio/flame_audio.dart';
import 'package:train_game/components/background_component.dart';
import 'package:train_game/components/dumbbell_component.dart';
import 'package:train_game/components/player_component.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/inputs/joystick.dart';

class FitFighterGame extends FlameGame with HasCollisionDetection {
  @override
  Future<void> onLoad() async {
    super.onLoad();
    add(BackgroundComponent());
    add(PlayerComponent(joystick: joystick));
    add(DumbbellComponent());
    add(joystick);

    FlameAudio.audioCache.loadAll([
      Globals.dumbbellSound,
      Globals.proteinSound,
      Globals.vaccineSound,
      Globals.virusSound
    ]);
  }
}
