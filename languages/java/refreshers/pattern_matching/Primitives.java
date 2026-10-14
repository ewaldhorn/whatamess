// NOTE
// You must run it with preview enabled
// java --enable-preview Primitives.java
void main() {
    // Java 23/25 Primitive patterns (JEP 455): zero autoboxing overhead!
    int statusCode = 200;
    String status = switch (statusCode) {
        case 200 -> "OK";
        case 404 -> "Not Found";
        case int code when code >= 500 -> "Server Error (" + code + ")";
        case int code -> "Other Status: " + code;
    };

    IO.println("Result is " + status);
}
