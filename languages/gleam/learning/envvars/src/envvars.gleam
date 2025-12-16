import argv
import envoy
import gleam/io
import gleam/result
import helpers

// ------------------------------------------------------------------------------------------------
pub fn main() {
  case argv.load().arguments {
    ["get", name] -> get(name)
    _ -> io.println("Usage: vars get <name>")
  }
}

// ------------------------------------------------------------------------------------------------
fn get(name: String) -> Nil {
  let value = envoy.get(name) |> result.unwrap("")
  io.println(helpers.format_pair(name, value))
}
