// TODO: Declare the two variables here
String? name = 'Jane', address;

// Tests your solution (Don't edit!):
void main() {
  try {
    if (name == 'Jane' && address == null) {
      // verify that "name" is nullable
      name = null;
      print('Success!');
    } else {
      print('Not quite right, try again!');
    }
  } catch (e) {
    print('Exception: ${e.runtimeType}');
  }
}
