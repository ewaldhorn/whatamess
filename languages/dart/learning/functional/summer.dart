void main() {
  addNumbersTraditional(10, 5);
  addNumbersLambda(10, 5);
}

void addNumbersTraditional(int a, int b) {
  var sum = a + b;
  print(sum);
}

var addNumbersLambda = (int a, int b) {
  var sum = a + b;
  print(sum);
};
