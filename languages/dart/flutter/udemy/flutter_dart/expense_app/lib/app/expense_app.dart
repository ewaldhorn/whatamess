import 'package:expense_app/models/expense.dart';
import 'package:expense_app/widgets/add_edit_expense.dart';
import 'package:expense_app/widgets/chart/chart.dart';
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
      useSafeArea: true,
      isScrollControlled: true,
      builder: (ctx) => AddEditExpense(onAddExpense: _addExpense),
    );
  }

  void _addExpense(Expense e) {
    setState(() {
      _expenses.add(e);
    });
  }

  void _removeExpense(Expense e) {
    final removedItemIndex = _expenses.indexOf(e);

    setState(() {
      _expenses.remove(e);
    });

    ScaffoldMessenger.of(context).clearSnackBars();

    ScaffoldMessenger.of(context).showSnackBar(SnackBar(
      content: Text('Removed "${e.title}".'),
      action: SnackBarAction(
          label: 'Undo',
          onPressed: () {
            setState(() {
              _expenses.insert(removedItemIndex, e);
            });
          }),
    ));
  }

  @override
  Widget build(BuildContext context) {
    final currentWidth = MediaQuery.of(context).size.width;
    Widget mainContent = const Center(child: Text('No expenses yet.'));

    if (_expenses.isNotEmpty) {
      mainContent = ExpensesList(
        expenses: _expenses,
        onRemoveExpense: _removeExpense,
      );
    }

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
      body: currentWidth < 600
          ? Column(
              children: [
                Chart(expenses: _expenses),
                Expanded(
                  child: mainContent,
                )
              ],
            )
          : Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Expanded(child: Chart(expenses: _expenses)),
                Expanded(
                  child: mainContent,
                )
              ],
            ),
    );
  }
}
