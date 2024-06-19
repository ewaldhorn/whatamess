import 'dart:io';
import 'dart:math';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  int rw = int.parse(readLineSync());
  int rh = int.parse(readLineSync());

  List<String> lines = [];

  for (int i = 0; i < rh; i++) {
    lines.add(readLineSync());
  }

  int w = lines[0].length;
  int h = lines.length;

  var answer = List.generate(h, (_) => List.filled(w, 0));

// walk through the grid, if it's an x, increment the neighbours
  for (int height = 0; height < h; height++) {
    for (int width = 0; width < w; width++) {
      if (lines[height][width] == 'x') {
        checkNeighbours(lines, answer, height, width);
      }
    }
  }

  // display grid
  for (int height = 0; height < h; height++) {
    for (int width = 0; width < w; width++) {
      if (answer[height][width] <= 0) {
        stdout.write('.');
      } else {
        stdout.write(answer[height][width]);
      }
    }
    print('');
  }
}

bool canUpdate(List<String> lines, int y, int x) => lines[y][x] != 'x';

void checkNeighbours(List<String> lines, List<List<int>> answer, int y, int x) {
  // can we check up?
  if (y > 0) {
    // top left, above, top right
    if (x > 0) {
      if (canUpdate(lines, y - 1, x - 1))
        answer[y - 1][x - 1] += 1; // top left -1,-1
    }
    if (canUpdate(lines, y - 1, x)) answer[y - 1][x] += 1; // above -1,0

    if (x < lines[0].length - 1) {
      if (canUpdate(lines, y - 1, x + 1))
        answer[y - 1][x + 1] += 1; // top right -1,+1
    }
  }

  if (x > 0) {
    if (canUpdate(lines, y, x - 1)) answer[y][x - 1] += 1; // left 0,-1
  }

  if (x < lines[0].length - 1) {
    if (canUpdate(lines, y, x + 1)) answer[y][x + 1] += 1; // right 0, +1
  }

  if (y < lines.length - 1) {
    if (x > 0) {
      if (canUpdate(lines, y + 1, x - 1))
        answer[y + 1][x - 1] += 1; // bottom left +1, -1
    }
    if (canUpdate(lines, y + 1, x)) answer[y + 1][x] += 1; //below +1, 0
    if (x < lines[0].length - 1) {
      if (canUpdate(lines, y + 1, x + 1))
        answer[y + 1][x + 1] += 1; // bottom right +1, +1
    }
  }
}
