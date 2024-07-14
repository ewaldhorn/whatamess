int sumOptionals([int? a, int? b, int? c, int? d]) {
  int total = 0;

  total += a ?? 0;
  total += b ?? 0;
  total += c ?? 0;
  total += d ?? 0;

  return total;
}

void main() {
  print('The sum of [1,2] should be 3: ${sumOptionals(1,2)}');
  print('The sum of [1] should be 1: ${sumOptionals(1)}');
  print('The sum of [1,null,4] should be 5: ${sumOptionals(1,null,4)}');
}
