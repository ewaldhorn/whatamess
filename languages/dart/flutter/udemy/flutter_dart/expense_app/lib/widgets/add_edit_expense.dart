import 'package:flutter/material.dart';

class AddEditExpense extends StatefulWidget {
  const AddEditExpense({super.key});

  @override
  State<AddEditExpense> createState() => _AddEditExpenseState();
}

class _AddEditExpenseState extends State<AddEditExpense> {
  @override
  Widget build(BuildContext context) {
    return const Padding(
      padding: EdgeInsets.all(16),
      child: Column(
        children: [
          TextField(
            maxLength: 100,
            keyboardType: TextInputType.text,
            decoration: InputDecoration(label: Text('Description')),
          ),
        ],
      ),
    );
  }
}
