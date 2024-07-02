import 'dart:io';
import 'dart:math';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  String LON = readLineSync();
  String LAT = readLineSync();
  int N = int.parse(readLineSync());

  List<String> fibs = [];

  for (int i = 0; i < N; i++) {
    fibs.add(readLineSync().replaceAll(',', '.'));
  }

  print('answer');
}
