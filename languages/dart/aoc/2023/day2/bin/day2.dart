// ----------------------------------------------------------------------- main
import 'dart:io';

const maxRed = 12;
const maxGreen = 13;
const maxBlue = 14;

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
  String gamePart = line.substring(0, colonIndex);
  String gameNumber = gamePart.substring(gamePart.indexOf(' ') + 1);
  List<String> resultsPart = line.substring(colonIndex + 2).split(';');

  print('Game Part: [$gamePart] [$gameNumber] - Results: [$resultsPart]');

  for (final result in resultsPart) {
    if (!isPossibleResult(result)) {
      return 0;
    }
  }

  return int.parse(gameNumber);
}

// ----------------------------------------------------------- isPossibleResult
bool isPossibleResult(String result) {
  return true;
}
