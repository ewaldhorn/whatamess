import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:train_game/game/fit_fighter_game.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  final FitFighterGame theGame = FitFighterGame();
  runApp(GameWidget(game: theGame));
}
