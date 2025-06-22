import 'package:dart_console/dart_console.dart';
import 'dart:io'; // For sleep and stdin/stdout

void main() {
  final console = Console();

  // Clear the screen and reset cursor
  console.clearScreen();
  console.resetCursorPosition();

  // Set colors
  console.setBackgroundColor(ConsoleColor.blue);
  console.setForegroundColor(ConsoleColor.white);

  // Print a title
  console.writeLine('Dart PONG (WIP)', TextAlignment.center);
  console.resetColorAttributes(); // Reset colors to default

  console.writeLine(''); // Empty line

  // Simulate a game loop - this is very basic!
  // In a real PONG, you'd have continuous drawing and input handling.
  for (int i = 0; i < 15; i++) {
    console.cursorPosition = Coordinate(5 + i, 10); // Move cursor
    console.write('Paddle A');
    console.cursorPosition = Coordinate(
      5 + i,
      console.windowWidth - 10,
    ); // Move cursor
    console.write('Paddle B');
    console.cursorPosition = Coordinate(8, 15 + i * 2); // Move ball
    console.write('O'); // Ball
    sleep(Duration(milliseconds: 500)); // Pause for a short time
    console
        .clearScreen(); // Clear for next frame (inefficient for games, but simple)
    console.resetCursorPosition();
    console.writeLine('Dart PONG (WIP)', TextAlignment.center); // Redraw title
  }

  console.cursorPosition = Coordinate(
    console.windowHeight - 2,
    0,
  ); // Move to bottom
  console.writeLine('Press any key to exit...');
  stdin.readLineSync(); // Wait for user input before closing
  console.resetColorAttributes(); // Ensure terminal colors are reset on exit
  console.clearScreen(); // Clear screen on exit
}
