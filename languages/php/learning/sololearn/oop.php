<?php
    class Person {
        public $age; //property
        public function speak() { //method
          echo "Hi!"
        }
      }
?>

<?php
class Dog {
    public $legs=4;
    public function display() {
        echo $this->legs;
    }
}
$d1 = new Dog();
$d1->display();

echo '<br />';

$d2 = new Dog();
$d2->legs = 2;
$d2->display();
?>