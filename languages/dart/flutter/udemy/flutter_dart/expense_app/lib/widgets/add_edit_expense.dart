import 'package:expense_app/models/expense.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:expense_app/models/expense.dart' as cats;

final dateFormatter = DateFormat.yMd();

class AddEditExpense extends StatefulWidget {
  const AddEditExpense({super.key, required this.onAddExpense});

  final void Function(Expense expense) onAddExpense;

  @override
  State<AddEditExpense> createState() => _AddEditExpenseState();
}

class _AddEditExpenseState extends State<AddEditExpense> {
  final _descriptionController = TextEditingController();
  final _amountController = TextEditingController();
  DateTime? _selectedDate;
  cats.Category _selectedCategory = cats.Category.leisure;

  @override
  void dispose() {
    _descriptionController.dispose();
    _amountController.dispose();
    super.dispose();
  }

  void _submitExpenseData() {
    final amount = double.tryParse(_amountController.text);
    final amountInvalid = amount == null || amount < 0.0;

    if (_descriptionController.text.trim().isEmpty ||
        amountInvalid ||
        _selectedDate == null) {
      showDialog(
        context: context,
        builder: (context) => AlertDialog(
          title: const Text("Invalid Input"),
          content: const Text("Please make sure you completed all fields."),
          actions: [
            TextButton(
                onPressed: () {
                  Navigator.pop(context);
                },
                child: const Text("OK"))
          ],
        ),
      );
      return;
    }

    var newExpense = Expense(
        _selectedCategory, _descriptionController.text, amount, _selectedDate!);

    widget.onAddExpense(newExpense);
    Navigator.pop(context);
  }

  void _presentDatePicker() async {
    // wait for the date picker to return a selected date
    final selected = await showDatePicker(
        context: context,
        initialDate: DateTime.now(),
        firstDate: DateTime.now().subtract(const Duration(days: 365)),
        lastDate: DateTime.now());

    setState(() {
      _selectedDate = selected;
    });
  }

  @override
  Widget build(BuildContext context) {
    final keyboardHeight = MediaQuery.of(context).viewInsets.bottom;

    return LayoutBuilder(builder: (ctx, constraints) {
      final availableWidth = constraints.maxWidth;

      return SizedBox(
        height: double.infinity,
        child: SingleChildScrollView(
          child: Padding(
            padding: EdgeInsets.fromLTRB(16, 16, 16, keyboardHeight + 16),
            child: Column(children: [
              if (availableWidth >= 600)
                Row(crossAxisAlignment: CrossAxisAlignment.start, children: [
                  Expanded(
                    child: TextField(
                      controller: _descriptionController,
                      maxLength: 100,
                      keyboardType: TextInputType.text,
                      decoration:
                          const InputDecoration(label: Text('Description')),
                    ),
                  ),
                  const SizedBox(width: 24),
                  Expanded(
                    child: TextField(
                      controller: _amountController,
                      maxLength: 10,
                      keyboardType: TextInputType.number,
                      decoration: const InputDecoration(
                          label: Text('Amount'), prefixText: 'R '),
                    ),
                  ),
                ])
              else
                TextField(
                  controller: _descriptionController,
                  maxLength: 100,
                  keyboardType: TextInputType.text,
                  decoration: const InputDecoration(label: Text('Description')),
                ),
              if (availableWidth >= 600)
                Row(
                  children: [
                    DropdownButton(
                        value: _selectedCategory,
                        items: cats.Category.values
                            .map(
                              (entry) => DropdownMenuItem(
                                value: entry,
                                child: Text(entry.name.toUpperCase()),
                              ),
                            )
                            .toList(),
                        onChanged: (value) {
                          if (value == null) {
                            return;
                          }

                          setState(() {
                            _selectedCategory = value;
                          });
                        }),
                    const SizedBox(width: 24),
                    Expanded(
                      child: Row(
                        children: [
                          Text(switch (_selectedDate) {
                            null => 'Select Date',
                            _ => dateFormatter.format(_selectedDate!)
                          }),
                          const SizedBox(width: 10),
                          IconButton(
                            color: Colors.green[900],
                            onPressed: _presentDatePicker,
                            icon: const Icon(Icons.calendar_month_outlined),
                          )
                        ],
                      ),
                    )
                  ],
                )
              else
                Row(
                  children: [
                    Expanded(
                      child: Row(
                        children: [
                          Text(switch (_selectedDate) {
                            null => 'Select Date',
                            _ => dateFormatter.format(_selectedDate!)
                          }),
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
              Row(mainAxisAlignment: MainAxisAlignment.center, children: [
                const SizedBox(width: 20),
                DropdownButton(
                    value: _selectedCategory,
                    items: cats.Category.values
                        .map(
                          (entry) => DropdownMenuItem(
                            value: entry,
                            child: Text(entry.name.toUpperCase()),
                          ),
                        )
                        .toList(),
                    onChanged: (value) {
                      if (value == null) {
                        return;
                      }

                      setState(() {
                        _selectedCategory = value;
                      });
                    }),
                const Spacer(),
                TextButton(
                  onPressed: () {
                    Navigator.pop(context);
                  },
                  child: const Text('Cancel'),
                ),
                const Spacer(),
                ElevatedButton(
                  onPressed: _submitExpenseData,
                  child: const Text('Save Expense'),
                ),
                const SizedBox(width: 20)
              ]),
            ]),
          ),
        ),
      );
    });
  }
}
