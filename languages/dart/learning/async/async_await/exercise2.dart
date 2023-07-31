// Implement changeUsername here
Future<String> changeUsername() async {
  try {
    var tmp = await fetchNewUsername();
    return tmp;
  } catch(err) {
    return err.toString();
  }
}
