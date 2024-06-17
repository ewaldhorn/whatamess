import 'dart:io';
import 'dart:math';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
    List inputs;
    inputs = readLineSync().split(' ');
    int S = int.parse(inputs[0]); // number of services
    int M = int.parse(inputs[1]); // number of lines to parse after this next one

    var services = new List<int>.filled(5, 0);
    inputs = readLineSync().split(' '); // max capacity per service

    for (int i = 0; i < S; i++) {
        services[i] = int.parse(inputs[i]);
    }
    
    for (int i = 0; i < M; i++) {
        inputs = readLineSync().split(' ');
        for (int j = 0; j < S; j++) {
            
            int clients = int.parse(inputs[j]);
        }
    }
    for (int i = 0; i < M; i++) {

        // Write an answer using print()
        // To debug: stderr.writeln('Debug messages...');

        print('Number of services to start / stop');
    }
}
