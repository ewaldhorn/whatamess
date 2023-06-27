<?php
class myClass {
    static $myProperty = 42;
    static function myMethod() {
        echo self::$myProperty;
    }
}

myClass::myMethod();
?>

<?php
class myClass {
    final function myFunction() {
        echo "Parent";
    }
}
// ERROR because a final method cannot be overridden in child classes.
class myClass2 extends myClass {
    function myFunction() {
        echo "Child";
    }
}
?>