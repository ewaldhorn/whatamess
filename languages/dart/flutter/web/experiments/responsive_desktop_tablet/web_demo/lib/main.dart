import 'package:flutter/material.dart';
import 'package:web_demo/responsive/desktop_view.dart';
import 'package:web_demo/responsive/mobile_view.dart';
import 'package:web_demo/responsive/responsive.dart';
import 'package:web_demo/responsive/tablet_view.dart';

void main() {
  runApp(ResponsiveDemo());
}

class ResponsiveDemo extends StatefulWidget {
  const ResponsiveDemo({super.key});

  @override
  State<ResponsiveDemo> createState() {
    return _ResponsiveDemoState();
  }
}

class _ResponsiveDemoState extends State<ResponsiveDemo> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(primarySwatch: Colors.red),
      home: ResponsiveUI(
        desktopView: DesktopView(),
        tabletView: TabletView(),
        mobileView: MobileView(),
      ),
    );
  }
}
