import 'package:flutter/material.dart';
import 'package:visual_sort/bubble_sort/bubble_sort_widget.dart';

import 'selection_sort/selection_sort_widget.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Visual Sorter',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepOrange),
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'Flutter Visual Sorter'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  // --------------------------------------------------------------------------
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      body: const Center(
        child: SingleChildScrollView(
          child: Column(
            children: [
              Text('Bubble Sort'),
              BubbleSorter(),
              SizedBox(height: 32),
              Text('Selection Sort'),
              SelectionSorter(),
            ],
          ),
        ),
      ),
    );
  }
}
