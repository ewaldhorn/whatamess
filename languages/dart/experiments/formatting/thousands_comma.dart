void main(List<String> args) {
  print("999     becomes: ${intComma(999)}");
  print("1000    becomes: ${intComma(1000)}");
  print("1001    becomes: ${intComma(1001)}");
  print("1010    becomes: ${intComma(1010)}");
  print("1234    becomes: ${intComma(1234)}");
  print("10000   becomes: ${intComma(10000)}");
  print("10001   becomes: ${intComma(10001)}");
  print("10010   becomes: ${intComma(10010)}");
  print("10100   becomes: ${intComma(10100)}");
  print("12345   becomes: ${intComma(12345)}");
  print("100000  becomes: ${intComma(100000)}");
  print("123456  becomes: ${intComma(123456)}");
  print("1000000 becomes: ${intComma(1000000)}");
  print("1234567 becomes: ${intComma(1234567)}");
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
