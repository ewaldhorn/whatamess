import 'sqlite_service.dart';
import 'user.dart';
import 'user_repository.dart';

void main(List<String> args) {
  final sqliteService = SqliteService();
  sqliteService.connect();

  final userRepository = UserRepository(sqliteService);
  final user = User(userRepository);
  user.insertUser();
}
