int findMaxSumOfSubArray(List<int> numbers, int windowSize) {
  int maxSum = 0;

  // don't bother if we don't have enough numbers to start with
  if (numbers.length < windowSize) {
    return 0;
  }

  for (int pos = 0; pos < numbers.length - windowSize; pos++) {
    int tmp = 0;

    for (int windowPos = pos; windowPos < pos + windowSize; windowPos++) {
      tmp += numbers[windowPos];
    }

    if (tmp > maxSum) {
      maxSum = tmp;
    }
  }

  return maxSum;
}
