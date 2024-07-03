import 'dart:io';

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
  if (isNumber(instruction)) {
    // all numbers, need to split off last character
    stdout.write(
        '${instruction[instruction.length - 1] * toNumber(instruction.substring(0, instruction.length - 1))}');
  } else {
    final numericPosition = findNumericPosition(instruction);
    var count = toNumber(instruction.substring(0, numericPosition));
    var chars = instruction.substring(count > 0 ? numericPosition : 0);
    switch (chars) {
      case "nl":
        print('');
        break;
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
        stdout.write('${chars * count}');
    }
  }
}

int findNumericPosition(String string) {
  int pos = 0;

  while (pos < string.length && isNumber(string.substring(0, pos + 1))) {
    pos += 1;
  }

  return pos;
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
