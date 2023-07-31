// Part 1
String addHello(String user) {
  return "Hello $user";
}

// Part 2
// You can call the provided async function fetchUsername()
// to return the username.
Future<String> greetUser() async {
  var un = await fetchUsername();
  return addHello(un);
}

// Part 3
// You can call the provided async function logoutUser()
// to log out the user.
Future<String> sayGoodbye() async {
  try {
    var res = await logoutUser();
    return "$res Thanks, see you next time";
  } catch (err) {
    return "No goodbyes for you!";
  }
}
