void main() {
    Person p1 = Person("Briana", "F", 28);
    p1.showData();

    Person p2 = Person("Karl", "M", 58);
    p2.showData();
}

class Person {
    // properties
    String? name, sex;
    int? age;

    // constructor
    Person(String name, sex, int age) {
        this.name = name;
        this.sex = sex;
        this.age = age;
    }

    // methods
    void showData() {
        print("$name is $age years old, and $sex.");
    }
}