import 'package:flutter/material.dart';

class ResponsiveUI extends StatelessWidget {
  const ResponsiveUI(
      {required this.desktopView,
      required this.tabletView,
      required this.mobileView,
      super.key});

  final Widget mobileView, tabletView, desktopView;

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) {
        if (constraints.maxWidth < 600) {
          return mobileView;
        } else if (constraints.maxWidth < 800) {
          return tabletView;
        } else {
          return desktopView;
        }
      },
    );
  }
}
