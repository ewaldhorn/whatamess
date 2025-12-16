// ------------------------------------------------------------------------------------------------
pub fn demo() {
  let sum = sum_list([18, 56, 35, 85, 91], 0)
  echo sum
}

// ------------------------------------------------------------------------------------------------
// recurses over the list, grabbing the first element each time and passing the rest as an argument
// along with the running total
fn sum_list(list: List(Int), total: Int) -> Int {
  case list {
    [first, ..rest] -> sum_list(rest, total + first)
    [] -> total
  }
}
