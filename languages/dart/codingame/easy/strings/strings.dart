import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
  String BEGINS = readLineSync();
  String ENDS = readLineSync();

  List<String> BEGIN = BEGINS.split('.');
  List<String> END = ENDS.split('.');

  int diffYears = int.parse(END[2]) - int.parse(BEGIN[2]);
  int diffMonths = int.parse(END[1]) - int.parse(BEGIN[1]);
  int diffDays = int.parse(END[0]) - int.parse(BEGIN[0]);

  stderr.writeln('Y: $diffYears, M: $diffMonths, D: $diffDays');

  if (diffMonths < 0) {
    diffYears -= 1;
    diffMonths = 12 + diffMonths;
  }

  if (diffDays < 0) {
    diffMonths -= 1;
    diffDays = 31 + diffDays;
  }

  if (diffYears > 0) {
    if (diffYears == 1) {
      stdout.write('1 year, ');
    } else {
      stdout.write('$diffYears years, ');
    }
  }

  if (diffMonths > 0) {
    if (diffMonths == 1) {
      stdout.write('1 month, ');
    } else {
      stdout.write('$diffMonths months, ');
    }
  }

  int days = parseAsDateTime(ENDS).difference(parseAsDateTime(BEGINS)).inDays;

  print('total $days days');
}

DateTime parseAsDateTime(String inp) {
  // dd.mm.yyyy
  List<String> values = inp.split('.');
  return DateTime.parse('${values[2]}${values[1]}${values[0]}');
}
