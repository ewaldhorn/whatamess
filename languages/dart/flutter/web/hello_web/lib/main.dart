import 'package:flutter/material.dart';
import 'package:hello_web/dashboard.dart';
import 'package:hello_web/style/colours.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: true,
      title: 'Flutter Responsive Dashboard',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
        scaffoldBackgroundColor: AppColours.primaryBg
      ),
      home: const Dashboard(),
    );
  }
}

