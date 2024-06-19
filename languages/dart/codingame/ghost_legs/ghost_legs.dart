import 'dart:io';
import 'dart:math';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  List inputs;
  inputs = readLineSync().split(' ');
  int W = int.parse(inputs[0]);
  int H = int.parse(inputs[1]);
  for (int i = 0; i < H; i++) {
    String line = readLineSync();
  }

  print('answer');
}
