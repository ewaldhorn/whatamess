// Revisiting my article from June 24, 2019
// https://www.thepolyglotdeveloper.com/2019/06/moving-to-dart-from-cpp/

import 'dart:io';
import 'dart:math';

void main() {
  List numbers = createList();
  print('');
  displayList('Unsorted list', numbers);
  sortList(numbers, 0, numbers.length - 1);
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

void sortList(List values, int left, int right) {
  if (left < right) {
    int pivotIndex = _partition(values, left, right);
    sortList(values, left, pivotIndex - 1);
    sortList(values, pivotIndex, right);
  }
}

int _partition(List numbers, int left, int right) {
  int pivotIndex = left + (right - left) ~/ 2;
  int pivotValue = numbers[pivotIndex];
  int i = left, j = right;
  int temp;
  while (i <= j) {
    while (numbers[i] < pivotValue) {
      i++;
    }
    while (numbers[j] > pivotValue) {
      j--;
    }
    if (i <= j) {
      temp = numbers[i];
      numbers[i] = numbers[j];
      numbers[j] = temp;
      i++;
      j--;
    }
  }
  return i;
}
