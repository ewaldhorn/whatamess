import 'package:shelf/shelf.dart';
import 'package:shelf_router/shelf_router.dart';

import 'pages/pages.dart' as pages;

// ----------------------------------------------------------- Configure routes
final router = Router(notFoundHandler: pages.oopsHandler)
  ..delete('/api/v1/users/<id>', pages.deleteUser)
  ..get('/', pages.homePageHandler)
  ..get('/users/list', pages.getUsers)
  ..get('/api/v1/users/<id>', pages.getUserById)
  ..get('/echo/<message>', pages.echoHandler)
  ..post('/api/v1/users', pages.addUser)
  ..put('/api/v1/users/<id>', pages.updateUser);

// ------------------------------------ Configure a pipeline that logs requests
final handler = Pipeline().addMiddleware(logRequests()).addHandler(router.call);
