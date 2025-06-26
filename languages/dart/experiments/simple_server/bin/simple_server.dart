import 'package:shelf/shelf.dart';
import 'package:shelf/shelf_io.dart' as shelf_io;
import 'dart:io';
import 'dart:convert';

// A simple handler for the root path
Response _rootHandler(Request request) {
  return Response.ok('Hello, world!\n');
}

// A handler for a specific API endpoint
Future<Response> _apiHandler(Request request) async {
  if (request.method == 'GET') {
    // Simulate fetching data
    final data = {
      'message': 'Data from Dart API',
      'timestamp': DateTime.now().toIso8601String(),
    };
    return Response.ok(
      jsonEncode(data),
      headers: {'Content-Type': 'application/json'},
    );
  } else if (request.method == 'POST') {
    final contentType = request.headers['content-type'];
    if (contentType != null && contentType.contains('application/json')) {
      try {
        final body = await request.readAsString();
        final json = jsonDecode(body);
        return Response.ok(
          jsonEncode({'received': json, 'status': 'success'}),
          headers: {'Content-Type': 'application/json'},
        );
      } catch (e) {
        return Response.badRequest(
          body: jsonEncode({'error': 'Invalid JSON in request body'}),
        );
      }
    }
    return Response.badRequest(
      body: 'Expected Content-Type: application/json for POST requests.',
    );
  }
  return Response.notFound('Not Found');
}

void main() async {
  // Create a pipeline that first logs requests, then handles API requests, then root requests.
  // The 'Cascade' middleware tries handlers in order until one returns a non-404 response.
  final handler = Pipeline()
      .addMiddleware(logRequests()) // Logs incoming requests
      .addHandler((Request request) {
        if (request.url.path == 'api') {
          return _apiHandler(request);
        }
        return _rootHandler(request);
      });

  final ip = InternetAddress.anyIPv4;
  final port = 8080;

  final server = await shelf_io.serve(handler, ip, port);
  print('Serving at http://${server.address.host}:${server.port}');
}
