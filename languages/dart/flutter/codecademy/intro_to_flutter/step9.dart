import 'package:flutter/material.dart';

void main() {
  const icon1 = Icon(
    Icons.looks_one,
    color: Colors.red,
    size: 100.0,
  );
  const icon2 = Icon(
    Icons.looks_two,
    color: Colors.grey,
    size: 100.0,
  );
  const icon3 = Icon(
    Icons.looks_3,
    color: Colors.blue,
    size: 100.0,
  );

  runApp(const MaterialApp(
    home: Scaffold(
      body: Column(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [icon1, icon2, icon3]
      )
    )
  ));
}
