package formatting

import java.text.SimpleDateFormat
import java.util.*

/**
 * Basic text formatting in Kotlin examples.
 */
fun main() {
    val integer = 10
    val double = 10.345
    val smallDouble = 0.0035

    // Java's String.format
    val rightPaddedInteger = String.format("%-12d", integer)
    val paddedInteger = String.format("%12d", integer)
    val rightPaddedDouble = String.format("%-12f", double)
    val paddedDouble = String.format("%12f", double)

    println("Standard integer is        : |$integer|")
    println("The right padded integer is: |$rightPaddedInteger|")
    println("The padded integer is      : |$paddedInteger|")

    println("Standard double is         : |$double|")
    println("Padded double is           : |$paddedDouble|")
    println("Right Padded double is     : |$rightPaddedDouble|")

    println("Small Double is (default)  : |$smallDouble|")
    println("Small Double is  (5 dec)   : |${"%.5f".format(smallDouble)}|")
    println("Small Double is  (4 dec)   : |${"%.4f".format(smallDouble)}|")
    println("Small Double is  (3 dec)   : |${"%.3f".format(smallDouble)}|")
    println("Small Double is  (2 dec)   : |${"%.2f".format(smallDouble)}|")

    println("Double (default)           : |$double|")
    println("Double (5 dec)             : |${"%.5f".format(double)}|")
    println("Double (4 dec)             : |${"%.4f".format(double)}|")
    println("Double (3 dec)             : |${"%.3f".format(double)}|")
    println("Double (2 dec)             : |${"%.2f".format(double)}|")
    println("Double (1 dec)             : |${"%.1f".format(double)}|")
    println("Double (0 dec)             : |${"%.0f".format(double)}|")

    // Declare variables of different types
    val boolValue = true
    val charValue = 'A'
    val intValue = -42
    val floatValue = 3.14159f
    val stringValue = "Hello, World!"
    val dateValue = Date()
    val hexValue = 255

    // Print values using different format specifiers
    println("Boolean: %b".format(boolValue))  // Format a boolean value
    println("Character: %c".format(charValue))  // Format a character value
    println("Signed Integer: %d".format(intValue))  // Format a signed integer value
    println("Float in Scientific Notation: %e".format(floatValue))  // Format a float value in scientific notation
    println("Float in Decimal Format: %f".format(floatValue))  // Format a float value in decimal format
    println("Float in Decimal or Scientific Notation: %g".format(floatValue))  // Format a float value in decimal or scientific notation
    println("Hashcode: %h".format(stringValue))  // Format the hashcode of a string value
    println("Newline Separator: %n")  // Print a newline separator
    println("Octal Integer: %o".format(intValue))  // Format an integer value in octal format
    println("String: %s".format(stringValue))  // Format a string value
    println("Date: ${SimpleDateFormat("yyyy-MM-dd").format(dateValue)}")  // Format the date value using SimpleDateFormat
    println("Hexadecimal Integer: %x".format(hexValue))  // Format an integer value in hexadecimal format
}