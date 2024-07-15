class Snake {
  // --------------------------------------------------------------- Directions
  static const ({int x, int y}) left = (x: -1, y: 0);
  static const ({int x, int y}) right = (x: 1, y: 0);
  static const ({int x, int y}) up = (x: 0, y: -1);
  static const ({int x, int y}) down = (x: 0, y: 1);

  // --------------------------------------------------------- Class properties
  ({int x, int y}) _direction = right;
  late List<({int x, int y})> _body;

  // -------------------------------------------------------------- Constructor
  Snake(int initialLength) {
    _body = List<({int x, int y})>.generate(initialLength, (index) {
      ({int x, int y}) part = (x: index, y: 0);
      return part;
    });
  }

  ({int x, int y}) get head => _body.first;

  // ---------------------------------------------------------- changeDirection
  void changeDirection(({int x, int y}) newDirection) {
    _direction = newDirection;
  }

  // ---------------------------------------------------------- reportDirection
  // used for debugging snake movement
  void reportDirection() {
    print('I\'m going $_direction');
  }
}
