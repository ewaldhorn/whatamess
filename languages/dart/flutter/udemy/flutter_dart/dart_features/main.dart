void main() {
  print('');
  print('Some general Dart features:');

  MyClass c = MyClass(1);

  print('Class C has ${c.value} and ${c.doubleValue}.');
  print('');
}

class MyClass {
  final int value;

  const MyClass(this.value);

  int get doubleValue => value * 2;
}
