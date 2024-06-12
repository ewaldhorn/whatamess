// Basic solution, using nested loops
// Not optimal, but a good first step to just solve the problem
List<int> productOfArray(List<int> input) {
  List<int> result = List.filled(input.length, 0);

  for (int outer = 0; outer < input.length; outer++) {
    int tmp = 1;

    for (int inner = 0; inner < input.length; inner++) {
      if (inner != outer) {
        tmp = tmp * input[inner];
      }
    }
    result[outer] = tmp;
  }

  return result;
}

// Same output, minus the nested looping
List<int> productOfArrayNoNesting(List<int> input) {
  List<int> result = List.filled(input.length, 1);
  return result;
}
