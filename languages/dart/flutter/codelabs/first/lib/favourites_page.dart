import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'main.dart';

class FavouritesPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    var appState = context.watch<MyAppState>();
    var favourites = <Widget>[];

    favourites.add(Padding(
      padding: const EdgeInsets.all(20),
      child: Text('You have '
          '${appState.favourites.length} favorites:'),
    ));

    for (var f in appState.favourites) {
      var tmp = ListTile(
        title: Text(f.asString),
      );
      favourites.add(tmp);
    }

    return Center(
      child: ListView(
        children: favourites,
      ),
    );
  }
}
