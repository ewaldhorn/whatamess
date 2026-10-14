record Point(double x, double y) {}
record Circle(Point center, double radius) implements Shape {}
record Rectangle(Point topLeft, double width, double height) implements Shape {}

sealed interface Shape permits Circle, Rectangle {}

double calculateArea(Shape shape) {
    return switch (shape) {
        // Discard unneeded center point with '_' (JEP 456)
        case Circle(_, var radius) -> Math.PI * radius * radius;

        // Nested deconstruction: extract coordinates directly
        case Rectangle(Point(var x, var y), var w, var h) -> {
            IO.println("Top-left: (" + x + ", " + y + ")");
            yield w * h;
        }
    }; // NO 'default:' needed! Compiler proves Circle and Rectangle cover all Shape permits.
}

void main() {
    IO.println();
    IO.println("====== First we do the circle");
    var circle = new Circle(new Point(10.5, 10.3), 88.3);
    IO.println("The circle has an area of " + "%.3f".formatted(calculateArea(circle)));
    System.out.printf("This still works        : %.3f\n", calculateArea(circle));
    IO.println();
    IO.println("====== Now the rectangle");
    var rectangle = new Rectangle(new Point(4.4, 5.4), 11.3, 19.1);
    IO.println("The rectangle area is : " + calculateArea(rectangle));
    IO.println();
}
