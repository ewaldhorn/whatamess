import 'package:web/web.dart' as web;

web.HTMLLIElement newLI(String itemText) =>
    (web.document.createElement('li') as web.HTMLLIElement)..text = itemText;

void main() {
  final now = DateTime.now();
  final element = web.document.querySelector('#output') as web.HTMLDivElement;
  element.text = 'The time is ${now.hour}:${now.minute} '
      'and your Dart web app is running!';

  final output = web.document.querySelector('#output');
  for (final item in thingsTodo()) {
    output?.appendChild(newLI(item));
  }
}

Iterable<String> thingsTodo() sync* {
  const actions = ['Walk', 'Wash', 'Feed'];
  const pets = ['cats', 'dogs'];

  for (final action in actions) {
    for (final pet in pets) {
      if (pet != 'cats' || action == 'Feed') {
        yield '$action the $pet';
      }
    }
  }
}
