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
  // set up the results list, populate it with 1's
  List<int> result = List.filled(input.length, 1);

  // prepopulate the result with the multipliers for that position
  for (int pos = 1; pos < input.length; pos++) {
    result[pos] = input[pos - 1] * result[pos - 1];
  }

  // now loop through the numbers, calculating the product
  int tmp = 1;

  for (int pos = input.length; pos > 0; pos--) {
    result[pos - 1] = result[pos - 1] * tmp;
    tmp = tmp * input[pos - 1];
  }
  return result;
}
