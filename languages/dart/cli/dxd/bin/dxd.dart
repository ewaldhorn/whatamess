import 'dart:io';

import 'package:dxd/dxd.dart';
import 'package:dxd/help.dart';

const hex = 16;

// ----------------------------------------------------------------------- main
void main(List<String> arguments) async {
  var path = getFileName(arguments);

  if (path != null) {
    readAndDisplayFile(path);
  } else {
    printHelp();
  }
}

// --------------------------------------------------------- readAndDisplayFile
void readAndDisplayFile(String path) async {
  var file = File(path);
  var counter = 0;

  final content = await file.readAsBytes();

  StringBuffer countBuffer =
      StringBuffer("${0.toRadixString(hex).padLeft(8, '0')}: ");
  StringBuffer hexCodeBuffer = StringBuffer();
  StringBuffer charBuffer = StringBuffer();

  for (var ch in content) {
    counter += 1;
    hexCodeBuffer.write(ch.toRadixString(hex).padLeft(2, '0').toUpperCase());

    if (ch >= '!'.codeUnitAt(0) && ch <= '~'.codeUnitAt(0)) {
      charBuffer.writeCharCode(ch);
    } else {
      charBuffer.write(' ');
    }

    if (counter % 2 == 0) {
      hexCodeBuffer.write(' ');
    }

    if (counter % 16 == 0) {
      stdout.writeln(
          '$countBuffer ${hexCodeBuffer.toString().padRight(40, " ")}  $charBuffer');
      countBuffer.clear();
      countBuffer.write(
          '${counter.toRadixString(hex).padLeft(8, "0").toUpperCase()}: ');
      hexCodeBuffer.clear();
      charBuffer.clear();
    }
  }

  // might have residual characters, display them
  if (hexCodeBuffer.isNotEmpty) {
    stdout.writeln(
        '$countBuffer ${hexCodeBuffer.toString().padRight(40, " ")}  $charBuffer');
  }
}
