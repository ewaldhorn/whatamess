import 'dart:convert';

import 'package:dart_rest_api/users/users.dart';
import 'package:shelf/shelf.dart';

Response getUsers(Request request) {
  // Convert the users list to JSON and return as response
  final jsonUsers = jsonEncode(users);
  return Response.ok(jsonUsers, headers: {'Content-Type': 'application/json'});
}
