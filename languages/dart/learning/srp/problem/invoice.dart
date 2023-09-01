// Not adhering to the single responsibility principle

import 'dart:io';

class Invoice {
  double sumTotal(double amount) {
    double total = 0.0;
    try {
      total = amount + sumTax();
    } catch (e) {
      final file = File('error.log');
      file.writeAsStringSync(e.toString());
    }
    return total;
  }

  double sumTax() {
    return 12.30;
  }

  void convertToExcel() {
    print('convert to excel');
  }

  void convertToPDF() {
    print('convert to pdf');
  }

  void print(String text) {
    print(text);
  }
}
