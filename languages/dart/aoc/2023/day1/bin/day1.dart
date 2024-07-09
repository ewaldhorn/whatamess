import 'dart:io';

// ----------------------------------------------------------------------- main
void main(List<String> arguments) async {
  if (arguments.isEmpty) {
    print('Please specify an input file.');
  } else {
    int val = await parseFile(arguments.first);
    print('Total is: $val');
  }
}

// ------------------------------------------------------------------ parseFile
Future<int> parseFile(String filename) async {
  int total = 0;

  await File(filename).readAsLines().then((lines) {
    for (var line in lines) {
      total += parseLine(line);
    }
  });

  return total;
}

// ------------------------------------------------------------------ parseLine
int parseLine(String line) {
  String? left, right;

  int len = line.length - 1;
  int zeroChar = '0'.codeUnitAt(0);
  int nineChar = '9'.codeUnitAt(0);

  for (var i = 0; i < line.length; i++) {
    if (left == null &&
        line[i].codeUnitAt(0) >= zeroChar &&
        line[i].codeUnitAt(0) <= nineChar) {
      left = line[i];
    }

    if (right == null &&
        line[len - i].codeUnitAt(0) >= zeroChar &&
        line[len - i].codeUnitAt(0) <= nineChar) {
      right = line[len - i];
    }
  }

  return int.parse('$left$right');
}
