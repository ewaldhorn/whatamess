// ------------------------------------------------------------------------------------------------
pub fn demo() {
  let numbers = [1, 2, 3, 4, 5]
  echo get_first_larger_than(numbers, 3)
  echo get_first_larger_than(numbers, 5)
}

// ------------------------------------------------------------------------------------------------
fn get_first_larger_than(numbers: List(Int), limit: Int) -> Int {
  case numbers {
    [first, ..] if first > limit -> first
    [_, ..rest] -> get_first_larger_than(rest, limit)
    [] -> 0
  }
}
