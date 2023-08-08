import 'dart:io';

import 'package:dart_frog/dart_frog.dart';

Future<HttpServer> run(Handler handler, InternetAddress ip, int port) {
  const usePort = 9000;
  if (port == 8080) {
    // ignore: avoid_print
    print("Can't use 8080, going to use $usePort instead.");
  }
  return serve(handler, ip, usePort, poweredByHeader: null);
}
