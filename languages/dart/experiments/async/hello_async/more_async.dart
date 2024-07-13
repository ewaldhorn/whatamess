import 'dart:io';

Future<void> printOrderMessage() async {
  stdout.write('Awaiting user order... ');
  var order = await fetchUserOrder();
  print('\nYour order is: $order');
}

Future<String> fetchUserOrder() {
  // Imagine that this function is more complex and slow.
  return Future.delayed(const Duration(seconds: 4), () => 'Large Latte');
}

void main() async {
  countSeconds(4);
  await printOrderMessage();
}

// You can ignore this function - it's here to visualize delay time in this example.
void countSeconds(int s) {
  for (var i = s; i > 0; i--) {
    Future.delayed(Duration(seconds: 1 + s - i), () => stdout.write('$i '));
  }
}
