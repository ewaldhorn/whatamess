import 'package:flutter/material.dart';
import 'package:train_game/screens/game_splash.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(const MaterialApp(
    debugShowCheckedModeBanner: false,
    home: GameSplashScreen(),
  ));
}
