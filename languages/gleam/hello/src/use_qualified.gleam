import gleam/io.{println}

pub fn this_is_qualified() {
  // Use the function in a qualified fashion
  io.println("This is qualified")

  // Or an unqualified fashion
  println("This is unqualified")
}
