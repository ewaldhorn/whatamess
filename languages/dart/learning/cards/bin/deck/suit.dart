/*
Marking a class as sealed has 2 main effects:
 - The class is abstract (you cannot create a concrete instance of Suit)
 - All subtypes must be defined in the same library
*/
sealed class Suit {}

class Club extends Suit {}

class Diamond extends Suit {}

class Heart extends Suit {}

class Spade extends Suit {}
