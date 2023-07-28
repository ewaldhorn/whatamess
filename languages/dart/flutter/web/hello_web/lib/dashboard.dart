import 'package:flutter/material.dart';
import 'package:hello_web/config/size_config.dart';
import 'package:hello_web/style/colours.dart';
import 'package:hello_web/widgets/app_bar.dart';
import 'package:hello_web/widgets/side_menu.dart';

class Dashboard extends StatelessWidget {
  const Dashboard({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
          child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Expanded(
            flex: 1,
            child: SideMenu(),
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
              padding:
                  const EdgeInsets.symmetric(vertical: 30.0, horizontal: 30.0),
              child: const Column(
                children: [AppBarActionItems()],
              ),
            ),
          )
        ],
      )),
    );
  }
}
