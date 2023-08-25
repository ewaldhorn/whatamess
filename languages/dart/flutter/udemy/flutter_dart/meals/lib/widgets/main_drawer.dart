import 'package:flutter/material.dart';

class MainDrawer extends StatelessWidget {
  const MainDrawer({super.key});

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: Column(
        children: [
          DrawerHeader(
              decoration: BoxDecoration(
                  gradient: LinearGradient(colors: [
                Theme.of(context).colorScheme.primaryContainer,
                Theme.of(context).colorScheme.primaryContainer.withOpacity(0.8),
                Theme.of(context).colorScheme.primaryContainer.withOpacity(0.6),
                Theme.of(context).colorScheme.primaryContainer.withOpacity(0.4)
              ], begin: Alignment.topLeft, end: Alignment.bottomRight)),
              padding: const EdgeInsets.all(20),
              child: Row(
                children: [
                  Icon(Icons.fastfood,
                      size: 48, color: Theme.of(context).colorScheme.primary),
                  const SizedBox(width: 18),
                  Text(
                    'Cooking up!',
                    style: Theme.of(context)
                        .textTheme
                        .titleLarge!
                        .copyWith(color: Theme.of(context).colorScheme.primary),
                  )
                ],
              ))
        ],
      ),
    );
  }
}
