import 'package:flutter/material.dart';

void main() {
  var text = 'Inside Container\n' * 10;

  const myMargin = EdgeInsets.all(10);
  const myPadding = EdgeInsets.all(30);

  runApp(MaterialApp(
      home: Scaffold(
          backgroundColor: Colors.yellow,
          body: Container(
              margin: myMargin,
              padding: myPadding,
              color: Colors.red,
              child:
                  Text(text, style: const TextStyle(color: Colors.white))))));
}

