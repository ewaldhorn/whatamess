import 'dart:io';

Map<String, String> validNumbers = {
  '1': '1',
  '2': '2',
  '3': '3',
  '4': '4',
  '5': '5',
  '6': '6',
  '7': '7',
  '8': '8',
  '9': '9',
  'one': '1',
  'two': '2',
  'three': '3',
  'four': '4',
  'five': '5',
  'six': '6',
  'seven': '7',
  'eight': '8',
  'nine': '9'
};

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

  // left
  for (var i = 0; i < line.length; i++) {
    var tmp = line.substring(i);
    validNumbers.forEach((k, v) {
      if (left == null && tmp.startsWith(k)) {
        left = v;
      }
    });
  }

  // right
  for (var i = line.length; i > 0; i--) {
    var tmp = line.substring(0, i);
    validNumbers.forEach((k, v) {
      if (right == null && tmp.endsWith(k)) {
        right = v;
      }
    });
  }

  return int.parse('$left$right');
}
