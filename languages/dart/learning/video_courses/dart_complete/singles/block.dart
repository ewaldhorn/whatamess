import 'dart:io';

void main() {
  print('Standard block with default character.');
  drawBlock(10, 5);

  print('\nAnother block, with specified character.');
  drawBlock(10, 5, '#');
}

void drawBlock(int width, int height, [String character = '*']) {
  for (int h = 0; h < height; h++) {
    for (int w = 0; w < width; w++) {
      stdout.write(character);
    }
    stdout.writeln('');
  }
}
