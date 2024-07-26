class DifferenceOfSquares {
  int squareOfSum(int num) {
    int total = 0;

    for (var i = 1; i <= num; i++) {
      total += i;
    }

    return total * total;
  }

  int sumOfSquares(int num) {
    int total = 0;

    for (var i = 1; i <= num; i++) {
      total += i * i;
    }

    return total;
  }

  int differenceOfSquares(int num) {
    return squareOfSum(num) - sumOfSquares(num);
  }
}
