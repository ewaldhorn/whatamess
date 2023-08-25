import 'package:flutter/material.dart';
import 'package:meals_app/models/meal.dart';
import 'package:meals_app/screens/meal_details.dart';
import 'package:meals_app/widgets/meal_item.dart';

class MealsScreen extends StatelessWidget {
  final String? title;
  final List<Meal> meals;

  const MealsScreen({super.key, this.title, required this.meals});

  void _onMealSelected(BuildContext context, Meal meal) => Navigator.of(context)
      .push(MaterialPageRoute(builder: (ctx) => MealDetailsScreen(meal: meal)));

  @override
  Widget build(BuildContext context) {
    Widget content = ListView.builder(
      itemCount: meals.length,
      itemBuilder: (ctx, idx) => MealItem(
          meal: meals[idx],
          onMealSelected: () => _onMealSelected(context, meals[idx])),
    );

    if (meals.isEmpty) {
      content = Center(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Text('No recipes found.'),
            SizedBox(height: 16),
            Text(
              'Try selecting a different category.',
              style: Theme.of(context)
                  .textTheme
                  .bodyLarge!
                  .copyWith(color: Theme.of(context).colorScheme.onBackground),
            )
          ],
        ),
      );
    }

    if (title == null) {
      return content;
    } else {
      return Scaffold(
        appBar: AppBar(
          title: Text(title!),
        ),
        body: content,
      );
    }
  }
}
