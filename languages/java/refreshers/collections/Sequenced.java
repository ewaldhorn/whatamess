void main() {
  SequencedCollection<String> items = new ArrayList<>(List.of("A", "B", "C", "D"));

  String first = items.getFirst();       // "A"
  String last  = items.getLast();        // "D"

  items.addFirst("Z");                  // ["Z", "A", "B", "C", "D"]
  items.removeLast();                   // Removes "D"

  // Constant-time reversed view (not a copy!)
  for (String s : items.reversed()) {
    IO.print(s + " ");                    // C B A Z
  }

  IO.println();
}
