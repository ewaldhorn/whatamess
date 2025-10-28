import argv
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import gleam_community/ansi
import simplifile

// ------------------------------------------------------------------------------------------------
// The file we will store todos in
const file_path = "my.todo"

// ------------------------------------------------------------------------------------------------
pub fn main() {
  case argv.load().arguments {
    ["add", title] -> add_todo(file_path, title)
    ["clear"] -> clear_todos(file_path)
    ["list"] -> list_todos(file_path)
    ["done", index_string] ->
      case parse_index(index_string) {
        Ok(index) -> done_todo(file_path, index)
        Error(_) -> {
          io.println("Error: 'done' requires a valid todo number.")
          Ok(Nil)
        }
      }
    _ -> {
      io.println("Usage:")
      io.println("  todo add <title>")
      io.println("  todo list")
      io.println("  todo done <number>")
      io.println("  todo clear")
      Ok(Nil)
    }
  }
}

// ------------------------------------------------------------------------------------------------
pub fn add_todo(file_path: String, title: String) -> Result(Nil, Nil) {
  simplifile.append(file_path, "[ ] " <> title <> "\n")
  |> result.map_error(with: fn(_) { Nil })
  |> result.replace(Nil)
}

// ------------------------------------------------------------------------------------------------
pub fn clear_todos(file_path: String) -> Result(Nil, Nil) {
  simplifile.write(file_path, "")
  |> result.map_error(with: fn(_) { Nil })
  |> result.replace(Nil)
}

// ------------------------------------------------------------------------------------------------
fn split_to_lines(content: String) -> List(String) {
  string.split(content, "\n")
}

// ------------------------------------------------------------------------------------------------
fn join_lines(lines: List(String)) -> String {
  string.join(lines, "\n")
}

// ------------------------------------------------------------------------------------------------
pub fn formatted_todos(file_path: String) -> Result(String, Nil) {
  case simplifile.read(file_path) {
    Ok(content) -> {
      Ok(content)
      |> result.map(split_to_lines)
      // Wrap list.filter in an anonymous function (fn(lines) { ... })
      // to correctly apply the predicate (fn(line) { line != "" }) to the list (lines)
      |> result.map(fn(lines) { list.filter(lines, fn(line) { line != "" }) })
      |> result.map(format_todo_lines)
      |> result.map(join_lines)
      |> result.replace_error(Nil)
    }
    Error(_) -> Ok("")
  }
}

// ------------------------------------------------------------------------------------------------
// No changes are strictly necessary here, but keeping it clean:
fn format_todo_lines(lines: List(String)) -> List(String) {
  list.index_map(lines, fn(line, index) {
    case line {
      "[ ] " <> rest -> int.to_string(index + 1) <> " " <> rest
      "[x] " <> rest ->
        int.to_string(index + 1) <> " " <> ansi.strikethrough(rest) <> " (done)"
      _ -> line
    }
  })
}

// ------------------------------------------------------------------------------------------------
pub fn list_todos(file_path: String) -> Result(Nil, Nil) {
  case formatted_todos(file_path) {
    Ok(output) -> {
      io.println(output)
      Ok(Nil)
    }
    Error(Nil) -> {
      io.println("Error reading todos.")
      Ok(Nil)
      // Returning Ok(Nil) here to avoid a program crash on I/O error
    }
  }
}

// ------------------------------------------------------------------------------------------------
fn parse_index(index_string: String) -> Result(Int, Nil) {
  int.parse(index_string)
  |> result.map(fn(i) { i - 1 })
  // Convert 1-based index (user) to 0-based (list)
  |> result.map_error(with: fn(_) { Nil })
}

// ------------------------------------------------------------------------------------------------
pub fn done_todo(file_path: String, index: Int) -> Result(Nil, Nil) {
  simplifile.read(file_path)
  |> result.map(split_to_lines)
  // Calls the new index_map-based function
  |> result.map(fn(lines) { update_line_to_done(lines, index) })
  |> result.map(join_lines)
  |> result.try(fn(content) { simplifile.write(file_path, content) })
  |> result.replace(Nil)
  |> result.map_error(with: fn(_) { Nil })
}

// ------------------------------------------------------------------------------------------------
fn update_line_to_done(lines: List(String), target_index: Int) -> List(String) {
  list.index_map(lines, fn(line, i) {
    // Check if the current index matches the index we want to change
    case i == target_index {
      True ->
        case line {
          "[ ] " <> rest -> "[x] " <> rest
          // Mark as done
          _ -> line
        }
      False ->
        // Not the target line, return it unchanged.
        line
    }
  })
}
// ------------------------------------------------------------------------------------------------
