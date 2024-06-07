import 'package:dcat/dcat.dart' as dcatlib;

import 'dart:io';

import 'package:args/args.dart';

const lineNumber = 'line-number';

void main(List<String> arguments) {
  exitCode = 0; // presume success
  final parser = ArgParser()..addFlag(lineNumber, negatable: false, abbr: 'n');

  ArgResults argResults = parser.parse(arguments);
  final paths = argResults.rest;

  dcatlib.dcat(paths, showLineNumbers: argResults[lineNumber] as bool);
}
