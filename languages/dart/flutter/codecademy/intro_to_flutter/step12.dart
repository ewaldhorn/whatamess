import 'package:flutter/material.dart';

void main() {
  //           Add decoration here v
  var myDecoration = BoxDecoration(
      color: Colors.orange,
      border: Border.all(color: Colors.yellowAccent, width: 8),
      borderRadius: BorderRadius.circular(5));

  runApp(MaterialApp(
      home: Scaffold(
          body: Container(
              decoration: myDecoration,
              child: const Icon(Icons.brush, size: 50)))));
}

