import gleam/list

// create a list of numbers
fn create_list(n: Int) -> List(Int) {
  let list_of_numbers =
    list.fold(
      list.range(1, n),
      [],
      fn(acc: List(Int), x: Int) { [x, ..acc] } 
    )
  list.reverse(list_of_numbers) // list needs to be reversed it's 10 9 8 7 6...
}

// sum list entries using recursion
fn sum_list(list: List(Int), total: Int) -> Int {
  case list {
    [first, ..rest] -> sum_list(rest, total + first)
    [] -> total
  }
}

fn square_list(list: List(Int), total: Int) -> Int {
  case list {
    [first, ..rest] -> square_list(rest, total + {first * first})
    []->total
  }
}

pub fn square_of_sum(n: Int) -> Int {
  let numbers = create_list(n)
  let total_sum = sum_list(numbers, 0)
  total_sum * total_sum
}

pub fn sum_of_squares(n: Int) -> Int {
  let numbers = create_list(n)
  square_list(numbers, 0)
}

pub fn difference(n: Int) -> Int {
  square_of_sum(n) - sum_of_squares(n)
}

