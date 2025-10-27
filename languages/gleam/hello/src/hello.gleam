import gleam/io
import gleam/string as text
import help_me_run
import patterns/alternative_patterns
import patterns/guards
import patterns/list_patterns
import patterns/matching
import patterns/pattern_alias
import patterns/string_patterns
import records/with_records
import recursion/list_recursion
import recursion/recurse_this
import recursion/tail_calls
import results/with_results
import types/custom_types

// ------------------------------------------------------------------------------------------------
pub fn main() -> Nil {
  io.println("Hello from hello!")
  reverse_string()
  help_me_run.do_a_little_something()
  matching.demo()
  do_recursions()
  do_patterns()
  do_tuples()
  custom_types.demo()
  with_records.demo()

  // Nil in Gleam
  let x = Nil
  echo x
  let result = io.println("Hello!")
  echo result == Nil

  with_results.demo()
  bit_arrays()

  // Ensure we return a Nil
  Nil
}

// ------------------------------------------------------------------------------------------------
fn do_patterns() {
  string_patterns.demo()
  list_patterns.demo()
  alternative_patterns.demo()
  pattern_alias.demo()
  guards.demo()
}

// ------------------------------------------------------------------------------------------------
fn do_recursions() {
  recurse_this.demo()
  tail_calls.demo()
  list_recursion.demo()
}

// ------------------------------------------------------------------------------------------------
fn do_tuples() {
  let triple = #(1, 2.2, "three")
  echo triple

  let #(a, _, _) = triple
  echo a
  echo triple.1
}

// ------------------------------------------------------------------------------------------------
fn reverse_string() {
  io.println(text.reverse("This will look funny!"))
}

// ------------------------------------------------------------------------------------------------
fn bit_arrays() {
  // 8 bit int. In binary: 00000011
  echo <<3>>
  echo <<3>> == <<3:size(8)>>

  // 16 bit int. In binary: 0001100000000011
  echo <<6147:size(16)>>

  // A bit array of UTF8 data
  echo <<"Hello, Joe!":utf8>>

  // Concatenation
  let first = <<4>>
  let second = <<2>>
  echo <<first:bits, second:bits>>
}
