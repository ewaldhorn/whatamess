void main(List<String> args) {
  final loc = location();
  print(loc);

  final latLong = getLatLon();
  print("Lat: ${latLong.lat}, Lon: ${latLong.lon}");

  final (lat: newLatName, lon: newLonName) = getLatLon();
  print("Lat: $newLatName, Lon: $newLonName");
}

({int lat, int lon}) getLatLon() {
  return (lat: 3, lon: 4);
}

Location location() {
  return Location(1, 2);
}

class Location {
  Location(this.x, this.y);

  final int x, y;

  @override
  String toString() {
    return 'Location{x: $x, y: $y}';
  }
}
