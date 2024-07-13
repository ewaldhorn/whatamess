import 'dart:io';

Future<void> printOrderMessage() async {
  try {
    stdout.write('Awaiting user order... ');
    var order = await fetchUserOrder();
    print(order);
  } catch (err) {
    print('Caught error: $err');
  }
}

Future<String> fetchUserOrder() {
  // Imagine that this function is more complex.
  var str = Future.delayed(
      const Duration(seconds: 1), () => throw 'Cannot locate user order');
  return str;
}

void main() async {
  await printOrderMessage();
}
