void main() {
    print("\nDart Variables:\n===============\n");

    // Strings
    var thisPerson = "Bob";
    String thisIsAString = "This is a string!";

    print("The variable thisPerson contains '$thisPerson' and thisIsAString == '$thisIsAString'.");

    // Numbers
    var myNumber = 14;
    int anyInteger = 7;

    print("$myNumber / $anyInteger = ${myNumber/anyInteger}");

    // Dynamic variables
    dynamic thisIsADynamic = 50;

    print("The dynamic contains $thisIsADynamic");
    thisIsADynamic = "Yolo";
    print("The dynamic contains $thisIsADynamic");

    // Const and Final
    // Const for compile-time constants, known at compile time
    // Final for runtime constants, unknown at compile time
    const String aintChanging = "This is it!";
    // aintChanging = "Nope"; // can't do this. Nope.

    final String changeable = "Bobby";
    print("changeable is $changeable");
    // changeable = "Billy"; // can't do this either...

    print("");
}