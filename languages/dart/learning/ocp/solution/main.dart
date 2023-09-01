import 'card.dart';
import 'cash.dart';
import 'payment.dart';

void main(List<String> args) {
  final payment = Payment();
  payment.makePayment(10.20, Cash());
  payment.makePayment(11.50, Card());
}
