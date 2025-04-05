void main() {
  final str = 'Hello, World!';
  final step = 2;
  final slice = StringBuffer();

  for (int i = 0; i < str.length; i += step) {
    slice.write(str[i]);
  }

  print(slice.toString()); // prints 'Hlo ol!'
}