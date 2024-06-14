// Verifies the Dart environment is working.

// Run with:
// dart run verify_dart_environment.dart

// ----------------------------------------------------------------------- main
void main() {
  print('Hello, Dart!');
  print('');

  var someItems = ['Item 1', 'Another item', 'The last item'];
  listItems(someItems);

  print('');
  print('The 9th Fibonacci number is: ${fibonacci(9)}');

  flightDataFun();
}

// ------------------------------------------------------------------ listItems
void listItems(List<String> items) {
  print('Items List:');
  for (final item in items) {
    print('Item: $item');
  }
}

// ------------------------------------------------------------------ fibonacci
int fibonacci(int n) {
  if (n == 0 || n == 1) return n;
  return fibonacci(n - 1) + fibonacci(n - 2);
}

// -------------------------------------------------------------- flightDataFun
void flightDataFun() {
  print('');

  var tmp = FlightData.withDefaults("A new flight", DateTime.now());
  tmp.report();
}

// ------------------------------------------------------------------ constants
const DefaultReserve = 3;
const DefaultSeatingCapacity = 25;

// ----------------------------------------------------------- class FlightData
class FlightData {
  String description;
  DateTime? departureTime;
  int seatingCapacity;
  int reserveCapacity;

  int get passengerCapacity => seatingCapacity - reserveCapacity;

  // ------------------------------------------------------------- constructors
  FlightData(this.description, this.departureTime, this.seatingCapacity,
      this.reserveCapacity);

  FlightData.withDefaultReserve(description, departureTime, seatingCapacity)
      : this(description, departureTime, seatingCapacity, DefaultReserve);

  FlightData.withDefaults(description, departuretime)
      : this(
            description, departuretime, DefaultSeatingCapacity, DefaultReserve);

  // ------------------------------------------------------------------ methods
  void report() {
    print('Flight "$description" has room for $passengerCapacity passengers.');
    if (departureTime != null) {
      print('It leaves at $departureTime');
    }
  }
}
