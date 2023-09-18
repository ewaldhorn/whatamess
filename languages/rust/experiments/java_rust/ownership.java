class Ownership {
    public static void main(String[] args) {
        var item = new Item(10);

        var item_reference = item;

        Ownership.doSomething(item);
        Ownership.doSomething(item_reference);
    }

    public static void doSomething(Item item) {
        item.count = item.count + 1;
        System.out.println("Doing something with " + item.count + " items.");
    }
}

class Item {
    int count;

    Item(int startWith) {
        this.count = startWith;
    }
}