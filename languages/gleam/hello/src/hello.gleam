import gleam/io
import gleam/string as text
import use_qualified
import with_ints

// ------------------------------------------------------------------------------------------------
pub fn main() -> Nil {
  io.println("Hello from hello!")
  reverse_string()
  use_qualified.this_is_qualified()
  with_ints.demo()
}

// ------------------------------------------------------------------------------------------------
fn reverse_string() {
  io.println(text.reverse("This will look funny!"))
}
