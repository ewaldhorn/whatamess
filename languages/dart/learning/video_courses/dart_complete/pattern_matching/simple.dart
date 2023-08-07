(String, int) userInfo(Map<String, dynamic> json) {
  return (json['name'] as String, json['age'] as int);
}

var json = {'name': 'Bob', 'age': 21};
var json2 = {'name': 'Sally', 'age': 20};

void main() {
  // using positions
  var info = userInfo(json);
  var name = info.$1;
  var age = info.$2;

  print("Name: $name, Age: $age");

  // using records and pattern matching to support multiple return values
  (name, age) = userInfo(json2);
  print("Name: $name, Age: $age");
}
