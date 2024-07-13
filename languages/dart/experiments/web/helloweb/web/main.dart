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
  alignImageProperly();
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

// --------------------------------------------------------- alignImageProperly
// set the parent width to 100%
// set the image to be displayed in a block with auto margins
void alignImageProperly() {
  var el = web.document.querySelector('#tears') as web.HTMLImageElement;
  var parent = el.parentElement as web.HTMLBodyElement;
  parent.style.setProperty('width', '100%');

  el
    ..style.setProperty("display", "block")
    ..style.setProperty("margin-right", "auto")
    ..style.setProperty("margin-left", "auto");
}
