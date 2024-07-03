import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

const faceValue = ['2', '3', '4', '5', '6', '7', '8', '9'];
const tenCards = ['10', 'J', 'Q', 'K'];

void main() {
  int bank = scoreCards(readLineSync().split(' '));
  int player = scoreCards(readLineSync().split(' '));

  if (player > 21 && bank > 21) {
    print('Bank');
  } else if (player == bank) {
    print('Draw');
  } else if (bank == 0) {
    print('Bank');
  } else if (player == 0) {
    print('Blackjack!');
  } else if (player > 21) {
    print('Bank');
  } else if (bank > 21) {
    print('Player');
  } else if (player > bank) {
    print('Player');
  } else if (bank > player) {
    print('Bank');
  }
}

int scoreCards(List<String> cards) {
  int total = 0;
  int aces = 0;

  for (String card in cards) {
    if (faceValue.contains(card)) {
      total += int.parse(card);
    } else if (tenCards.contains(card)) {
      total += 10;
    } else if (card == 'A') {
      aces += 1;
    }
  }
  // check if we have blackjack, then return that instead
  if (cards.length == 2 && aces == 1 && total == 21) {
    return 0;
  }

  while (aces > 0) {
    if (total < 20) {
      total += 11;
    } else {
      total += 1;
    }
    aces -= 1;
  }

  return total;
}
