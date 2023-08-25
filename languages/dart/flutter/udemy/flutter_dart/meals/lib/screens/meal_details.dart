import 'package:flutter/material.dart';
import 'package:meals_app/models/meal.dart';

class MealDetailsScreen extends StatelessWidget {
  final Meal meal;
  final void Function(Meal meal) onToggleFavourite;

  const MealDetailsScreen(
      {super.key, required this.meal, required this.onToggleFavourite});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text(meal.title),
          actions: [
            IconButton(
              onPressed: () => onToggleFavourite(meal),
              icon: Icon(Icons.star),
            )
          ],
        ),
        body: SingleChildScrollView(
          child: Column(children: [
            Image.network(
              meal.imageUrl,
              height: 300,
              width: double.infinity,
              fit: BoxFit.cover,
            ),
            const SizedBox(height: 14),
            Text('Ingredients',
                style: Theme.of(context)
                    .textTheme
                    .titleLarge!
                    .copyWith(color: Theme.of(context).colorScheme.primary)),
            Text(meal.ingredients.join('\n'),
                style: TextStyle(
                    color: Theme.of(context).colorScheme.onBackground)),
            const SizedBox(height: 28),
            Text('Steps',
                style: Theme.of(context)
                    .textTheme
                    .titleLarge!
                    .copyWith(color: Theme.of(context).colorScheme.primary)),
            Text(
              meal.steps.join('\n'),
              style: TextStyle(color: Colors.white),
            ),
          ]),
        ));
  }
}
