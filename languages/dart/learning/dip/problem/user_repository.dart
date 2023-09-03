import 'sqlite_service.dart';

class UserRepository {
  final SqliteService service;

  UserRepository(this.service);

  void connect() {
    service.connect();
  }

  void insert() {
    service.insert();
  }
}
