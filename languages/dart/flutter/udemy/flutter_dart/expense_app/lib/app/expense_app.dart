import 'package:expense_app/models/expense.dart';
import 'package:flutter/material.dart';

class ExpenseApp extends StatefulWidget {
  const ExpenseApp({super.key});

  @override
  State<ExpenseApp> createState() => _ExpenseAppState();
}

class _ExpenseAppState extends State<ExpenseApp> {
  final List<Expense> _expenses = [
    Expense(Category.food, "Big Mac", 65.95,
        DateTime.now().add(const Duration(hours: 11))),
    Expense(Category.travel, "Petrol", 455.78, DateTime.now()),
    Expense(Category.work, "Flutter Course", 78.00,
        DateTime.now().subtract(const Duration(days: 3, hours: 2, minutes: 7)))
  ];
  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Column(
        children: [Text('Chart'), Text('Expenses')],
      ),
    );
  }
}
