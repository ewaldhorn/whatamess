import 'package:flame/components.dart';
import 'package:flame/palette.dart';
import 'package:flutter/material.dart';

JoystickComponent joystick = JoystickComponent(
  knob: CircleComponent(
    radius: 30,
    paint: BasicPalette.red.withAlpha(200).paint(),
  ),
  background: CircleComponent(
      radius: 80,
      paint: BasicPalette.lightGray.withAlpha(100).paint(),
      children: [
        PolygonComponent.relative([
          // up arrow
          Vector2(0.0, -0.5),
          Vector2(0.5, 0.0),
          Vector2(-0.5, 0.0),
        ],
            position: Vector2(60, 5),
            paint: BasicPalette.darkGray.withAlpha(100).paint(),
            parentSize: Vector2(80, 80)),
        PolygonComponent.relative([
          // down arrow
          Vector2(0.5, 1.0),
          Vector2(1.0, 0.5),
          Vector2(0.0, 0.5),
        ],
            position: Vector2(60, 135),
            paint: BasicPalette.darkGray.withAlpha(100).paint(),
            parentSize: Vector2(80, 80)),
        PolygonComponent.relative([
          // left arrow
          Vector2(-0.5, 0.0),
          Vector2(0.0, -0.5),
          Vector2(0.0, 0.5),
        ],
            position: Vector2(5, 60),
            paint: BasicPalette.darkGray.withAlpha(100).paint(),
            parentSize: Vector2(80, 80)),
        PolygonComponent.relative([
          // down arrow
          Vector2(0.5, 0.0),
          Vector2(0.0, 0.5),
          Vector2(0.0, -0.5),
        ],
            position: Vector2(135, 60),
            paint: BasicPalette.darkGray.withAlpha(100).paint(),
            parentSize: Vector2(80, 80))
      ]),
  margin: const EdgeInsets.only(left: 40, bottom: 40),
);
