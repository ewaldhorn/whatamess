import 'dart:js_interop';

import 'package:web/web.dart' as web;

// -------------------------------------------------------------------- Globals
final List<String> fruits = ['💰', '🥝', '😵', '🍓', '🍒', '🍉', '🍋', '🍍'];

// ------------------------------------------------------------------- initHelp
void initHelp() {
  final helpElement =
      web.document.getElementById('help') as web.HTMLParagraphElement;

  helpElement.innerText = fruits.join(' ');
}

// ----------------------------------------------------------------- setupSlots
void setupSlots() {
  final slotsElement =
      web.document.querySelector('.slots') as web.HTMLDivElement;
}

// -------------------------------------------------------- setupButtonHandlers
void setupButtonHandlers() {
  final spinButton =
      web.document.getElementById('spin') as web.HTMLButtonElement;
  spinButton.addEventListener('click', handleSpinClicked.toJS);

  final clearButton =
      web.document.getElementById('clear') as web.HTMLButtonElement;
  clearButton.addEventListener('click', handleClearClicked.toJS);
}

// ---------------------------------------------------------- handleSpinClicked
void handleSpinClicked(web.Event event) {
  web.console.log('Spin'.toJS);
}

// --------------------------------------------------------- handleClearClicked
void handleClearClicked(web.Event event) {
  web.console.log('Clear'.toJS);
}

// ----------------------------------------------------------------------- main
void main() {
  initHelp();
  setupSlots();
  setupButtonHandlers();
}
