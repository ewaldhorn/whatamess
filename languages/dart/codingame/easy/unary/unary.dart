import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  String MESSAGE = readLineSync();

  var chars = MESSAGE.split('');
  var answer = '';

  List<String> binaries = [];

  for (var char in chars) {
    var asBin = char.codeUnitAt(0).toRadixString(2);

    while (asBin.length < 7) {
      asBin = '0' + asBin;
    }

    asBin.split('').forEach((s) => binaries.add(s));
  }

  stderr.write("We have $binaries\n");

  var pos = 0;
  var c = binaries[pos];
  pos += 1;

  answer += startNewBit(c);

  while (pos < binaries.length) {
    if (c == binaries[pos]) {
      answer += '0';
      pos += 1;
    } else {
      c = binaries[pos];
      answer += ' ';
      answer += startNewBit(c);
      pos += 1;
    }
  }

  print(answer);
}

String startNewBit(String c) {
  if (c == '0') {
    return '00 0';
  } else {
    return '0 0';
  }
}
