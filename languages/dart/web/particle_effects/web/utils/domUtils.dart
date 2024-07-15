import 'package:web/web.dart' as web;

// ------------------------------------------------------------- insertHRBefore
void insertHRBefore(String element) {
  final hrElement = web.HTMLHRElement();
  final outputElement =
      web.document.getElementById(element) as web.HTMLDivElement;
  web.document.body?.insertBefore(hrElement, outputElement);
}
