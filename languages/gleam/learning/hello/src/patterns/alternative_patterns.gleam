import gleam/int

// ------------------------------------------------------------------------------------------------
pub fn demo() {
  let number = int.random(12)
  echo number

  let result = case number {
    2 | 4 | 6 | 8 -> "This is an even number"
    1 | 3 | 5 | 7 -> "This is an odd number"
    _ -> "I'm not sure"
  }
  echo result
}
