import 'dart:io';

import 'logger.dart';

class LoggerImpl extends Logger {
  @override
  void logToFile(String text) {
    final file = File('error_file.log');
    file.writeAsStringSync(text);
  }
}
