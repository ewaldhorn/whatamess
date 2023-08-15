String singleWhere(Iterable<String> items) {
  return items.singleWhere(
      (element) => element.startsWith('M') && element.contains('a'));
}

