import 'dart:io';
import 'dart:math';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

const name_pos = 1;
const lon_pos = 4;
const lat_pos = 5;

void main() {
  String LON = readLineSync();
  String LAT = readLineSync();
  int N = int.parse(readLineSync());

  // read all the defib entry lines as a list of strings, replacing the , with . where needed
  List<List<String>> fibs = [];

  for (int i = 0; i < N; i++) {
    fibs.add(readLineSync().replaceAll(',', '.').split(';'));
  }

  // set up our temp vars
  var name = fibs[0][name_pos];
  var closest = 9999999999.99;
  var plon = double.parse(LON.replaceAll(',', '.')); // get position lon
  var plat = double.parse(LAT.replaceAll(',', '.')); // get position lat

  for (var pos = 0; pos < fibs.length; pos++) {
    var lon = double.parse(fibs[pos][lon_pos]);
    var lat = double.parse(fibs[pos][lat_pos]);
    var tmp = distanceBetween(plon, plat, lon, lat);

    if (tmp < closest) {
      closest = tmp;
      name = fibs[pos][name_pos];
    }
  }

  print('$name');
}

// calculate distance between two points as per problem statement
double distanceBetween(double x1, double y1, double x2, double y2) {
  var x = (x2 - x1) * cos((y1 + y2) / 2);
  var y = (y2 - y1);
  return sqrt(6371.0 * (pow(x, 2) + pow(y, 2)));
}
