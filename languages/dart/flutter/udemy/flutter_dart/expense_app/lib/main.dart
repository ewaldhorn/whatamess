import 'package:expense_app/app/expense_app.dart';
import 'package:flutter/material.dart';

var defaultColourScheme = ColorScheme.fromSeed(seedColor: Colors.orange);

void main() {
  runApp(
    MaterialApp(
        theme: ThemeData().copyWith(
          useMaterial3: true,
          colorScheme: defaultColourScheme,
          appBarTheme: const AppBarTheme().copyWith(
              backgroundColor: defaultColourScheme.onPrimaryContainer,
              foregroundColor: defaultColourScheme.primaryContainer),
        ),
        home: const ExpenseApp()),
  );
}
