String stringify(int x, int y) => '$x $y';

// Tests your solution (Don't edit!):
void main() {
  assert(stringify(2, 3) == '2 3',
      "Your stringify method returned '${stringify(2, 3)}' instead of '2 3'");
  print('Success!');
}
