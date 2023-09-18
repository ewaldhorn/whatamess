import java.util.Optional;

class Optionals {
    public static void main(String[] args) {
        WorkItem item1 = null;
        System.out.println("item1: " + item1);

        Optional<WorkItem> item2 = Optional.of(new WorkItem());
        System.out.println("item2: " + item2);

        Optional<WorkItem> item3 = Optional.empty();
        System.out.println("item3: " + item3);

    }
}

class WorkItem {

}