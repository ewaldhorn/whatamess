import 'dart:io';
import 'dart:math';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

/**
 * The while loop represents the game.
 * Each iteration represents a turn of the game
 * where you are given inputs (the heights of the mountains)
 * and where you have to print an output (the index of the mountain to fire on)
 * The inputs you are given are automatically updated according to your last actions.
 **/
void main() {
  // game loop
  while (true) {
    int highest = 0;
    int position = 0;

    for (int i = 0; i < 8; i++) {
      int mountainH =
          int.parse(readLineSync()); // represents the height of one mountain.
      if (mountainH > highest) {
        highest = mountainH;
        position = i;
      }
    }

    // Write an action using print()
    // To debug: stderr.writeln('Debug messages...');

    print(position); // The index of the mountain to fire on.
  }
}
