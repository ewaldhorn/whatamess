Future<void> fetchUserOrder() {
  // Imagine that this function is fetching user info from another service or database.
  return Future.delayed(const Duration(seconds: 3), () => print('Large Latte'));
}

void main() {
  fetchUserOrder();
  print('Fetching user order...');
}

// dart run simple.dart results in:
// Fetching user order...
// <a 3 second delay>
// Large Latte
