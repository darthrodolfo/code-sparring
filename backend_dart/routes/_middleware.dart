import 'dart:io';

import 'package:backend_dart/repositories/aircraft_repository.dart';
import 'package:backend_dart/repositories/sqlite_aircraft_repository.dart';
import 'package:dart_frog/dart_frog.dart';

final _repository = SqliteAircraftRepository(
  databasePath: Platform.environment['DB_PATH'] ?? 'data/aircraft.db',
);

Handler middleware(Handler handler) {
  return handler.use(provider<AircraftRepository>((_) => _repository));
}
