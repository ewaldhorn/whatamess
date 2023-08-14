import 'package:flutter/material.dart';

void main() {
  
  //              v Change me v
  const double size = 100.0;
  
  const icon1 = Icon(
    Icons.looks_one,
    color: Colors.red,
    size: size,
  );
  const icon2 = Icon(
    Icons.looks_two,
    color: Colors.grey,
    size: size,
  );
  const icon3 = Icon(
    Icons.looks_3,
    color: Colors.blue,
    size: size,
  );

  runApp(const MaterialApp(
    home: Scaffold(
      body: Row(
        children: [icon1, icon2, icon3]
      )
    )
  ));
}
