import 'dart:io';

void main() {
  // For loop
  print("For loop from 1 to 5:");
  var num = 5;
  for (var i = 1; i <= num; i++) {
    stdout.write(i);

    if (i < num) {
      stdout.write(',');
    }
  }

  // For in loop
  print("\n\nFor In loop:");
  var numbers = [1, 2, 3, 4, 5];

  for (var i in numbers) {
    stdout.write(i);

    if (i < num) {
      stdout.write(',');
    }
  }

  // While loop
  print("\n\nWhile loop:");
  var counter = 1;

  while (counter <= num) {
    stdout.write(counter);

    if (counter < num) {
        stdout.write(',');
    }

    counter++;
  }

  print("\n");
}
