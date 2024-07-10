import 'dart:io';

final validNumbers = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];

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
    int height = lines.length;
    int width = lines[0].length;

    for (int y = 0; y < height; y++) {
      for (int x = 0; x < width; x++) {
        stdout.write(lines[y][x]);
      }
      print('');
    }
  });

  return total;
}

// -------------------------------------------------------- findAdjacentNumbers
int findAdjacentNumbers(List<String> lines, int x, int y) {
  int total = 0;

  return total;
}
