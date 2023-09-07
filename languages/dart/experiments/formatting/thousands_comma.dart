void main(List<String> args) {
  print("9         becomes: ${intComma(9)}");
  print("9         becomes: ${intCommaII(9)}");
  print("99        becomes: ${intCommaII(99)}");
  print("999       becomes: ${intComma(999)}");
  print("999       becomes: ${intCommaII(999)}");
  print("1000      becomes: ${intComma(1000)}");
  print("1000      becomes: ${intCommaII(1000)}");
  print("1001      becomes: ${intComma(1001)}");
  print("1001      becomes: ${intCommaII(1001)}");
  print("1010      becomes: ${intComma(1010)}");
  print("1234      becomes: ${intComma(1234)}");
  print("10000     becomes: ${intComma(10000)}");
  print("10000     becomes: ${intCommaII(10000)}");
  print("10001     becomes: ${intComma(10001)}");
  print("10010     becomes: ${intComma(10010)}");
  print("10100     becomes: ${intComma(10100)}");
  print("12345     becomes: ${intComma(12345)}");
  print("12345     becomes: ${intCommaII(12345)}");
  print("100000    becomes: ${intComma(100000)}");
  print("100000    becomes: ${intCommaII(100000)}");
  print("123456    becomes: ${intComma(123456)}");
  print("123456    becomes: ${intCommaII(123456)}");
  print("1000000   becomes: ${intComma(1000000)}");
  print("1234567   becomes: ${intComma(1234567)}");
  print("1234567   becomes: ${intCommaII(1234567)}");
  print("12345678  becomes: ${intComma(12345678)}");
  print("12345678  becomes: ${intCommaII(12345678)}");
  print("123456789 becomes: ${intComma(123456789)}");
  print("123456789 becomes: ${intCommaII(123456789)}");
}

String intComma(int i) {
  if (i < 0) {
    return "-" + intComma(-i);
  }

  if (i < 1000) {
    return "$i";
  }

  return "${intComma(i ~/ 1000)},${padWithZero(i % 1000)}";
}

String padWithZero(int number) {
  if (number < 10) {
    return "00$number";
  } else if (number < 100) {
    return "0$number";
  } else
    return "$number";
}

String intCommaII(int i) {
  if (i < 0) {
    return "-" + intComma(-i);
  }

  if (i < 1000) {
    return "$i";
  }

  return "${intCommaII(i ~/ 1000)},${padToThree(i % 1000)}";
}

String padToThree(int number) {
  int zeroes = 3 - '${number}'.length;
  if (zeroes > 0) {
    return '${"0" * zeroes}${number}';
  } else {
    return '$number';
  }
}
