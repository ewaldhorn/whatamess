import 'dart:convert';

import 'package:dart_rest_api/users/users.dart';
import 'package:shelf/shelf.dart';
import 'package:shelf_router/shelf_router.dart';

// ------------------------------------------------------------------- getUsers
Response getUsers(Request request) {
  // Convert the users list to JSON and return as response
  final jsonUsers = jsonEncode(users);
  return Response.ok(jsonUsers, headers: {'Content-Type': 'application/json'});
}

// ---------------------------------------------------------------- getUserByID
Response getUserById(Request request) {
  final id = request.params['id'];

  final user = users.firstWhere(
    (obj) => obj['id'] == id,
    orElse: () => <String, String>{}, // return an empty map if no item is found
  );

  if (user['id'] == null) {
    return Response.notFound('User not found');
  }

  return Response.ok(
    jsonEncode(user),
    headers: {'Content-Type': 'application/json'},
  );
}
