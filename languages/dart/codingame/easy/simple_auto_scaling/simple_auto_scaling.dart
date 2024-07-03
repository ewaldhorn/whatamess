import 'dart:io';

String readLineSync() {
  String? s = stdin.readLineSync();
  return s == null ? '' : s;
}

void main() {
    List inputs;
    inputs = readLineSync().split(' ');
    int S = int.parse(inputs[0]); // number of services
    int M = int.parse(inputs[1]); // number of lines to parse after this next one

    var services = new List<int>.filled(S, 0);
    var current = new List<int>.filled(S, 0);

    inputs = readLineSync().split(' '); // max capacity per service

    for (int i = 0; i < S; i++) {
        services[i] = int.parse(inputs[i]);
    }

    stderr.write("${services}\n");

    for (int i = 0; i < M; i++) {
        inputs = readLineSync().split(' ');
        for (int j = 0; j < S; j++) {
            int clients = int.parse(inputs[j]);
            var requiredServices = (clients / services[j].toDouble()).ceil();

            if (current[j] > requiredServices) {
              stdout.write('${requiredServices-current[j]}');
              current[j] -= current[j]-requiredServices;
            } else if (current[j] < requiredServices) {
              stdout.write('${requiredServices-current[j]}');
              current[j] += requiredServices - current[j];
            } else {
              stdout.write('0');
            }

            if (j < S-1) {
              stdout.write(' ');
            }
        }
        print('');
    }
}