import 'dart:async';
import 'dart:math';

import 'package:flame/components.dart';
import 'package:flame/game.dart';
import 'package:flame/palette.dart';
import 'package:flame_audio/flame_audio.dart';
import 'package:flutter/painting.dart';
import 'package:train_game/components/background_component.dart';
import 'package:train_game/components/dumbbell_component.dart';
import 'package:train_game/components/player_component.dart';
import 'package:train_game/components/vaccine_component.dart';
import 'package:train_game/components/virus_component.dart';
import 'package:train_game/constants/globals.dart';
import 'package:train_game/inputs/joystick.dart';
import 'package:train_game/screens/game_over_menu.dart';

class FitFighterGame extends FlameGame with HasCollisionDetection {
  int score = 0;
  late Timer _gameTimer;
  late int _vaccineTimerAppearance;
  int _remainingTime = 30;
  late TextComponent _scoreText;
  late TextComponent _timeText;
  VaccineComponent vaccineComponent = VaccineComponent();

  @override
  Future<void> onLoad() async {
    super.onLoad();
    FlameAudio.audioCache.loadAll([
      Globals.dumbbellSound,
      Globals.proteinSound,
      Globals.vaccineSound,
      Globals.virusSound
    ]);

    _vaccineTimerAppearance = Random().nextInt(_remainingTime - 20) + 20;

    add(BackgroundComponent());
    add(VirusComponent(startPosition: Vector2(100, 150)));
    add(VirusComponent(startPosition: Vector2(size.x - 50, size.y - 200)));
    add(PlayerComponent(joystick: joystick));
    add(DumbbellComponent());
    add(joystick);

    add(ScreenHitbox());

    _gameTimer = Timer(1, repeat: true, onTick: () {
      if (_remainingTime <= 0) {
        pauseEngine();
        overlays.add(GameOverMenu.id);
      } else if (_remainingTime == _vaccineTimerAppearance) {
        add(vaccineComponent);
      } else {
        _remainingTime -= 1;
      }
    });

    _gameTimer.start();
    _scoreText = TextComponent(
        text: 'Score: $score',
        position: Vector2(40, 40),
        anchor: Anchor.topLeft,
        textRenderer: TextPaint(
            style: TextStyle(
          color: BasicPalette.white.color,
          fontSize: 25,
        )));

    _timeText = TextComponent(
        text: 'Time Left: $_remainingTime s',
        position: Vector2(size.x - 40, 40),
        anchor: Anchor.topRight,
        textRenderer: TextPaint(
            style: TextStyle(
          color: BasicPalette.white.color,
          fontSize: 25,
        )));

    add(_scoreText);
    add(_timeText);
  }

  @override
  void update(double dt) {
    super.update(dt);
    _gameTimer.update(dt);
    _scoreText.text = 'Score: $score';
    _timeText.text = 'Time Left: $_remainingTime s';
  }

  void reset() {
    score = 0;
    _remainingTime = 30;
  }
}
