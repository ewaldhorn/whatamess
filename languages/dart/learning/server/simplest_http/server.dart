import 'dart:io';

void main() {
 var server = HttpServer.bind(InternetAddress.loopbackIPv4, 8080);
 print('Server running on localhost:${server.port}');
 
 server.listen((HttpRequest request) {
   request.response.write('Hello, Dart!');
   request.response.close();
 });
}
