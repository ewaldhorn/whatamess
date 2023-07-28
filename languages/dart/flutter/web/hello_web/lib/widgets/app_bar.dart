import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

class AppBarActionItems extends StatelessWidget {
  const AppBarActionItems({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        IconButton(
            onPressed: () {},
            icon: SvgPicture.asset('assets/settings.svg', width: 20)),
        const SizedBox(width: 10.0),
        IconButton(
            onPressed: () {},
            icon: SvgPicture.asset('assets/profile.svg', width: 20))
      ],
    );
  }
}
