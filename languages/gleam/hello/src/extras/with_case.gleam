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
