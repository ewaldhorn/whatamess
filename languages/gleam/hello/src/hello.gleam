import extras/with_blocks
import extras/with_lists
import extras/with_type_alias
import gleam/io
import gleam/string as text
import use_qualified
import with_floats
import with_ints
import with_let
import with_strings
import with_types

// ------------------------------------------------------------------------------------------------
pub fn main() -> Nil {
  io.println("Hello from hello!")
  reverse_string()
  use_qualified.this_is_qualified()
  with_ints.demo()
  with_floats.demo()
  with_strings.demo()
  with_let.demo()
  let _a = with_types.demo()
  with_type_alias.demo()
  echo with_blocks.demo()
  with_lists.demo()

  Nil
}

// ------------------------------------------------------------------------------------------------
fn reverse_string() {
  io.println(text.reverse("This will look funny!"))
}
