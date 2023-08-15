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

  void _presentDatePicker() {
    showDatePicker(
        context: context,
        initialDate: DateTime.now(),
        firstDate: DateTime.now().subtract(const Duration(days: 365)),
        lastDate: DateTime.now());
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
          Row(
            children: [
              Expanded(
                child: Row(
                  children: [
                    const Text('Selected Date'),
                    const SizedBox(width: 10),
                    IconButton(
                      color: Colors.green[900],
                      onPressed: _presentDatePicker,
                      icon: const Icon(Icons.calendar_month_outlined),
                    )
                  ],
                ),
              ),
              const SizedBox(width: 20),
              Expanded(
                child: TextField(
                  controller: _amountController,
                  maxLength: 10,
                  keyboardType: TextInputType.number,
                  decoration: const InputDecoration(
                      label: Text('Amount'), prefixText: 'R '),
                ),
              ),
            ],
          ),
          const SizedBox(height: 20),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const SizedBox(width: 20),
              TextButton(
                onPressed: () {
                  Navigator.pop(context);
                },
                child: const Text('Cancel'),
              ),
              const Spacer(),
              ElevatedButton(
                onPressed: () {},
                child: const Text('Save Expense'),
              ),
              const SizedBox(width: 20)
            ],
          ),
        ],
      ),
    );
  }
}
