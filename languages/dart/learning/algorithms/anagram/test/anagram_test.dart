import 'package:anagram/anagram.dart';
import 'package:test/test.dart';

void main() {
  group("anagrams", () {
    test("one,two : false", () {
      expect(isAnagram("one", "two"), false);
    });

    test("two,wat : false", () {
      expect(isAnagram("two", "wat"), false);
    });

    test("two,wot : true", () {
      expect(isAnagram("two", "wot"), true);
    });

    test("loop, pool : true", () {
      expect(isAnagram("loop", "pool"), true);
    });
  });
}
