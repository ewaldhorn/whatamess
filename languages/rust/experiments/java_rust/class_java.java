class Classes {
    public static void main(String[] args) {
        var contact = new Contact("Jan", 29);
        System.out.println("Hello " + contact.name + ", you are " + contact.age + " years old!");
    }
}

class Contact {
    String name;
    short age;

    Contact(String name, short age) {
        this.name = name;
        this.age = age;
    }
}