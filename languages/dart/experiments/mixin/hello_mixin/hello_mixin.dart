mixin AreaFunctionsMixin {
  int areaOfSquare(int side) {
    return side * side;
  }
}

class MySquare with AreaFunctionsMixin {
  final int side;

  const MySquare(this.side);

  void reportAreaSize() {
    print(
        'A square with sides sized $side has an area of ${areaOfSquare(side)}.');
  }
}

void main() {
  MySquare mySquare = MySquare(8);
  mySquare.reportAreaSize();
}
