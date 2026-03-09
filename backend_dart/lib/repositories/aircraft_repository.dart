import 'package:backend_dart/models/aircraft.dart';

abstract class AircraftRepository {
  List<Aircraft> getAll();
  Aircraft? getById(String id);
  Aircraft add(Aircraft aircraft);
  Aircraft? update(String id, Aircraft aircraft);
  bool delete(String id);
}
