import 'package:anagram/anagram.dart';
import 'package:test/test.dart';

void main() {
  group("anagrams", () {
    test("empty strings : true", () {
      expect(isAnagram("", ""), true);
    });

    test("one,two : false", () {
      expect(isAnagram("one", "two"), false);
    });

    test("two,wat : false", () {
      expect(isAnagram("two", "wat"), false);
    });

    test("two,wot : true", () {
      expect(isAnagram("two", "wot"), true);
    });

    test("potato,peel : false", () {
      expect(isAnagram("potato", "peel"), false);
    });

    test("loop, pool : true", () {
      expect(isAnagram("loop", "pool"), true);
    });

    test("solo, solos : false", () {
      expect(isAnagram("solo", "solos"), false);
    });
    test("racecar, fastcar : false", () {
      expect(isAnagram("racecar", "fastcar"), false);
    });
    test("petrol, patrol : false", () {
      expect(isAnagram("petrol", "patrol"), false);
    });
    test("spud,puds : true", () {
      expect(isAnagram("spud", "puds"), true);
    });
  });
}
