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
        if (isSymbol(lines, x, y)) {
          stdout.write(lines[y][x]);
          total += walkLeft(lines, x, y - 1);
          total += walkLeft(lines, x, y);
          total += walkLeft(lines, x, y + 1);
        }
      }
      print('');
    }
  });

  return total;
}

// ------------------------------------------------------------------- isSymbol
bool isSymbol(List<String> list, int x, int y) {
  return !(validNumbers.contains(list[y][x]) || list[y][x] == '.');
}

// -------------------------------------------------------- findAdjacentNumbers
int findAdjacentNumbers(List<String> lines, int x, int y) {
  int total = 0;

  return total;
}

// ------------------------------------------------------------------- walkLeft
// attempt to find a number in the left direction
int walkLeft(List<String> lines, int x, int y) {
  int number = 0;
  int pos = x;

  String buffer = '';

  if (y >= 0 || y <= lines.length) {
    while (pos > 0) {
      pos -= 1;
      if (validNumbers.contains(lines[y][pos])) {
        buffer = lines[y][pos] + buffer;
      } else {
        pos = 0;
      }
    }
  }

  print('Found: \'$buffer\'');

  return number;
}
