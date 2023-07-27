void main() {
  const String sourceString = "thiS is the SOurCE stRing";
  var lowerCased = sourceString.toLowerCase();
  var upperCased = sourceString.toUpperCase();
  var prettified = capitalWords(sourceString);

  print("");
  print("Original string: $sourceString");
  print("Lowercase      : $lowerCased");
  print("Uppercase      : $upperCased");
  print("Prettified     : $prettified");
  print("");
}

String capitalWords(String input) {
  if (input.isEmpty) return input;

  return input
      .toLowerCase()
      .split(' ')
      .map((word) =>
          word.isEmpty ? word : '${word[0].toUpperCase()}${word.substring(1)}')
      .join(' ');
}
