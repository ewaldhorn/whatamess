import gleeunit
import internal

// ------------------------------------------------------------------------------------------------
pub fn main() {
  gleeunit.main()
}

// ------------------------------------------------------------------------------------------------
pub fn format_pair_test() {
  let greeting = internal.format_pair("hello", "world")
  assert greeting == "hello=world"
}
