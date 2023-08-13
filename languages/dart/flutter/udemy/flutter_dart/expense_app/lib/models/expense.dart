import 'package:uuid/uuid.dart';

const uuid = Uuid();

class Expense {
  final String id;
  final String title;
  final double amount;
  final DateTime date;
  final Category category;

  Expense(this.category, this.title, this.amount, this.date) : id = uuid.v4();
}

enum Category { food, travel, work, leisure }
