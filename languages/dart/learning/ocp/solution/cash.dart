import 'pay.dart';

class Cash implements Pay {
  @override
  void makePayment(double amount) {
    print('$amount payment made in cash.');
  }
}
