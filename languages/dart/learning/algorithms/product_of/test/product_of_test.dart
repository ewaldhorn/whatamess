import 'package:product_of/product_of.dart';
import 'package:test/test.dart';

void main() {
  group("basic nested loops", () {
    test("1,2,3,4 => 24,12,8,6", () {
      var input = [1, 2, 3, 4];
      var answer = [24, 12, 8, 6];

      expect(productOfArray(input), answer);
    });

    test("4,3,2,1 => 6,8,12,24", () {
      var input = [4, 3, 2, 1];
      var answer = [6, 8, 12, 24];

      expect(productOfArray(input), answer);
    });

    test("10,3,5,6,2 => 180,600,360,300,900", () {
      var input = [10, 3, 5, 6, 2];
      var answer = [180, 600, 360, 300, 900];

      expect(productOfArray(input), answer);
    });
  });

  group("without nested loops", () {
    test("1,2,3,4 => 24,12,8,6", () {
      var input = [1, 2, 3, 4];
      var answer = [24, 12, 8, 6];

      expect(productOfArrayNoNesting(input), answer);
    });

    test("4,3,2,1 => 6,8,12,24", () {
      var input = [4, 3, 2, 1];
      var answer = [6, 8, 12, 24];

      expect(productOfArrayNoNesting(input), answer);
    });

    test("10,3,5,6,2 => 180,600,360,300,900", () {
      var input = [10, 3, 5, 6, 2];
      var answer = [180, 600, 360, 300, 900];

      expect(productOfArrayNoNesting(input), answer);
    });
  });
}
