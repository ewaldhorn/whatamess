import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  int L = int.parse(readLineSync());
  int H = int.parse(readLineSync());
  List<String> T = readLineSync().toUpperCase().split('');
  List<String> rows = [];
  for (int i = 0; i < H; i++) {
    rows.add(readLineSync());
  }

  List<String> answer = List.filled(H, "");

  int aValue = 'A'.codeUnitAt(0);

  for (var c in T) {
    int pos = c.codeUnitAt(0) - aValue;

    if (pos >= 0) {
      pos = pos * L; // calculate character offset
    } else {
      pos = 26 * L; // last character which should be '?'
    }

    for (var x = 0; x < H; x++) {
      for (var y = 0; y < L; y++) {
        answer[x] += rows[x][pos + y];
      }
    }
  }

  for (var l in answer) {
    print('$l');
  }
}
