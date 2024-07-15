import 'package:web/web.dart' as web;

// ----------------------------------------------------------------- setHeading
void setHeading() {
  final headingElement =
      web.document.getElementById('heading') as web.HTMLHeadingElement;
  headingElement.text = 'Dart-to-JS Experiments';
}

// ------------------------------------------------------------- insertHRBefore
void insertHRBefore(String element) {
  final hrElement = web.HTMLHRElement();
  final outputElement =
      web.document.getElementById(element) as web.HTMLDivElement;
  web.document.body?.insertBefore(hrElement, outputElement);
}

// ----------------------------------------------------------------------- main
void main() {
  setHeading();
  insertHRBefore('output');
}
