import 'package:expense_app/models/expense.dart';
import 'package:expense_app/widgets/add_edit_expense.dart';
import 'package:expense_app/widgets/expenses_list/expenses_list.dart';
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
    Expense(
      Category.travel,
      "Petrol",
      455.78,
      DateTime.now(),
    ),
    Expense(
      Category.work,
      "Flutter Course",
      78.00,
      DateTime.now().subtract(const Duration(days: 3, hours: 2, minutes: 7)),
    )
  ];

  void _showAddExpenseModal() {
    showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      builder: (ctx) => AddEditExpense(onAddExpense: _addExpense),
    );
  }

  void _addExpense(Expense e) {
    setState(() {
      _expenses.add(e);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Flutter Expense Tracker'),
        actions: [
          IconButton(
            onPressed: _showAddExpenseModal,
            icon: const Icon(Icons.plus_one),
          )
        ],
      ),
      body: Column(
        children: [
          const Text('Chart'),
          Expanded(child: ExpensesList(expenses: _expenses))
        ],
      ),
    );
  }
}
