import 'package:shelf_router/shelf_router.dart';

import 'pages/pages.dart' as pages;

// Configure routes.
final router = Router(notFoundHandler: pages.oopsHandler)
  ..get('/', pages.homePageHandler)
  ..get('/echo/<message>', pages.echoHandler);
