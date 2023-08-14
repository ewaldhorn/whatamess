import 'package:expense_app/models/expense.dart';
import 'package:flutter/material.dart';

class ExpenseItem extends StatelessWidget {
  const ExpenseItem(
    this.expense, {
    super.key,
  });

  final Expense expense;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 20.0, vertical: 16.0),
        child: Column(
          children: [
            Text(expense.title),
            const SizedBox(height: 4),
            Row(children: [
              Row(
                children: [
                  Icon(categoryIcons[expense.category]),
                  const SizedBox(width: 8),
                  Text(expense.formattedDate),
                  const SizedBox(width: 8),
                  Text(expense.category.name)
                ],
              ),
              const Spacer(),
              Text('R${expense.amount.toStringAsFixed(2)}')
            ])
          ],
        ),
      ),
    );
  }
}
