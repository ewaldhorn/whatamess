import gleam/io
import gleam/list

// List module
// The gleam/list standard library module contains functions for working with lists. A Gleam
// program will likely make heavy use of this module, the various functions serving as different
// types of loops over lists.
//
// map makes a new list by running a function on each element in a list.
//
// filter makes a new list containing only the elements for which a function returns true.
//
// fold combines all the elements in a list into a single value by running a function
// left-to-right on each element, passing the result of the previous call to the next call.
//
// find returns the first element in a list for which a function returns True.
//
// It's worth getting familiar with all the functions in this module when writing Gleam code,
// you'll be using them a lot!

// ------------------------------------------------------------------------------------------------
pub fn demo() {
  let ints = [0, 1, 2, 3, 4, 5]

  io.println("=== map ===")
  echo list.map(ints, fn(x) { x * 2 })

  io.println("=== filter ===")
  echo list.filter(ints, fn(x) { x % 2 == 0 })

  io.println("=== fold ===")
  echo list.fold(ints, 0, fn(count, e) { count + e })

  io.println("=== find ===")
  let _ = echo list.find(ints, fn(x) { x > 3 })
  let _ = echo list.find(ints, fn(x) { x > 13 })

  Nil
}
