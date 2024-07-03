import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  String b = readLineSync();
  var max = 0;
  var chars = b.split('');

  for (int i = 0; i < chars.length; i++) {
    if (chars[i] == '0') {
      var tmp = b.split('');
      tmp[i] = '1';
      var segments = tmp.join('').split('0');
      segments.forEach((seg) {
        if (seg.length > max) {
          max = seg.length;
        }
        ;
      });
    }
  }

  print(max);
}
