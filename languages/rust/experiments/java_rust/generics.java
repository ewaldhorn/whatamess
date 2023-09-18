class Generics {
    public static void main(String[] args) {
        var containerInt = new Container<Integer>(456);
        var containerString = new Container<String>("Yolo");
    }
}

class Container<T> {
    T item;

    Container(T t) {
        this.item = t;
    }
}