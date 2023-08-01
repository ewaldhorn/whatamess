import 'package:flutter/material.dart';

class TabletView extends StatefulWidget {
  const TabletView({super.key});

  @override
  State<TabletView> createState() => _TabletViewState();
}

class _TabletViewState extends State<TabletView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.red,
      appBar: AppBar(
        title: const Text('Tablet View'),
      ),
    );
  }
}
