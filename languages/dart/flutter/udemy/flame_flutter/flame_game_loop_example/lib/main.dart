import 'package:flame/flame.dart';
import 'package:flame/game.dart';
import 'package:flame/input.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  Flame.device.fullScreen();

  runApp(GameWidget(game: MyFirstGame()));
}

class MyFirstGame extends FlameGame with TapDetector, FPSCounter {
  static final fpsTextPaint =
      TextPaint(style: const TextStyle(color: Colors.black, fontSize: 12));

  @override
  Future<void>? onLoad() {
    //print('<game loop> onLoad() called');
    return super.onLoad();
  }

  @override
  void update(double dt) {
    super.update(dt);

    //print('<game loop> update called at delta time $dt');
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    canvas.drawPaint(Paint()..color = Colors.red);

    //
    // format and show the FPS for the game loop
    var stringFormatterFPS = NumberFormat('##', "en_US");
    String fpsString = "fps: " + stringFormatterFPS.format(fps(60));
    fpsTextPaint.render(canvas, fpsString, Vector2(10, 10));
    //print('<game loop> render called');
  }

  @override
  void onTapUp(TapUpInfo info) {
    print(
        '<game loop> onTap location: (${info.eventPosition.game.x}, ${info.eventPosition.game.y}');
  }
}
