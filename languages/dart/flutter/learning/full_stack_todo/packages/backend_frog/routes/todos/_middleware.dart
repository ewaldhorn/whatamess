import 'package:backend_frog/data/in_memory_todo_repository.dart';
import 'package:dart_frog/dart_frog.dart';
import 'package:todo/todo.dart';

InMemoryTodoRepository? _todoData;

Handler middleware(Handler handler) {
  return handler.use(
    provider<TodoRepository>(
      (context) => _todoData ??= InMemoryTodoRepository(),
    ),
  );
}
