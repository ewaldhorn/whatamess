abstract class Store {
  int? storedNullableValue;

  /// If [storedNullableValue] is currently `null`,
  /// set it to the result of [calculateValue] 
  /// or `0` if [calculateValue] returns `null`.
  void updateStoredValue() {
    var val = calculateValue();
    storedNullableValue ??= (val == null) ? 0 : calculateValue();
  }

  /// Calculates a value to be used,
  /// potentially `null`.
  int? calculateValue();
}
