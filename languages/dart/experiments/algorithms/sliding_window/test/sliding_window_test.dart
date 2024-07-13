import 'package:sliding_window/sliding_window.dart';
import 'package:test/test.dart';

void main() {
  group('findMaxSumOfSubArray', () {
    test('([1, 2, 6, 2, 4, 1], 3), 12)', () {
      expect(findMaxSumOfSubArray([1, 2, 6, 2, 4, 1], 3), 12);
    });

    test('([1, 2, 6, 2, 4, 1], 3), 12)', () {
      expect(findMaxSumOfSubArray([2, 2, 2, 2, 2, 2], 3), 6);
    });
  });
}
