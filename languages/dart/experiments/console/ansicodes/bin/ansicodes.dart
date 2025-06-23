import 'dart:io';
import 'dart:async';

// This is a console-based Dart application that allows a user to move an
// asterisk character around the terminal using arrow keys.
// It requires a local Dart environment to run and interact with the console.

void main() async {
  // Ensure the cursor is shown and modes are restored even if the program crashes
  // or exits unexpectedly. This 'finally' block guarantees cleanup.
  StreamSubscription<List<int>>? subscription; // The declaration is here
  try {
    // Hide the cursor (ANSI escape code: CSI ?25l)
    stdout.write('\x1b[?25l');

    // Disable line buffering and echo for immediate character input.
    // This allows the program to read individual key presses without waiting
    // for the Enter key.
    stdin.echoMode = false;
    stdin.lineMode = false;

    // Initial position of the asterisk
    int x = 0; // column
    int y = 0; // row

    // Get the terminal size to prevent the asterisk from going out of bounds.
    // stdout.terminalColumns and stdout.terminalLines might return null
    // in some environments, so provide a fallback.
    final int terminalWidth = stdout.terminalColumns ?? 80;
    final int terminalHeight = stdout.terminalLines ?? 24;

    // Function to clear the console and draw the asterisk at its current position.
    void draw() {
      // Clear the entire screen (ANSI escape code: CSI 2 J)
      stdout.write('\x1b[2J');
      // Move cursor to top-left (ANSI escape code: CSI 1;1 H)
      stdout.write('\x1b[1;1H');
      // Move cursor to the current (x, y) position (ANSI escape code: CSI Y;X H)
      stdout.write('\x1b[${y + 1};${x + 1}H');
      stdout.write('*');
    }

    // Draw the initial state
    draw();

    // Create a stream subscription to listen for raw byte input from stdin.
    // This allows reading individual key presses without blocking the main thread.
    final Completer<void> completer = Completer<void>();

    subscription = stdin.listen((List<int> event) {
      // ANSI escape sequences for arrow keys often start with 0x1B (ESC)
      // followed by other characters.
      if (event.length == 1) {
        // Handle single character inputs (like 'q', Ctrl+D, or ESC)
        final int charCode = event.first;
        if (charCode == 'q'.codeUnits.first ||
            charCode == 0x04 ||
            charCode == 0x1B) {
          // 'q', Ctrl+D (EOT), or ESC
          print('\nExiting...');
          // It's crucial to complete the completer here to allow the 'finally' block to execute.
          // The cleanup will be handled in the 'finally' block.
          completer.complete();
          return;
        }
      } else if (event.length == 3 && event[0] == 0x1B && event[1] == 0x5B) {
        // Handle common 3-byte ANSI escape sequences for arrow keys:
        // ESC [ A for Up (0x1B 0x5B 0x41)
        // ESC [ B for Down (0x1B 0x5B 0x42)
        // ESC [ C for Right (0x1B 0x5B 0x43)
        // ESC [ D for Left (0x1B 0x5B 0x44)
        switch (event[2]) {
          case 0x41: // A - Up arrow
            y = (y - 1).clamp(0, terminalHeight - 1);
            break;
          case 0x42: // B - Down arrow
            y = (y + 1).clamp(0, terminalHeight - 1);
            break;
          case 0x43: // C - Right arrow
            x = (x + 1).clamp(0, terminalWidth - 1);
            break;
          case 0x44: // D - Left arrow
            x = (x - 1).clamp(0, terminalWidth - 1);
            break;
        }
        draw(); // Redraw the asterisk after movement
      }
    });

    // Wait until the program is explicitly told to exit (e.g., by pressing 'q' or Ctrl+D)
    await completer.future;
  } finally {
    // This block will always execute when the 'try' block finishes,
    // whether normally or due to an error.

    // Show the cursor (ANSI escape code: CSI ?25h)
    stdout.write('\x1b[?25h');
    // Re-enable echo mode for the terminal
    stdin.echoMode = true;
    // Re-enable line mode for the terminal
    stdin.lineMode = true;
    // Cancel the stdin subscription to prevent memory leaks if it's still active
    subscription?.cancel(); // The usage is here
    // Clear the screen once more and reset cursor to top-left for clean exit
    stdout.write('\x1b[2J\x1b[1;1H');
  }
  // Exit the application cleanly
  exit(0);
}
