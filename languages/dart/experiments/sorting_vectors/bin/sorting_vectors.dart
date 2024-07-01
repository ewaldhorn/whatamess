import 'dart:io';
import 'dart:math';

// Slightly updated code for https://www.thepolyglotdeveloper.com/2019/06/moving-to-dart-from-cpp/

void main() {
  print('');
  List numbers = createList(20);
  displayList("Unsorted list", numbers);
  sortList(numbers, 0, numbers.length - 1);
  displayList("\nSorted list  ", numbers);
  print('\n');
}

List createList(int howMany) {
  var rnd = Random();
  List numbers = List.empty(growable: true);

  for (var i = 0; i < howMany; i++) {
    numbers.add(rnd.nextInt(250));
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
  int pivotIndex = left +
      (right - left) ~/
          2; // notice the ~/ operator that casts the result to int
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
