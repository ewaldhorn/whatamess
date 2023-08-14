import 'package:expense_app/app/expense_app.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp(theme: ThemeData(useMaterial3: true), home: const ExpenseApp()),
  );
}
