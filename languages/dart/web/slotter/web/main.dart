import 'package:web/web.dart' as web;

// -------------------------------------------------------------------- Globals
final List<String> fruits = ['💰', '🥝', '😵', '🍓', '🍒', '🍉', '🍋', '🍍'];

// ------------------------------------------------------------------- initHelp
void initHelp() {
  final helpElement =
      web.document.getElementById('help') as web.HTMLParagraphElement;

  helpElement.innerText = fruits.join(' ');
}

// ----------------------------------------------------------------------- main
void main() {
  initHelp();
}
