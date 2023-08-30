import 'dart:async';

import 'package:flame/game.dart';
import 'package:train_game/components/background_component.dart';

class FitFighterGame extends FlameGame {
  @override
  Future<void> onLoad() async {
    super.onLoad();
    add(BackgroundComponent());
  }
}
