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

// ------------------------------------------------------------------- addUser
Future<Response> addUser(Request request) async {
  final payload = await request.readAsString();
  final data = jsonDecode(payload);

  // mimic a database insert
  data['id'] = (users.length + 1).toString();
  users.add(data);

  return Response.ok(
    {
      'message': '${data['id']} added successfully',
      'status': 'success',
    },
    headers: {'Content-Type': 'application/json'},
  );
}

// ----------------------------------------------------------------- updateUser
Future<Response> updateUser(Request request) async {
  final id = request.params['id'];
  final existingUserIndex = users.indexWhere((u) => u['id'] == id);

  if (existingUserIndex == -1) {
    return Response.notFound('User not found');
  }

  final payload = await request.readAsString();

  final updateData = jsonDecode(payload);

  // Update the user data accordingly
  users[existingUserIndex] = updateData;

  return Response.ok(
    jsonEncode(users[existingUserIndex]),
    headers: {'Content-Type': 'application/json'},
  );
}

// ----------------------------------------------------------------- deleteUser
Response deleteUser(Request request) {
  final id = request.params['id'];
  final userIndex = users.indexWhere((u) => u['id'] == id);
  if (userIndex == -1) {
    return Response.notFound('User not found');
  }
  users.removeAt(userIndex);
  return Response.ok('User deleted successfully');
}
