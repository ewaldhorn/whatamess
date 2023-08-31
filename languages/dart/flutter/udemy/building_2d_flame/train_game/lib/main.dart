import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:train_game/screens/main_menu.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(const MaterialApp(
    debugShowCheckedModeBanner: false,
    home: MainMenu(),
  ));
}
