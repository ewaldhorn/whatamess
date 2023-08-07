void main() {
  simpleList();
}

void simpleList() {
  final names = [
    'Rob',
    "Karen",
    'Nick',
    "Elma",
    "Richard",
    'Anne',
    'Arnold',
    'Leo',
    "William"
  ];

  print(
      'There are ${names.length} names in the list. Currently, it is unsorted.');
  print(
      'The first name is "${names.first}" and the last name is "${names.last}".');
  print('This list has a runtime type of ${names.runtimeType}.');
  print('This list is empty: ${names.isEmpty}.');

  names.add('Nelly');
  print('Added a name to the list, for a total of ${names.length} entries.');
  print(
      'The first name is still "${names.first}" and the last name in the list is now "${names.last}".');
}
