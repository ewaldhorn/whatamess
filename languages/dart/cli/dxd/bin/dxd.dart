import 'dart:convert';
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
  stdout.write("${0.toRadixString(hex).padLeft(8, '0')} ");

  for (var ch in content) {
    stdout.write(ch.toRadixString(hex).padLeft(2, '0').toUpperCase());
    counter += 1;

    if (counter % 2 == 0) {
      stdout.write(' ');
    }

    if (counter % 16 == 0) {
      stdout.write(
          '\n${counter.toRadixString(hex).padLeft(8, "0").toUpperCase()} ');
    }
  }
}
