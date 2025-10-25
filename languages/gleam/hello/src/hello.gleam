import gleam/io
import gleam/string as text
import help_me_run
import patterns/alternative_patterns
import patterns/list_patterns
import patterns/matching
import patterns/pattern_alias
import patterns/string_patterns
import recursion/list_recursion
import recursion/recurse_this
import recursion/tail_calls

// ------------------------------------------------------------------------------------------------
pub fn main() -> Nil {
  io.println("Hello from hello!")
  reverse_string()
  help_me_run.do_a_little_something()
  do_patterns()
  matching.demo()
  do_recursions()

  Nil
}

// ------------------------------------------------------------------------------------------------
fn do_patterns() {
  string_patterns.demo()
  list_patterns.demo()
  alternative_patterns.demo()
  pattern_alias.demo()
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
