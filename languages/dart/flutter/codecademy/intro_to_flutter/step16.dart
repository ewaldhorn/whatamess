import 'package:flutter/material.dart';

void main() {
  var myMargin = const EdgeInsets.all(25);
  var myPadding = const EdgeInsets.all(25);
  var myDecoration = BoxDecoration(
                     border: Border.all(color: Colors.blueAccent, width: 10),
                     color: Colors.yellow);
  var myTransform = Matrix4.skewX(0.3);

  const icon1 = Icon(
    Icons.looks_one,
    color: Colors.red,
    size: 50.0,
  );
  const icon2 = Icon(
    Icons.looks_two,
    color: Colors.grey,
    size: 50.0,
  );
  const icon3 = Icon(
    Icons.looks_3,
    color: Colors.blue,
    size: 50.0,
  );

  runApp(MaterialApp(
      home: Scaffold(
          body: Column(children: [
    Container(padding: myPadding, child: icon1, decoration: myDecoration ),
    Container(margin: myMargin, child: icon2, decoration: myDecoration),
    Container(padding: myPadding, transform: myTransform, child: icon3)
  ]))));
}
