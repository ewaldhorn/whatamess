// FreeCodeCamp Article
// https://www.freecodecamp.org/news/how-to-use-enhanced-enums-in-dart/
// ====================================================================== ENUMS
enum Priority { low, medium, high }

enum Day {
  monday("Mon"),
  tuesday("Tu"),
  wednesday("Wed"),
  thursday("Thu"),
  friday("Fri"),
  saturday("Sat"),
  sunday("Sun");

  const Day(this.abbreviation);

  final String abbreviation;
}

enum Month {
  January("Jan"),
  February("Feb"),
  March("Mar"),
  April("Apr"),
  May("May"),
  June("Jun"),
  July("Jul"),
  August("Aug"),
  September("Sep"),
  October("Oct"),
  November("Nov"),
  December("Dec");

  final String abbreviation;

  const Month(this.abbreviation);

  Month operator +(int other) {
    // Ensure the result stays within 0-11 range
    int result = (this.index + other) % 12;
    return Month.values[result];
  }
}

// ======================================================================= MAIN
void main() {
  Day today = Day.friday;
  print('So today is ${today.name} ("${today.abbreviation}").');
  print('');

  Month thisMonth = Month.July;
  print('Current month is ${thisMonth.name}.');

  Month nextMonth = thisMonth + 1;
  print('Next month is ${nextMonth.name}.');

  print('And then the next month is ${(nextMonth + 1).name}.');
}
