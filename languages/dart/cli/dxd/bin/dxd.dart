import 'dart:io';

import 'package:dxd/dxd.dart';
import 'package:dxd/help.dart';

const hex = 16;

void main(List<String> arguments) async {
  var path = getFileName(arguments);

  if (path != null) {
    readAndDisplayFile(path);
  } else {
    printHelp();
  }
}

void readAndDisplayFile(String path) async {
  var file = File(path);
  var counter = 0;

  final content = await file.readAsBytes();

  StringBuffer countBuffer =
      StringBuffer("${0.toRadixString(hex).padLeft(8, '0')}: ");
  StringBuffer hexCodeBuffer = StringBuffer();
  StringBuffer charBuffer = StringBuffer();

  for (var ch in content) {
    hexCodeBuffer.write(ch.toRadixString(hex).padLeft(2, '0').toUpperCase());
    counter += 1;

    if (ch >= '!'.codeUnitAt(0) && ch <= '~'.codeUnitAt(0)) {
      charBuffer.writeCharCode(ch);
    } else {
      charBuffer.write('.');
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

  if (hexCodeBuffer.isNotEmpty) {
    stdout.writeln(
        '$countBuffer ${hexCodeBuffer.toString().padRight(40, " ")}  $charBuffer');
  }
}
