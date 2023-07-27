void main() {
  // Maps (Key/Value pairs)
  var clients = {"ABC Trading": "Noel Barbarossa", "Top Turf": "Nadia Calico"};
  print(clients);
  print(clients["Top Turf"]);

  // Show Values
  print(clients.values);

  // Show Keys
  print(clients.keys);

  // Show Length
  print(clients.length);

  // Add something
  clients["Surf Up!"] = "Dan Vlu";
  print(clients);

  // Add many somethings
  clients.addAll({"Bo's":"Bo Tipple", "Nick's Bar":"Nick Cave"});
  print(clients);

  // Remove something
  clients.remove("Top Turf");
  print(clients);

  // Remove (clear) everything
  clients.clear();
  print(clients);
}
