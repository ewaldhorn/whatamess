import 'car.dart';

class ElectricCar extends Car {
  @override
  void brake() {
    super.brake();
  }

  @override
  void drive() {
    super.drive();
  }

  @override
  void rechargeBattery() {
    super.rechargeBattery();
  }

  @override
  void fillGas() {
    print('electric car doesn\'t use gas');
  }

  @override
  void fillPetrol() {
    print('electric car doesn\'t use petrol');
  }
}
