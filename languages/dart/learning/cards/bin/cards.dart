typedef Point = ({int x, int y});

void main(List<String> arguments) {
  print('Point is: $Point');
  ({int x, int y}) t = (x: 5, y: 5);

  Point p = (x: 1, y: 2);
  p = t;
}
