import 'package:backend_dart/models/aircraft.dart';

class AircraftStore {
  final Map<String, Aircraft> _store = {};

  List<Aircraft> getAll() => List.unmodifiable(_store.values.toList());

  Aircraft? getById(String id) => _store[id];

  Aircraft add(Aircraft aircraft) {
    _store[aircraft.id] = aircraft;
    return aircraft;
  }

  Aircraft? update(String id, Aircraft aircraft) {
    if (!_store.containsKey(id)) return null;
    _store[id] = aircraft;
    return aircraft;
  }

  bool delete(String id) => _store.remove(id) != null;
}
