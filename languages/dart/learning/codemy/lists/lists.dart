// Growable (dynamic) lists
void main() {
    // Lists
    var myList = [0,1,2];
    print("\n\tOur list is currently $myList");

    // Change item
    print("Modifying list...");
    myList[0] = 1;
    myList[1] = 2;
    myList[2] = 3;
    print("\tOur list is currently $myList");

    // Create empty list
    print("Creating a new list...");
    var newList = [];
    print("\tOur list is currently $newList");

    // Add to list
    print("Adding to the list...");
    newList.add(0);
    print("\tOur list is currently $newList");
    print("Adding some more...");
    newList.add(1);
    newList.add(2);
    print("\tOur list is currently $newList");

    // Add multiple items to list
    print("Adding multiple items to the list...");
    newList.addAll([3,4,6,7,8,9]);
    print("\tOur list is currently $newList");

    // Insert at a specific position in the list
    print("Adding the missing 5...");
    newList.insert(5, 5);
    print("\tOur list is currently $newList");

    // Insert many items
    print("Reset list a bit...");
    newList = [0,1,2,7,8,9];
    print("\tOur list is currently $newList");
    print("Adding the missing items...");
    newList.insertAll(3, [3,4,5,6]);
    print("\tOur list is currently $newList");

    // Mixed lists
    var anyList = [1,2,"Some","Thing","Strange",false];
    print("\tOur mixed list is currently $anyList");

    // Remove from list
    print("Remove 'Strange'");
    anyList.remove('Strange');
    print("\tOur mixed list is currently $anyList");

    // Remove from a specific location
    print("Remove the 3rd item...");
    anyList.removeAt(3);
    print("\tOur mixed list is currently $anyList");

    print("");
}