import gleam/io
import gleam/string as text
import help_me_run
import patterns/list_patterns
import patterns/string_patterns
import recursion/list_recursion
import recursion/recurse_this
import recursion/tail_calls

// ------------------------------------------------------------------------------------------------
pub fn main() -> Nil {
  io.println("Hello from hello!")
  reverse_string()
  help_me_run.do_a_little_something()
  string_patterns.demo()
  list_patterns.demo()
  do_recursions()

  Nil
}

// ------------------------------------------------------------------------------------------------
fn do_recursions() {
  recurse_this.demo()
  tail_calls.demo()
  list_recursion.demo()
}

// ------------------------------------------------------------------------------------------------
fn reverse_string() {
  io.println(text.reverse("This will look funny!"))
}
