import 'pay.dart';

class Card implements Pay {
  @override
  void makePayment(double amount) {
    print('$amount paid with card.');
  }
}
