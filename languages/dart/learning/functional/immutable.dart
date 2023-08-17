void main() {
  var p1 = Person(name: 'Bob', age: 10);
  var p1Updated = p1.copyWith(age: 11);

  print('Person1        : $p1');
  print('Person1 Updated: $p1Updated');
}

class Person {
  final String name;
  final int age;

  const Person({required this.name, required this.age});

  @override
  String toString() {
    return '{name:$name, age:$age}';
  }

  Person copyWith({String? name, int? age}) {
    return Person(name: name ?? this.name, age: age ?? this.age);
  }
}
