fun main() {
    println("Basic collections in Kotlin.")

    println("\n Lists:")
    // Read only list
    val readOnlyShapes = listOf("triangle", "square", "circle")
    println(readOnlyShapes) // [triangle, square, circle]

    // Mutable list with explicit type declaration
    val shapes: MutableList<String> = mutableListOf("triangle", "square", "circle")
    println(shapes) // [triangle, square, circle]

    // Create a read-only view of a mutable list
    val shapesLocked: List<String> = shapes
    shapes.add("rectangle")
    println(shapes)
    // can't modify shapesLocked, but the underlying list can still be modified
    println(shapesLocked)
}