import 'dart:io';

// ----------------------------------------------------------------------------
void main(List<String> args) {
  if (args.length != 1) {
    print('Usage: dart ${Platform.script.path} string');
    return;
  }
  print(toLeet(args[0]));
}

// ----------------------------------------------------------------------------
String toLeet(String inputString) {
  final replacements = {
    'a': '4',
    'A': '4',
    'e': '3',
    'E': '3',
    'i': '1',
    'I': '1',
    'o': '0',
    'O': '0',
    's': '5',
    'S': '5',
    't': '7',
    'T': '7',
  };
  return replacements.entries.fold(
    inputString,
    (acc, entry) => acc.replaceAll(entry.key, entry.value),
  );
}
