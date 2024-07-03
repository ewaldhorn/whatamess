import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  List inputs = readLineSync().split(' ');
  int W = int.parse(inputs[0]);
  int H = int.parse(inputs[1]);

  List<String> lines = [];

  for (int i = 0; i < H; i++) {
    lines.add(readLineSync());
  }

  lines[0].split('').asMap().forEach((key, value) {
    if (value != ' ') {
      var xpos = key;
      var ypos = 1;

      while (ypos < H - 1) {
        if (xpos > 0) {
          // check if we can go LEFT
          if (lines[ypos][xpos - 1] == '-') {
            xpos -= 3;
          } else if (xpos < W - 1 && lines[ypos][xpos + 1] == '-') {
            xpos += 3;
          }
        } else if (xpos == 0) {
          if (lines[ypos][xpos + 1] == '-') {
            xpos += 3;
          }
        } else {
          if (xpos < W - 1) {
            if (lines[ypos][xpos + 1] == '-') {
              xpos += 3;
            }
          }
        }

        ypos++;
      }

      print('${value}${lines[ypos][xpos]}');
    }
  });
}
