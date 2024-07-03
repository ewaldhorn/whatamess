import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  int W = int.parse(readLineSync());
  int H = int.parse(readLineSync());

  stderr.writeln('W:$W, H:$H');

  List<List<String>> map = [];

  for (int i = 0; i < H; i++) {
    List<String> inputs = readLineSync().split(' ');
    map.add(inputs);
  }

  bool found = false;
  int x = 0, y = 0;

  // first check the corners
  for (var xPos = 0; xPos < W; xPos++) {
    for (var yPos = 0; yPos < H; yPos++) {
      if (map[yPos][xPos] == '0') {
        int count = countObstacles(map, W, H, xPos, yPos);

        if (count > 0) {
          if (count == 3 &&
              ((x == 0 && y == 0) ||
                  (x == W - 1 && y == 0) ||
                  (x == 0 && y == H - 1) ||
                  (x == W - 1 && y == H - 1))) {
            // all the corners
            x = xPos;
            y = yPos;
            found = true;
          } else if (count == 8) {
            // fully surrounded
            x = xPos;
            y = yPos;
            found = true;
          } else if (count == 5) {
            // check if we are on an edge or not
            if (x == 0 || x == W - 1 || y == 0 || y == H - 1) {
              x = xPos;
              y = yPos;
              found = true;
            }
          }
        }

        if (found == true) {
          break;
        }
      }
    }
  }

  print('$x $y');
}

// counts the number of obstacles around the specified position
// if there's another empty block, or the block itself is not empty, -99 results
int countObstacles(List<List<String>> map, int W, int H, int x, int y) {
  int count = 0;

  if (map[y][x] != '0') {
    return -99;
  }

  // not the top row
  if (y > 0) {
    // check right above us
    if (map[y - 1][x] == '1') {
      count += 1;
    } else {
      return -99;
    }

    // left and above
    if (x > 0) {
      if (map[y - 1][x - 1] == '1') {
        count += 1;
      } else {
        return -99;
      }
    }

    // right and above
    if (x < W - 1) {
      if (map[y - 1][x + 1] == '1') {
        count += 1;
      } else {
        return -99;
      }
    }
  }

  // check to the left if possible
  if (x > 0) {
    if (map[y][x - 1] == '1') {
      count += 1;
    } else {
      return -99;
    }
  }

  // check to the right if possible
  if (x < W - 1) {
    if (map[y][x + 1] == '1') {
      count += 1;
    } else {
      return -99;
    }
  }

  // not bottom row
  if (y < H - 1) {
    // below
    if (map[y + 1][x] == '1') {
      count += 1;
    } else {
      return -99;
    }

    // below left
    if (x > 0) {
      if (map[y + 1][x - 1] == '1') {
        count += 1;
      } else {
        return -99;
      }
    }

    // below right
    if (x < W - 1) {
      if (map[y + 1][x + 1] == '1') {
        count += 1;
      } else {
        return -99;
      }
    }
  }
  /*
      -1,-1   0,-1,   1,-1

      -1, 0   0,0     1,0

      -1,1    0,1     1,1
    */

  return count;
}
