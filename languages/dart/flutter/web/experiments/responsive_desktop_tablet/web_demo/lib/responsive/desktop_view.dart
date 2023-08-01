import 'package:flutter/material.dart';

class DesktopView extends StatefulWidget {
  const DesktopView({super.key});

  @override
  State<DesktopView> createState() => _DesktopViewState();
}

class _DesktopViewState extends State<DesktopView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.red,
      appBar: AppBar(
        title: const Text('Desktop View'),
      ),
    );
  }
}
