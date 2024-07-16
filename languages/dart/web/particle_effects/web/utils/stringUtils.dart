// --------------------------------------------------------------- lineSplitter
// splits a long string into a list of strings of the specified length
// first break up the string into words
// now start putting them back together
List<String> lineSplitter(String msg, {int maxLen = 20}) {
  List<String> words = msg.split(' ');
  List<String> strings = [];

  StringBuffer sb = StringBuffer();

  words.forEach((w) {
    if (sb.length < maxLen) {
      sb.write('$w ');
    } else {
      strings.add(sb.toString().trim());
      sb.clear();
      sb.write('$w ');
    }
  });

  if (sb.isNotEmpty) {
    strings.add(sb.toString().trim());
  }

  sb.clear();

  return strings;
}
