import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  int result = 0;
  List inputs;
  int n = int.parse(readLineSync());
  inputs = readLineSync().split(' ');

  if (n > 0) {
    result = int.parse(inputs[0]);

    for (int i = 0; i < n; i++) {
      int t = int.parse(inputs[i]);
      if (t.abs() < result.abs()) {
        result = t;
      } else if (t.abs() == result.abs()) {
        if (result < 0 && t < 0) {
          result = t;
        } else {
          result = t.abs();
        }
      }
    }
  }

  print(result);
}
