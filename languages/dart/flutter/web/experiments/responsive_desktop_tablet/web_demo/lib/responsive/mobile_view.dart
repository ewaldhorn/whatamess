import 'package:flutter/material.dart';

class MobileView extends StatefulWidget {
  const MobileView({super.key});

  @override
  State<MobileView> createState() => _MobileViewState();
}

class _MobileViewState extends State<MobileView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.red[100],
      appBar: AppBar(
        title: const Text('Mobile View'),
      ),
      body: Column(
        children: [
          AspectRatio(
              aspectRatio: 16 / 9,
              child: Padding(
                padding: const EdgeInsets.all(10.0),
                child: Container(
                  height: 180,
                  color: Colors.red,
                ),
              )),
          Expanded(
              child: ListView.builder(
            itemCount: 10,
            itemBuilder: (context, index) {
              return Padding(
                padding: const EdgeInsets.all(10.0),
                child: Container(
                  height: 100,
                  color: Colors.blue,
                ),
              );
            },
          ))
        ],
      ),
    );
  }
}
