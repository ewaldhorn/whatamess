// Closely adheres to the single responsibility principle
import 'logger.dart';

class Invoice {
  final Logger logger;

  Invoice({required this.logger});

  double sumTotal(double amount) {
    double total = 0.0;
    try {
      total = amount + sumTax();
    } catch (e) {
      logger.logToFile(e.toString());
    }
    return total;
  }

  double sumTax() {
    return 12.30;
  }
}
