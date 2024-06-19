import 'dart:io';
import 'dart:math';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  String T = readLineSync();

  for (var instruction in T.split(' ')) {
    performInstruction(instruction);
  }
}

void performInstruction(String instruction) {
  if (instruction == 'nl') {
    // new line
    print('');
  } else if (isNumber(instruction)) {
    // all numbers, need to split off last character
    stdout.write(
        '${instruction[instruction.length - 1] * toNumber(instruction.substring(0, instruction.length - 1))}');
  } else {
    final numeric = RegExp(r'\d+/g');
    final chars = RegExp(r'[^0-9]/g');
    var t = numeric.firstMatch(instruction);
    var t2 = chars.firstMatch(instruction);
    var count = 0;

    if (t != null) {
      count = int.parse(t.toString());
    }

    if (t2 != null) {
      switch (t2.toString()) {
        case "sp":
          stdout.write('${" " * count}');
          break;
        case "bS":
          stdout.write('${"\\" * count}');
          break;
        case "sQ":
          stdout.write('${"'" * count}');
          break;
        default:
          stdout.write('${t2.toString() * count}');
      }
    }
  }
}

int toNumber(String string) {
  int result = 0;

  try {
    result = int.parse(string);
  } catch (ex) {}

  return result;
}

bool isNumber(String string) {
  try {
    int.parse(string);
    return true;
  } catch (ex) {
    return false;
  }
}
