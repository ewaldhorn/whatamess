import 'package:flame/components.dart';
import 'package:flame/palette.dart';
import 'package:flutter/painting.dart';

JoystickComponent joystick = JoystickComponent(
  knob: CircleComponent(
    radius: 30,
    paint: BasicPalette.red
        .withAlpha(200)
        .paint(), // the inner circle which the user will drag
  ),
  background: CircleComponent(
    radius: 80,
    paint: BasicPalette.red.withAlpha(100).paint(), // the outer
  ),
  margin:
      const EdgeInsets.only(left: 40, bottom: 40), // position of the joystick
);
