import 'package:flutter/material.dart';
import 'package:hello_web/config/size_config.dart';
import 'package:hello_web/style/colours.dart';

class Dashboard extends StatelessWidget {
  const Dashboard({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
          child: Row(
        children: [
          Expanded(
            flex: 1,
            child: Container(
              width: double.infinity,
              height: SizeConfig.screenHeight,
              color: AppColours.secondaryBg,
            ),
          ),
          Expanded(
            flex: 10,
            child: Container(
              width: double.infinity,
              height: SizeConfig.screenHeight,
            ),
          ),
          Expanded(
            flex: 4,
            child: Container(
              width: double.infinity,
              height: SizeConfig.screenHeight,
              color: AppColours.secondaryBg,
            ),
          )
        ],
      )),
    );
  }
}
