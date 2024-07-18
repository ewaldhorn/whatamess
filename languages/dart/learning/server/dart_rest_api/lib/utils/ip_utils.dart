import 'dart:io';

// --------------------------------------------------------------- getIPAddress
// Use any available host or container IP (usually `0.0.0.0`).
InternetAddress getIPAddress() => InternetAddress.anyIPv4;

// -------------------------------------------------------------------- getPort
// // For running in containers, we respect the PORT environment variable.
int getPort() => int.parse(Platform.environment['PORT'] ?? '8080');
