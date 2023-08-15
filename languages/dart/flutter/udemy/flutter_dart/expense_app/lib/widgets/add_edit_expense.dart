import 'package:flutter/material.dart';

class AddEditExpense extends StatefulWidget {
  const AddEditExpense({super.key});

  @override
  State<AddEditExpense> createState() => _AddEditExpenseState();
}

class _AddEditExpenseState extends State<AddEditExpense> {
  final _descriptionController = TextEditingController();
  final _amountController = TextEditingController();

  @override
  void dispose() {
    _descriptionController.dispose();
    _amountController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16),
      child: Column(
        children: [
          TextField(
            controller: _descriptionController,
            maxLength: 100,
            keyboardType: TextInputType.text,
            decoration: const InputDecoration(label: Text('Description')),
          ),
          TextField(
            controller: _amountController,
            maxLength: 10,
            keyboardType: TextInputType.number,
            decoration: const InputDecoration(label: Text('Amount')),
          ),
          const SizedBox(height: 10),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              TextButton(onPressed: () {}, child: const Text('Cancel')),
              ElevatedButton(
                  onPressed: () {}, child: const Text('Save Expense'))
            ],
          ),
        ],
      ),
    );
  }
}
