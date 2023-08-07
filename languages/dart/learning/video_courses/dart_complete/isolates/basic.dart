import 'dart:isolate';

void performTask(var message) {
  print('performTask: $message');
}

void main() {
  print('');
  Isolate.spawn(performTask, 'One');
  Isolate.spawn(performTask, 'Two');
  Isolate.spawn(performTask, 'Three');
  Isolate.spawn(performTask, 'Four');
  Isolate.spawn(performTask, 'Five');
  Isolate.spawn(performTask, 'Six');

  for (var i = 0; i < 10; i++) {
    print("${i + 1}. This is the main thread.");
  }
  print('');
}
