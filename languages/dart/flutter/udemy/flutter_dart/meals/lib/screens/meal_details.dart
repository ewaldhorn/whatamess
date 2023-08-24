import 'package:flutter/material.dart';
import 'package:meals_app/models/meal.dart';

class MealDetailsScreen extends StatelessWidget {
  final Meal meal;

  const MealDetailsScreen({super.key, required this.meal});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(meal.title),
        ),
        body: SingleChildScrollView(
          child: Column(children: [
            Text(meal.steps.join('\n'), style: TextStyle(color: Colors.white))
          ]),
        ));
  }
}
