import 'payment.dart';

void main(List<String> args) {
  final payment = Payment();

  payment.pay(150.00, "cash");
  payment.pay(150.00, "credit");
}
