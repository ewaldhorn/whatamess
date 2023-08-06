import 'dart:io';

void main() {
  print('Standard block with default character.');
  drawBlock(10, 5);

  print('\nAnother block, with specified character.');
  drawBlock(10, 5, '#');

  print('\nDraw an outline of a block with the default character.');
  drawOutline(10, 5);

  print('\nDraw another outline, with a custom character.');
  drawOutline(10, 5, '=');
}

void drawBlock(int width, int height, [String character = '*']) {
  for (int h = 0; h < height; h++) {
    for (int w = 0; w < width; w++) {
      stdout.write(character);
    }
    stdout.writeln('');
  }
}

void drawOutline(int width, int height, [String character = '*']) {
  for (int h = 0; h < height; h++) {
    for (int w = 0; w < width; w++) {
      if ((h == 0 || h == (height - 1)) || (w == 0 || w == (width - 1))) {
        stdout.write(character);
      } else {
        stdout.write(' ');
      }
    }
    stdout.writeln('');
  }
}
