import 'dart:io';

class Logger {
  void logToFile(String text) {
    final file = File('error.log');
    file.writeAsStringSync(text);
  }
}
