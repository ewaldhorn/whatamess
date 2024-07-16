import 'dart:math';

BigInt square(final int n) {
  BigInt result = BigInt.from(2);
  
  if (n < 1 || n > 64) {
    throw ArgumentError('square must be between 1 and 64');
  }

  if (n == 1) {
    return BigInt.from(1);
  } else {
    return result.pow(n-1);
  }
}

BigInt total() {
  BigInt result = BigInt.from(0);

  for (int g = 1; g <= 64; g++) {
    result = result + square(g);
  }
  
  return result;
}

