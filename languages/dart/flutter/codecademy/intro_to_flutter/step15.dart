import 'package:flutter/material.dart';

void main() {
  var myTransform = Matrix4.skewX(15);
  var myContainer = Container(transform: myTransform, child: const Text('Hi'));

  runApp(MaterialApp(
      home: Scaffold(
          body: Container(padding: const EdgeInsets.all(50), child: myContainer))));
}
