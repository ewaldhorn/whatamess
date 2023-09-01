import 'invoice.dart';

void main(List<String> args) {
  final invoice = Invoice();
  invoice.sumTotal(10);
  invoice.convertToExcel();
  invoice.convertToPDF();
}
