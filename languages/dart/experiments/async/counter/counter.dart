import 'dart:async';

// Use a timer to count from 1 to 10

void main() {
  var x = 0;
  var period = const Duration(seconds: 1);
  Timer? timer;

  timer = Timer.periodic(period, (arg) {
    x = x + 1;
    print(x);

    if (x >= 10) {
      timer?.cancel();
    }
  });
}
