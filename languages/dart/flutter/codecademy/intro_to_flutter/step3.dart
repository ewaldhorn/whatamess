import 'package:flutter/material.dart';

void main() {
  
    const bigRedStyle = TextStyle(color:Colors.red, fontSize:24);
  
    runApp(
    const MaterialApp(
      home: Scaffold(
        body: Text('Codecademy: Flutter Edition', style: bigRedStyle)
      )
    )
  );
}
