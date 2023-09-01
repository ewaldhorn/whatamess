import 'petrol_car.dart';

void main(List<String> args) {
  final car1 = PetrolCar();

  car1.fillGas();
  car1.brake();
  car1.rechargeBattery();
}
