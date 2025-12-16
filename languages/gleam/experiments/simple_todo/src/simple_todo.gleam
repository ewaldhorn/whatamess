import gleam/io
import gleam/list
import gleam/string
import gleam/result

pub type Todo {
  Todo(id: Int, text: String, completed: Bool)
}
pub type TodoApp {
  TodoApp(todos: List(Todo), next_id: Int)
}
pub fn new_app() -> TodoApp {
  TodoApp(todos: [], next_id: 1)
}
pub fn add_todo(app: TodoApp, text: String) -> TodoApp {
  let new_todo = Todo(id: app.next_id, text: text, completed: False)
  TodoApp(
    todos: [new_todo, ..app.todos],
    next_id: app.next_id + 1
  )
}
pub fn complete_todo(app: TodoApp, id: Int) -> Result(TodoApp, String) {
  case list.find(app.todos, fn(todo) { todo.id == id }) {
    Ok(_) -> {
      let updated_todos = list.map(app.todos, fn(t) { // <-- Renamed 'todo' to 't'
        case t.id == id {
          True -> Todo(..t, completed: True) // <-- Fix applied here
          False -> t
        }
      })
      Ok(TodoApp(..app, todos: updated_todos))
    }
    Error(_) -> Error("Todo not found")
  }
}

pub fn remove_todo(app: TodoApp, id: Int) -> TodoApp {
  let filtered_todos = list.filter(app.todos, fn(todo) { todo.id != id })
  TodoApp(..app, todos: filtered_todos)
}
pub fn list_todos(app: TodoApp) -> Nil {
  case app.todos {
    [] -> io.println("No todos yet!")
    todos -> {
      io.println("Your todos:")
      list.each(todos, fn(todo) {
        let status = case todo.completed {
          True -> "[✓]"
          False -> "[ ]"
        }
        io.println(status <> " " <> string.inspect(todo.id) <> ". " <> todo.text)
      })
    }
  }
}
pub fn main() {
  let initial_app = new_app()
    |> add_todo("Learn Gleam")
    |> add_todo("Build a todo app")
    |> add_todo("Compare with Elixir")

  list_todos(initial_app)

  let updated_app = case complete_todo(initial_app, 1) {
    Ok(app) -> app // Use a new variable name or pipe the result
    Error(msg) -> {
      io.println("Error: " <> msg)
      initial_app // Return the original app on error
    }
  }

  io.println("\nAfter completing first todo:")
  list_todos(updated_app)
}
