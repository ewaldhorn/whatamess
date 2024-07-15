import 'dart:collection';
import 'package:web/web.dart' as web;

class Keyboard {
  final HashMap<int, num> _keys = HashMap<int, num>();

  Keyboard() {
    web.window.onKeyPress.listen((web.KeyboardEvent event) {
      _keys.clear();
      _keys.putIfAbsent(event.keyCode, () => event.timeStamp);
    });

    // web.window.onKeyUp.listen((web.KeyboardEvent event) {
    //   _keys.remove(event.keyCode);
    // });
  }

  bool isPressed(int keyCode) => _keys.containsKey(keyCode);
}
