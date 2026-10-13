String describe(Object obj) {
    return switch (obj) {
        case null                  -> "Null reference";
        case Integer i when i < 0  -> "Negative integer: " + i;
        case Integer i             -> "Non-negative integer: " + i;
        case String s when s.isEmpty() -> "Empty string";
        case String s              -> "String of length " + s.length();
        case List<?> list          -> "List containing " + list.size() + " elements";
        default                    -> "Unknown object type: " + obj.getClass().getSimpleName();
    };
}

void main() {
  IO.println("Null     : " + describe(null));
  IO.println("-50      : " + describe(-50));
  IO.println("100      : " + describe(100));
  IO.println("\"\"       : " + describe(""));
  IO.println("\"123\"    : " + describe("123"));
  
var myList = List.of("a", "b", "c");
  IO.println("[a,b,c]  : " + describe(myList));
  IO.println("123.45   : " + describe(123.45));
}
