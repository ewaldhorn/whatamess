import 'package:flutter/material.dart';

void main() {
  runApp(
    const MaterialApp(
      home: Scaffold(
        body: Align(
          alignment: Alignment.bottomRight,
          child: Icon(
            Icons.center_focus_strong,
            color: Colors.red,
            size: 100.0,
          )
        )
      )
    )
  );
}
