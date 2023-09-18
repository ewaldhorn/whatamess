class Generics {
    public static void main(String[] args) {
        var containerInt = new Container<Integer>(456);
        var containerString = new Container<String>("Yolo");

        System.out.println("Compare with other: " + containerInt.compareItem(11));
        System.out.println("And more: " + Generics.compare("aab", "bba"));
    }

    public static <T extends Comparable> int compare(T item1, T item2) {
        return item1.compareTo(item2);
    }
}

class Container<T extends Comparable> {
    T item;

    Container(T t) {
        this.item = t;
    }

    int compareItem(T other) {
        return item.compareTo(other);
    }
}