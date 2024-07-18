import 'package:shelf/shelf.dart';
import 'package:shelf/shelf_io.dart';

import 'router.dart';
import 'package:dart_rest_api/utils/ip_utils.dart';

// ----------------------------------------------------------------------- main
void main(List<String> args) async {
  // Configure a pipeline that logs requests.
  final handler =
      Pipeline().addMiddleware(logRequests()).addHandler(router.call);
  final server = await serve(handler, getIPAddress(), getPort());

  print('Server listening on port ${server.port}');
}
