import 'dart:async';

void main() async {
  // Create a StreamController with String events
  StreamController<String> controller = StreamController<String>();

  // Get the sink to add events to the stream
  Sink<String> sink = controller.sink;

  // Listen to events from the stream
  controller.stream.listen((data) {
    print('Received: $data');
  });

  for (int i = 1; i <= 10; i++) {
    sink.add('Event $i');
    await Future.delayed(Duration(milliseconds: 500));
  }

  // Close the stream when done
  sink.close();
}
