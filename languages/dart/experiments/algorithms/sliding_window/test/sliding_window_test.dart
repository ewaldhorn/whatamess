import 'package:sliding_window/sliding_window.dart';
import 'package:test/test.dart';

void main() {
  group('findMaxSumOfSubArrayLooped', () {
    test('([1, 2, 6, 2, 4, 1], 3), 12)', () {
      expect(findMaxSumOfSubArrayLooped([1, 2, 6, 2, 4, 1], 3), 12);
    });

    test('([1, 2, 6, 2, 4, 1], 1), 6)', () {
      expect(findMaxSumOfSubArrayLooped([1, 2, 6, 2, 4, 1], 1), 6);
    });

    test('([2, 2, 2, 2, 2, 2], 3), 6)', () {
      expect(findMaxSumOfSubArrayLooped([2, 2, 2, 2, 2, 2], 3), 6);
    });

    test('([2, 2, 2, 2, 2, 2], 4), 8)', () {
      expect(findMaxSumOfSubArrayLooped([2, 2, 2, 2, 2, 2], 4), 8);
    });
  });

  group('findMaxSumOfSubArray', () {
    test('([1, 2, 6, 2, 4, 1], 3), 12)', () {
      expect(findMaxSumOfSubArray([1, 2, 6, 2, 4, 1], 3), 12);
    });

    test('([1, 2, 6, 2, 4, 1], 1), 6)', () {
      expect(findMaxSumOfSubArray([1, 2, 6, 2, 4, 1], 1), 6);
    });

    test('([2, 2, 2, 2, 2, 2], 3), 6)', () {
      expect(findMaxSumOfSubArray([2, 2, 2, 2, 2, 2], 3), 6);
    });

    test('([2, 2, 2, 2, 2, 2], 4), 8)', () {
      expect(findMaxSumOfSubArray([2, 2, 2, 2, 2, 2], 4), 8);
    });
  });
}
