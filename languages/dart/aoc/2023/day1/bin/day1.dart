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
      if (line.isNotEmpty) {
        total += parseLine(line);
      }
    }
  });

  return total;
}

// ------------------------------------------------------------------ parseLine
int parseLine(String line) {
  int total = 1;
  return total;
}
