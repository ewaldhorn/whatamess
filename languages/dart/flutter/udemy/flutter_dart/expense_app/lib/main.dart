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
          cardTheme: const CardTheme().copyWith(
            color: defaultColourScheme.secondaryContainer,
            margin: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
          ),
          elevatedButtonTheme: ElevatedButtonThemeData(
            style: ElevatedButton.styleFrom(
                backgroundColor: defaultColourScheme.primaryContainer),
          ),
          textTheme: ThemeData().textTheme.copyWith(
                titleLarge: const TextStyle(fontWeight: FontWeight.bold),
              ),
        ),
        home: const ExpenseApp()),
  );
}
