import 'package:web/web.dart' as web;
import 'helpers/dom_helpers.dart' as helpers;
import 'types.dart';

// -------------------------------------------------------------------- initApp
void initApp() {
  final Size gridSize = (width: 20, height: 10);
  // get output element
  final outputElement = helpers.getOutputElement();
  outputElement.innerHTML = '';

  final titleElement = helpers.getTitleElement();
  titleElement.innerHTML =
      '<h3>The Grid (${gridSize.width} x ${gridSize.height})</h3>';

  // create main grid container and add to output
  final mainElement = helpers.addNewDivTo(outputElement, null, 'gridContainer');

  buildGrid(mainElement, gridSize);

  mainElement.style.gridTemplateColumns = ' auto ' * gridSize.width;
  mainElement.style.gridTemplateRows = ' auto ' * gridSize.height;
  mainElement.style.gridGap = '10px';
}

// ------------------------------------------------------------------ buildGrid
void buildGrid(web.HTMLDivElement mainElement, Size gridSize) {
  for (int height = 0; height < gridSize.height; height++) {
    for (int width = 0; width < gridSize.width; width++) {
      helpers.addNewDivTo(
          mainElement, '${(height * gridSize.width) + width + 1}', 'grid-item');
    }
  }
}

void main() {
  initApp();
}
