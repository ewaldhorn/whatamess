int score(String s) {
  int total = 0;

  for (final c in s.toUpperCase().split('')) {
    total += getScoreFor(c);
  }

  return total;
}

int getScoreFor(String c) {
  int score = 0;
  const values = {
    1: 'AEIOULNRST',
    2: 'DG',
    3: 'BCMP',
    4: 'FHVWY',
    5: 'K',
    8: 'JX',
    10: 'QZ'
  };

  for (var entry in values.entries) {
    if (entry.value.contains(c)) {
      score = entry.key;
      break;
    }
  }

  return score;
}
