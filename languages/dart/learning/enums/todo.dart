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

// ======================================================================= MAIN
void main() {
  Day today = Day.friday;
  print('So today is ${today.name} ("${today.abbreviation}").');
}
