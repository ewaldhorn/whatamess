import 'package:web/web.dart' as web;

// ----------------------------------------------------------- getOutputElement
web.HTMLDivElement getOutputElement() =>
    web.document.querySelector('#output') as web.HTMLDivElement;

web.HTMLDivElement getTitleElement() =>
    web.document.querySelector('#title') as web.HTMLDivElement;

// --------------------------------------------------------------- createNewDiv
web.HTMLDivElement createNewDiv(String? cssClass) {
  final newDiv = web.document.createElement('div') as web.HTMLDivElement;

  if (cssClass != null) {
    newDiv.classList.add(cssClass);
  }

  return newDiv;
}

// ---------------------------------------------------------------- addNewDivTo
web.HTMLDivElement addNewDivTo(
    web.HTMLElement parent, String? html, String? cssClass) {
  final newDiv = createNewDiv(cssClass);

  if (html != null) {
    newDiv.innerHTML = html;
  }

  parent.append(newDiv);

  return newDiv;
}
