import 'package:args/args.dart';

const splitColumns = 'split-columns';

// ---------------------------------------------------------------- getFileName
String? getFileName(List<String> arguments) {
  final parser = ArgParser()..addFlag(splitColumns, negatable: false);

  ArgResults argResults = parser.parse(arguments);
  final paths = argResults.rest;

  if (paths.isEmpty) {
    return null;
  } else {
    return paths[0];
  }
}
