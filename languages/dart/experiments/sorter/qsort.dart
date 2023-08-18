// Revisiting my article from June 24, 2019
// https://www.thepolyglotdeveloper.com/2019/06/moving-to-dart-from-cpp/

import 'dart:io';
import 'dart:math';

void main() {
  List numbers = createList();
  print('');
  displayList('Unsorted list', numbers);
  sortList();
  displayList('\nSorted list  ', numbers);
  print('\n');
}

List createList([int howMany = 30]) {
  var rnd = Random();
  List numbers = List.filled(howMany, 0);
  for (var i = 0; i < howMany; i++) {
    numbers[i] = rnd.nextInt(250);
  }
  return numbers;
}

void displayList(String description, List numbers) {
  stdout.write("$description : ");
  for (var i in numbers) {
    stdout.write("$i ");
  }
}

void sortList() {}
