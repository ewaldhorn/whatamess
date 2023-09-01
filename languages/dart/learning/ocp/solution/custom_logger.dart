import 'logger.dart';

class CustomLogger extends Logger {
  @override
  void logToFile(String text) {
    print('This should log "$text"');
  }
}
