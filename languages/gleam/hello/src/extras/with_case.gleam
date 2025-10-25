import gleam/int

// ------------------------------------------------------------------------------------------------
pub fn demo() {
  let x = int.random(5)
  echo x

  let result = case x {
    // Match specific values
    0 -> "Zero"
    1 -> "One"

    // Match any other value
    _ -> "Other"
  }
  echo result
}

// ------------------------------------------------------------------------------------------------
pub fn demo_assignment() {
  let result = case int.random(5) {
    // Match specific values
    0 -> "Zero"
    1 -> "One"

    // Match any other value and assign it to a variable
    other -> "It is " <> int.to_string(other)
  }
  echo result
}
