// Imports all public packages exported by java.base
import module java.base;

void main() {
    List<String> names = List.of("Alice", "Bob", "Charlie");
    Map<Integer, List<String>> byLength = names.stream()
        .collect(Collectors.groupingBy(String::length));
    IO.println(byLength);
}
