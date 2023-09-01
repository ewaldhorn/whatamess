import 'package:fit_fighter/screens/main_menu.dart';
import 'package:flutter/material.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(
      const MaterialApp(debugShowCheckedModeBanner: false, home: MainMenu()));
}
