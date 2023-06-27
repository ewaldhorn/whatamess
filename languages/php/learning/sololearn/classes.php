<?php
class Person {
    public $name;
    public $age;
    public function __construct($name, $age) {
        $this->name = $name;
        $this->age = $age;
    }
}
$p = new Person("David", 42);
echo $p->name;
?>

<?php
class Person {
    public function __destruct() {
        echo "Object destroyed";
    }
}
$p = new Person();
?>