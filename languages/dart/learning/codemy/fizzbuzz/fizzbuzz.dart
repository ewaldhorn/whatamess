import 'dart:io';

void main() {
    int num = 1;
    const int max = 35;

    while (num <= max) {
        if (num % 5 == 0 && num % 3 == 0) {
            print("FizzBuzz");
        } else if (num % 5 == 0) {
            print("Fizz");
        } else if (num % 3 == 0) {
            print("Buzz");
        } else {
            print(num);
        }

        num++;
    }
}