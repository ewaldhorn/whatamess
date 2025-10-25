import extras/arguments/with_arguments as argies
import extras/shorthand/label_shorthand as short
import extras/with_blocks
import extras/with_case
import extras/with_closures
import extras/with_constants
import extras/with_function_captures
import extras/with_functions
import extras/with_lists
import extras/with_pipes
import extras/with_type_alias
import gleam/io
import gleam/string as text
import use_qualified
import with_floats
import with_generics
import with_high_order
import with_ints
import with_let
import with_strings
import with_types

// ------------------------------------------------------------------------------------------------
pub fn main() -> Nil {
  io.println("Hello from hello!")
  reverse_string()
  do_a_little_something()
  with_case.demo()
  with_case.demo_assignment()

  Nil
}

// ------------------------------------------------------------------------------------------------
fn do_a_little_something() {
  use_qualified.this_is_qualified()
  with_ints.demo()
  with_floats.demo()
  with_strings.demo()
  with_let.demo()
  let _a = with_types.demo()
  with_type_alias.demo()
  do_a_little_something_more()
}

// ------------------------------------------------------------------------------------------------
fn do_a_little_something_more() {
  echo with_blocks.demo()
  with_lists.demo()
  with_constants.demo()
  echo with_functions.double(5)
  echo with_functions.multiply(2, 3)
  with_high_order.demo()
  with_closures.demo()
  with_function_captures.demo()
  with_generics.demo()
  with_pipes.demo()
  argies.demo()
  short.demo()
}

// ------------------------------------------------------------------------------------------------
fn reverse_string() {
  io.println(text.reverse("This will look funny!"))
}
