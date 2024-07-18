import 'package:shelf/shelf.dart';
import 'package:shelf_router/shelf_router.dart';

import 'pages/pages.dart' as pages;

// ----------------------------------------------------------- Configure routes
final router = Router(notFoundHandler: pages.oopsHandler)
  ..get('/', pages.homePageHandler)
  ..get('/users/list', pages.getUsers)
  ..get('/echo/<message>', pages.echoHandler);

// ------------------------------------ Configure a pipeline that logs requests
final handler = Pipeline().addMiddleware(logRequests()).addHandler(router.call);
