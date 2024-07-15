import 'package:web/web.dart' as web;

// ----------------------------------------------------------------- setHeading
void setHeading() {
  final headingElement =
      web.document.getElementById('heading') as web.HTMLHeadingElement;
  headingElement.text = 'Dart-to-JS Refresher';
}
