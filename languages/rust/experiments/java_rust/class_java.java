class Classes {
    public static void main(String[] args) {
        var contact = new Contact("Jan", 29);
        System.out.println("Hello " + contact.name + ", you are " + contact.age + " years old!");
    }
}

class Contact implements BusinessCard {
    String name;
    short age;

    Contact(String name, short age) {
        this.name = name;
        this.age = age;
    }

    public String getName() {
        return name;
    }

    public String card() {
        return "Business Card for " + name;
    }
}

interface BusinessCard {
    public String card();
}