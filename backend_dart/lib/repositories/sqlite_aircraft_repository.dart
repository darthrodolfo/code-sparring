import 'dart:convert';
import 'dart:io';

import 'package:backend_dart/models/aircraft.dart';
import 'package:backend_dart/repositories/aircraft_repository.dart';
import 'package:decimal/decimal.dart';
import 'package:sqlite3/sqlite3.dart';

class SqliteAircraftRepository implements AircraftRepository {
  SqliteAircraftRepository({required String databasePath})
      : _db = _open(databasePath) {
    _initSchema();
  }

  final Database _db;

  static Database _open(String path) {
    final file = File(path);
    file.parent.createSync(recursive: true);
    return sqlite3.open(path);
  }

  void _initSchema() {
    _db.execute('''
      CREATE TABLE IF NOT EXISTS aircraft (
        id TEXT PRIMARY KEY,
        model TEXT NOT NULL,
        manufacturer TEXT NOT NULL,
        serial_number TEXT,
        year_of_manufacture INTEGER NOT NULL,
        price_millions TEXT NOT NULL,
        empty_weight_kg REAL NOT NULL,
        status TEXT NOT NULL,
        role TEXT NOT NULL,
        tags_json TEXT NOT NULL,
        first_flight_date TEXT NOT NULL,
        last_maintenance_time TEXT NOT NULL,
        base_location_json TEXT NOT NULL,
        specs_json TEXT NOT NULL,
        conflicts_json TEXT NOT NULL,
        metadata_json TEXT NOT NULL,
        estimated_units_produced INTEGER,
        estimated_active_units INTEGER,
        photo_url TEXT,
        manual_archive BLOB
      );
    ''');
  }

  @override
  List<Aircraft> getAll() {
    final rows = _db.select('SELECT * FROM aircraft ORDER BY model');
    return rows.map(_fromRow).toList(growable: false);
  }

  @override
  Aircraft? getById(String id) {
    final rows = _db.select('SELECT * FROM aircraft WHERE id = ?', [id]);
    if (rows.isEmpty) return null;
    return _fromRow(rows.first);
  }

  @override
  Aircraft add(Aircraft aircraft) {
    _db.execute('''
      INSERT INTO aircraft (
        id, model, manufacturer, serial_number, year_of_manufacture,
        price_millions, empty_weight_kg, status, role, tags_json,
        first_flight_date, last_maintenance_time, base_location_json, specs_json,
        conflicts_json, metadata_json, estimated_units_produced,
        estimated_active_units, photo_url, manual_archive
      ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    ''', _toParams(aircraft));
    return aircraft;
  }

  @override
  Aircraft? update(String id, Aircraft aircraft) {
    if (getById(id) == null) return null;
    _db.execute('''
      UPDATE aircraft SET
        model = ?, manufacturer = ?, serial_number = ?, year_of_manufacture = ?,
        price_millions = ?, empty_weight_kg = ?, status = ?, role = ?, tags_json = ?,
        first_flight_date = ?, last_maintenance_time = ?, base_location_json = ?,
        specs_json = ?, conflicts_json = ?, metadata_json = ?, estimated_units_produced = ?,
        estimated_active_units = ?, photo_url = ?, manual_archive = ?
      WHERE id = ?
    ''', [
      aircraft.model,
      aircraft.manufacturer,
      aircraft.serialNumber,
      aircraft.yearOfManufacture,
      aircraft.priceMillions.toString(),
      aircraft.emptyWeightKg,
      aircraft.status.name,
      aircraft.role.name,
      jsonEncode(aircraft.tags),
      aircraft.firstFlightDate.toIso8601String(),
      aircraft.lastMaintenanceTime.toIso8601String(),
      jsonEncode(aircraft.baseLocation.toJson()),
      jsonEncode(aircraft.specs.toJson()),
      jsonEncode(aircraft.conflicts.map((c) => c.toJson()).toList()),
      jsonEncode(aircraft.metadata),
      aircraft.estimatedUnitsProduced,
      aircraft.estimatedActiveUnits,
      aircraft.photoUrl,
      aircraft.manualArchive,
      id,
    ]);
    return aircraft;
  }

  @override
  bool delete(String id) {
    if (getById(id) == null) return false;
    _db.execute('DELETE FROM aircraft WHERE id = ?', [id]);
    return true;
  }

  List<Object?> _toParams(Aircraft a) => [
        a.id,
        a.model,
        a.manufacturer,
        a.serialNumber,
        a.yearOfManufacture,
        a.priceMillions.toString(),
        a.emptyWeightKg,
        a.status.name,
        a.role.name,
        jsonEncode(a.tags),
        a.firstFlightDate.toIso8601String(),
        a.lastMaintenanceTime.toIso8601String(),
        jsonEncode(a.baseLocation.toJson()),
        jsonEncode(a.specs.toJson()),
        jsonEncode(a.conflicts.map((c) => c.toJson()).toList()),
        jsonEncode(a.metadata),
        a.estimatedUnitsProduced,
        a.estimatedActiveUnits,
        a.photoUrl,
        a.manualArchive,
      ];

  Aircraft _fromRow(Row row) {
    final status = AircraftStatus.values.asNameMap()[row['status'] as String];
    final role = AircraftRole.values.asNameMap()[row['role'] as String];
    if (status == null || role == null) {
      throw const FormatException('Invalid enum value in SQLite row');
    }

    final baseLocation = Map<String, dynamic>.from(
      jsonDecode(row['base_location_json'] as String) as Map,
    );
    final specs = Map<String, dynamic>.from(
      jsonDecode(row['specs_json'] as String) as Map,
    );
    final conflicts = List<Map<String, dynamic>>.from(
      jsonDecode(row['conflicts_json'] as String) as List,
    );
    final metadata = Map<String, String>.from(
      jsonDecode(row['metadata_json'] as String) as Map,
    );

    final rawBlob = row['manual_archive'];
    final manualArchive =
        rawBlob == null ? null : List<int>.from(rawBlob as List<int>);

    return Aircraft(
      id: row['id'] as String,
      model: row['model'] as String,
      manufacturer: row['manufacturer'] as String,
      serialNumber: row['serial_number'] as String?,
      yearOfManufacture: row['year_of_manufacture'] as int,
      priceMillions: Decimal.parse(row['price_millions'] as String),
      emptyWeightKg: (row['empty_weight_kg'] as num).toDouble(),
      status: status,
      role: role,
      tags: List<String>.from(jsonDecode(row['tags_json'] as String) as List),
      firstFlightDate: DateTime.parse(row['first_flight_date'] as String),
      lastMaintenanceTime:
          DateTime.parse(row['last_maintenance_time'] as String),
      baseLocation: GeoLocation.fromJson(baseLocation),
      specs: AircraftSpecs.fromJson(specs),
      conflicts: conflicts.map(ConflictHistory.fromJson).toList(),
      metadata: metadata,
      estimatedUnitsProduced: row['estimated_units_produced'] as int?,
      estimatedActiveUnits: row['estimated_active_units'] as int?,
      photoUrl: row['photo_url'] as String?,
      manualArchive: manualArchive,
    );
  }
}
