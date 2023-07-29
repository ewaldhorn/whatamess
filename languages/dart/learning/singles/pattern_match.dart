extension DescribeDate on DateTime {
  void describe() {
    final now = DateTime.now();
    final diff = this.difference(DateTime(now.year, now.month, now.day));
    String description = switch(diff) {
      Duration(inDays: -1) => 'yesterday',
      Duration(inDays: 0) => 'today',
      Duration(inDays: 1) => 'tomorrow',
      Duration(inDays: 2) => 'the day after tomorrow',
      Duration(inDays: int d) => '${d} days from now'
    };

    print('$year/$month/$day is $description');
  }
}

void main() {
  DateTime(2023, 7, 28).describe(); // yesterday
  DateTime(2023, 7, 29).describe(); // today
  DateTime(2023, 7, 30).describe(); // tomorrow
  DateTime(2023, 7, 31).describe(); // the day after tomorrow
  DateTime(2023, 8, 1).describe(); // 3 days from now
}
