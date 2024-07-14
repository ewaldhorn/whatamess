String joinWithCommas(int a, [int? b, int? c, int? d, int? e]) {
  String answer = '$a';

  answer += (b != null) ? ',$b' : '';
  answer += (c != null) ? ',$c' : '';
  answer += (d != null) ? ',$d' : '';
  answer += (e != null) ? ',$e' : '';

  return answer;
}

// Tests your solution (Don't edit!):
void main() {
  final errs = <String>[];

  try {
    final value = joinWithCommas(1);

    if (value != '1') {
      errs.add(
          'Tried calling joinWithCommas(1) \n and got $value instead of the expected (\'1\').');
    }
  } on UnimplementedError {
    print(
        'Tried to call joinWithCommas but failed. \n Did you implement the method?');
    return;
  } catch (e) {
    print(
        'Tried calling joinWithCommas(1), \n but encountered an exception: ${e.runtimeType}.');
    return;
  }

  try {
    final value = joinWithCommas(1, 2, 3);

    if (value != '1,2,3') {
      errs.add(
          'Tried calling joinWithCommas(1, 2, 3) \n and got $value instead of the expected (\'1,2,3\').');
    }
  } on UnimplementedError {
    print(
        'Tried to call joinWithCommas but failed. \n Did you implement the method?');
    return;
  } catch (e) {
    print(
        'Tried calling joinWithCommas(1, 2 ,3), \n but encountered an exception: ${e.runtimeType}.');
    return;
  }

  try {
    final value = joinWithCommas(1, 2, 3, 4, 5);

    if (value != '1,2,3,4,5') {
      errs.add(
          'Tried calling joinWithCommas(1, 2, 3, 4, 5) \n and got $value instead of the expected (\'1,2,3,4,5\').');
    }
  } on UnimplementedError {
    print(
        'Tried to call joinWithCommas but failed. \n Did you implement the method?');
    return;
  } catch (e) {
    print(
        'Tried calling stringify(1, 2, 3, 4 ,5), \n but encountered an exception: ${e.runtimeType}.');
    return;
  }

  if (errs.isEmpty) {
    print('Success!');
  } else {
    errs.forEach(print);
  }
}
