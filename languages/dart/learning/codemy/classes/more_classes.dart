void main() {
    Person p1 = Person("Frank", "M", 38);
    p1.showData();

    Person p2 = Person("Lola", "F", 40);
    p2.showData();
}

class Person {
    String? name, sex;
    int age = 0;

    Person(String name, sex, int age) {
        this.name = name;
        this.sex = sex;
        this.age = age;
    }

    void showData() {
        print("$name is $age years old and $sex.");
    }
}