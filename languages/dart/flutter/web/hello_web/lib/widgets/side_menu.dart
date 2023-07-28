import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:hello_web/config/size_config.dart';
import 'package:hello_web/style/colours.dart';

class SideMenu extends StatelessWidget {
  const SideMenu({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Drawer(
      elevation: 0,
      child: Container(
        width: double.infinity,
        height: SizeConfig.screenHeight,
        color: AppColours.secondaryBg,
        child: SingleChildScrollView(
          child: Column(
            children: [
              Container(
                height: 100,
                alignment: Alignment.topCenter,
                padding: const EdgeInsets.only(top: 20),
                child: SizedBox(
                  width: 35,
                  height: 35,
                  child: SvgPicture.asset('assets/folder-plus.svg'),
                ),
              ),
              IconButton(
                onPressed: () {},
                icon: SvgPicture.asset(
                  'assets/house.svg',
                  color: AppColours.iconGray,
                ),
                iconSize: 20,
                padding: const EdgeInsets.symmetric(vertical: 20.0),
              ),
              IconButton(
                onPressed: () {},
                icon: SvgPicture.asset(
                  'assets/chart.svg',
                  color: AppColours.iconGray,
                ),
                iconSize: 20,
                padding: const EdgeInsets.symmetric(vertical: 20.0),
              ),
              IconButton(
                onPressed: () {},
                icon: SvgPicture.asset(
                  'assets/bill.svg',
                  color: AppColours.iconGray,
                ),
                iconSize: 20,
                padding: const EdgeInsets.symmetric(vertical: 20.0),
              )
            ],
          ),
        ),
      ),
    );
  }
}
