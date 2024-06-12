bool isAnagram(String left, String right) {
  // short-circuit on empty inputs
  if (left.isEmpty && right.isEmpty) {
    return true;
  }

  // different length strings aren't worth testing
  if (left.length != right.length) {
    return false;
  }

  // split into characters
  List<String> leftArray = left.split("");
  List<String> rightArray = right.split("");

  // sort them
  leftArray.sort((l, r) => l.compareTo(r));
  rightArray.sort((l, r) => l.compareTo(r));

  // if the two arrays are not identical, it's not an anagram
  for (int pos = 0; pos < leftArray.length; pos++) {
    if (leftArray[pos] != rightArray[pos]) {
      return false;
    }
  }

  return true;
}
