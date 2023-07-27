import "dart:io";

void main() {
    stdout.write("Who goes there?: ");
    var name = stdin.readLineSync();
    print("Oh, hello there $name!\nWelcome!");

    stdout.write("Enter a number (1 - 100):");
    String? input = stdin.readLineSync();
    int choice = int.parse(input ?? "0");
    print("Your choice is $choice");
}
