// ----------------------------------------------------------------------------
// use two loops to control the sliding window position
int findMaxSumOfSubArrayLooped(List<int> numbers, int windowSize) {
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

// ----------------------------------------------------------------------------
// only iterate over the numbers list once, using a sliding window
// as the window moves, subtract the number leaving the window and then
// add the number entering the window (start, end)
int findMaxSumOfSubArray(List<int> numbers, int windowSize) {
  // still don't bother if we don't have enough numbers to start with
  if (numbers.length < windowSize) {
    return 0;
  }

  int maxSum = 0, start = 0, end = 0, windowSum = 0;

  // calculate the initial window sum (first [windowSize] numbers)
  while (end < windowSize) {
    windowSum += numbers[end];
    end++;
  }
  maxSum = windowSum;

  // now perform the sliding window operation to check the rest.
  while (start + windowSize < numbers.length) {
    windowSum = windowSum - numbers[start] + numbers[end];

    if (windowSum > maxSum) {
      maxSum = windowSum;
    }

    start++;
    end++;
  }

  return maxSum;
}
