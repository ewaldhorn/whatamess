// ----------------------------------------------------------------------- main
import 'dart:io';

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
  int colonIndex = line.indexOf(':');
  List<String> resultsPart = line.substring(colonIndex + 2).split(';');

  return calculateLeastNumbersNeeded(resultsPart);
}

// ------------------------------------------------ calculateLeastNumbersNeeded
int calculateLeastNumbersNeeded(List<String> results) {
  int maxRed = 0;
  int maxGreen = 0;
  int maxBlue = 0;

  for (final result in results) {
    List<String> parts = result.trim().split(',');

    for (final part in parts) {
      List<String> detail = part.trim().split(' ');

      switch (detail[1]) {
        case 'red':
          int tmpRed = int.parse(detail[0]);
          if (tmpRed > maxRed) {
            maxRed = tmpRed;
          }
        case 'green':
          int tmpGreen = int.parse(detail[0]);
          if (tmpGreen > maxGreen) {
            maxGreen = tmpGreen;
          }
        case 'blue':
          int tmpBlue = int.parse(detail[0]);
          if (tmpBlue > maxBlue) {
            maxBlue = tmpBlue;
          }
      }
    }
  }

  return maxRed * maxGreen * maxBlue;
}
