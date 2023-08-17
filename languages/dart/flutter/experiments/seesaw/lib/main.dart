import 'package:flutter/material.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Data Vis',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: VisApp(),
    );
  }
}

class VisApp extends StatelessWidget {
  const VisApp({super.key});

  @override
  Widget build(BuildContext context) {
    {
      return Scaffold(
        appBar: AppBar(
          title: const Text('Data Vis App'),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[Text('This is it!')],
          ),
        ),
      );
    }
  }
}
