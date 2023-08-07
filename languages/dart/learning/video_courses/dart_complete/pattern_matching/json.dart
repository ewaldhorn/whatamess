var json = {
  'user': ['Bob', 40]
};
void main() {
  var old_way_name = json['user']![0];
  var old_way_age = json['user']![1];

  print("The old way says it's $old_way_name, aged $old_way_age.");

  var {'user': [name, age]} = json;
  print("The new way says it's $name, aged $age.");

  old_school();
  new_way();
}

void old_school() {
  var json = {
    'user': ['Bob', 40]
  };

  if (json is Map<String, dynamic> &&
      json.length == 1 &&
      json.containsKey('user')) {
    var user = json['user'];
    if (user is List<dynamic> &&
        user?.length == 2 &&
        user?[0] is String &&
        user?[1] is int) {
      var name = user?[0] as String;
      var age = user?[1] as int;
      print("User $name is $age years old.");
    }
  }
}

void new_way() {
  var json = {
    'user': ['Bob', 40]
  };

  switch (json) {
    case {'user': [String name, int age]}:
      print("User $name is $age years old.");
    default:
      throw 'Unknown JSON structure';
  }
}
