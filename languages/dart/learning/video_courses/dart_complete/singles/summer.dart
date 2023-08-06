void main() {
  var target = 10000000;
  // calculate performance of totalling algorithm
  var startTime = DateTime.now();
  var total = sumAllUpTo(target);
  var endTime = DateTime.now();
  var tookTime = endTime.difference(startTime);
  print(
      'It took $tookTime to get a total of $total with the standard algorithm.');

  startTime = DateTime.now();
  total = fasterSumTo(target);
  endTime = DateTime.now();
  tookTime = endTime.difference(startTime);
  print(
      'It took $tookTime to get a total of $total with the improved algorithm.');
}

int sumAllUpTo(int target) {
  var total = 0;

  for (var i = 0; i <= target; i++) {
    total += i;
  }

  return total;
}

int fasterSumTo(int target) {
  return target * (target + 1) ~/ 2;
}
