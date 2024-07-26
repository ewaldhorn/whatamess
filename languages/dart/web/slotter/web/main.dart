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
  final slotsList =
      web.document.querySelectorAll('.slot') as List<web.HTMLDivElement>;

  for (var element in slotsList) {
    final fruit = element.querySelector('.fruit') as web.HTMLDivElement;
    final duration = 0.5;
    fruit.style.transform = "translateY(0)";
  }
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

// ------------------------------------------------------------------ clearGame
void clearGame(bool firstClear, {int groups = 1, int duration = 1}) {}

// ----------------------------------------------------------------------- main
void main() {
  initHelp();
  setupSlots();
  setupButtonHandlers();
}
