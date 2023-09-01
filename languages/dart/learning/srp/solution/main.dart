import 'invoice.dart';
import 'logger.dart';

void main(List<String> args) {
  final invoice = Invoice(logger: Logger());
  invoice.sumTotal(10);
}
