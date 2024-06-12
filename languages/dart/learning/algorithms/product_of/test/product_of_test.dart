import 'package:product_of/product_of.dart';
import 'package:test/test.dart';

void main() {
  test("1,2,3,4 => 24,12,8,6", () {
    var input = [1, 2, 3, 4];
    var answer = [24, 12, 8, 6];

    expect(productOfArray(input), answer);
  });
}
