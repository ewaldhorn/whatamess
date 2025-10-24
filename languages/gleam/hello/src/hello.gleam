import gleam/io
import gleam/string as text
import use_qualified

// ------------------------------------------------------------------------------------------------
pub fn main() -> Nil {
  io.println("Hello from hello!")
  reverse_string()
  use_qualified.this_is_qualified()
}

// ------------------------------------------------------------------------------------------------
fn reverse_string() {
  io.println(text.reverse("This will look funny!"))
}
