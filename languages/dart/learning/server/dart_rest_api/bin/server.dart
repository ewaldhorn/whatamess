import 'package:shelf/shelf_io.dart';

import 'server_helpers.dart';
import 'package:dart_rest_api/utils/ip_utils.dart';

// ----------------------------------------------------------------------- main
void main(List<String> args) async {
  print('Starting up server');

  final server = await serve(handler, getIPAddress(), getPort());

  print('Server listening on port ${server.port}');
}
