import 'package:shelf/shelf.dart';

Response homePageHandler(Request req) {
  return Response.ok('Hello, World!\n');
}
