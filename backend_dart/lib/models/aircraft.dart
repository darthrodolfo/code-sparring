import 'package:decimal/decimal.dart';

enum AircraftRole { fighter, bomber, transport, reconnaissance, trainer, drone }

enum AircraftStatus { active, maintenance, retired, stored }

class GeoLocation {
  final double latitude;
  final double longitude;

  GeoLocation({required this.latitude, required this.longitude});

  factory GeoLocation.fromJson(Map<String, dynamic> json) {
    final latitudeValue = json['latitude'];
    final longitudeValue = json['longitude'];
    if (latitudeValue is! num) {
      throw const FormatException('baseLocation.latitude must be a number');
    }
    if (longitudeValue is! num) {
      throw const FormatException('baseLocation.longitude must be a number');
    }
    return GeoLocation(
      latitude: latitudeValue.toDouble(),
      longitude: longitudeValue.toDouble(),
    );
  }

  Map<String, dynamic> toJson() => {
        'latitude': latitude,
        'longitude': longitude,
      };
}

class AircraftSpecs {
  final int maxSpeedKmh;
  final double wingspanMeters;
  final int rangeKm;
  final int? maxAltitudeMeters;
  final Duration flightEndurance;

  AircraftSpecs({
    required this.maxSpeedKmh,
    required this.wingspanMeters,
    required this.rangeKm,
    this.maxAltitudeMeters,
    required this.flightEndurance,
  });

  factory AircraftSpecs.fromJson(Map<String, dynamic> json) {
    final maxSpeedKmhValue = json['maxSpeedKmh'];
    final wingspanMetersValue = json['wingspanMeters'];
    final rangeKmValue = json['rangeKm'];
    final maxAltitudeMetersValue = json['maxAltitudeMeters'];
    final flightEnduranceValue = json['flightEndurance'];

    if (maxSpeedKmhValue is! int) {
      throw const FormatException('specs.maxSpeedKmh must be an integer');
    }
    if (wingspanMetersValue is! num) {
      throw const FormatException('specs.wingspanMeters must be a number');
    }
    if (rangeKmValue is! int) {
      throw const FormatException('specs.rangeKm must be an integer');
    }
    if (maxAltitudeMetersValue != null && maxAltitudeMetersValue is! int) {
      throw const FormatException('specs.maxAltitudeMeters must be an integer');
    }
    if (flightEnduranceValue != null && flightEnduranceValue is! int) {
      throw const FormatException('specs.flightEndurance must be an integer');
    }

    return AircraftSpecs(
        maxSpeedKmh: maxSpeedKmhValue,
        wingspanMeters: wingspanMetersValue.toDouble(),
        rangeKm: rangeKmValue,
        maxAltitudeMeters: maxAltitudeMetersValue as int?,
        flightEndurance:
            Duration(microseconds: flightEnduranceValue as int? ?? 0));
  }

  Map<String, dynamic> toJson() => {
        'maxSpeedKmh': maxSpeedKmh,
        'wingspanMeters': wingspanMeters,
        'rangeKm': rangeKm,
        'maxAltitudeMeters': maxAltitudeMeters,
        'flightEndurance': flightEndurance.inMicroseconds,
      };
}

class ConflictHistory {
  final String name;
  final int startYear;
  final int endYear;

  ConflictHistory({
    required this.name,
    required this.startYear,
    required this.endYear,
  });

  factory ConflictHistory.fromJson(Map<String, dynamic> json) {
    return ConflictHistory(
      name: json['name'] as String,
      startYear: json['startYear'] as int,
      endYear: json['endYear'] as int,
    );
  }

  Map<String, dynamic> toJson() => {
        'name': name,
        'startYear': startYear,
        'endYear': endYear,
      };
}

class Aircraft {
  final String id;
  final String model;
  final String manufacturer;
  final String? serialNumber;
  final int yearOfManufacture;
  final Decimal priceMillions;
  final double emptyWeightKg;
  final AircraftStatus status;
  final AircraftRole role;
  final List<String> tags;
  final DateTime firstFlightDate;
  final DateTime lastMaintenanceTime;
  final GeoLocation baseLocation;
  final AircraftSpecs specs;
  final List<ConflictHistory> conflicts;
  final Map<String, String> metadata;
  final int? estimatedUnitsProduced;
  final int? estimatedActiveUnits;
  final String? photoUrl;
  final List<int>? manualArchive;

  Aircraft({
    required this.id,
    required this.model,
    required this.manufacturer,
    this.serialNumber,
    required this.yearOfManufacture,
    required this.priceMillions,
    required this.emptyWeightKg,
    required this.status,
    required this.role,
    required List<String> tags,
    required this.firstFlightDate,
    required this.lastMaintenanceTime,
    required this.baseLocation,
    required this.specs,
    required List<ConflictHistory> conflicts,
    required this.metadata,
    this.estimatedUnitsProduced,
    this.estimatedActiveUnits,
    this.photoUrl,
    List<int>? manualArchive,
  })  : tags = List.unmodifiable(tags),
        conflicts = List.unmodifiable(conflicts),
        manualArchive =
            manualArchive != null ? List.unmodifiable(manualArchive) : null;

  factory Aircraft.fromJson(Map<String, dynamic> json) {
    return Aircraft(
      id: json['id'] as String,
      model: json['model'] as String,
      manufacturer: json['manufacturer'] as String,
      serialNumber: json['serialNumber'] as String?,
      yearOfManufacture: json['yearOfManufacture'] as int,
      priceMillions: Decimal.parse(json['priceMillions'].toString()),
      emptyWeightKg: (json['emptyWeightKg'] as num).toDouble(),
      status: AircraftStatus.values.asNameMap()[json['status'] as String] ??
          (throw ArgumentError.value(
            json['status'],
            'status',
            'Must be one of: ${AircraftStatus.values.map((e) => e.name).join(', ')}',
          )),
      role: AircraftRole.values.asNameMap()[json['role'] as String] ??
          (throw ArgumentError.value(
            json['role'],
            'role',
            'Must be one of: ${AircraftRole.values.map((e) => e.name).join(', ')}',
          )),
      tags: List<String>.from(json['tags'] as List),
      firstFlightDate: DateTime.parse(json['firstFlightDate'] as String),
      lastMaintenanceTime:
          DateTime.parse(json['lastMaintenanceTime'] as String),
      baseLocation:
          GeoLocation.fromJson(json['baseLocation'] as Map<String, dynamic>),
      specs: AircraftSpecs.fromJson(json['specs'] as Map<String, dynamic>),
      conflicts: (json['conflicts'] as List)
          .map((e) => ConflictHistory.fromJson(e as Map<String, dynamic>))
          .toList(),
      metadata: Map<String, String>.from(json['metadata'] as Map),
      estimatedUnitsProduced: json['estimatedUnitsProduced'] as int?,
      estimatedActiveUnits: json['estimatedActiveUnits'] as int?,
      photoUrl: json['photoUrl'] as String?,
      manualArchive: json['manualArchive'] != null
          ? List<int>.from(json['manualArchive'] as List)
          : null,
    );
  }

  Map<String, dynamic> toJson() => {
        'id': id,
        'model': model,
        'manufacturer': manufacturer,
        'serialNumber': serialNumber,
        'yearOfManufacture': yearOfManufacture,
        'priceMillions': priceMillions.toString(),
        'emptyWeightKg': emptyWeightKg,
        'status': status.name,
        'role': role.name,
        'tags': tags,
        'firstFlightDate': firstFlightDate.toIso8601String(),
        'lastMaintenanceTime': lastMaintenanceTime.toIso8601String(),
        'baseLocation': baseLocation.toJson(),
        'specs': specs.toJson(),
        'conflicts': conflicts.map((c) => c.toJson()).toList(),
        'metadata': metadata,
        'estimatedUnitsProduced': estimatedUnitsProduced,
        'estimatedActiveUnits': estimatedActiveUnits,
        'photoUrl': photoUrl,
        'manualArchive': manualArchive,
      };
}

class CreateAircraftRequest {
  final String model;
  final String manufacturer;
  final String? serialNumber;
  final int yearOfManufacture;
  final String priceMillions;
  final double emptyWeightKg;
  final String status;
  final String role;
  final List<String> tags;
  final String firstFlightDate;
  final String lastMaintenanceTime;
  final Map<String, dynamic> baseLocation;
  final Map<String, dynamic> specs;
  final List<Map<String, dynamic>> conflicts;
  final Map<String, String> metadata;
  final int? estimatedUnitsProduced;
  final int? estimatedActiveUnits;
  final String? photoUrl;
  final List<int>? manualArchive;

  CreateAircraftRequest({
    required this.model,
    required this.manufacturer,
    this.serialNumber,
    required this.yearOfManufacture,
    required this.priceMillions,
    required this.emptyWeightKg,
    required this.status,
    required this.role,
    required this.tags,
    required this.firstFlightDate,
    required this.lastMaintenanceTime,
    required this.baseLocation,
    required this.specs,
    required this.conflicts,
    required this.metadata,
    this.estimatedUnitsProduced,
    this.estimatedActiveUnits,
    this.photoUrl,
    this.manualArchive,
  });

  factory CreateAircraftRequest.fromJson(Map<String, dynamic> json) {
    final model = json['model'];
    final manufacturer = json['manufacturer'];
    final yearOfManufacture = json['yearOfManufacture'];
    final emptyWeightKg = json['emptyWeightKg'];
    final status = json['status'];
    final role = json['role'];
    final firstFlightDate = json['firstFlightDate'];
    final lastMaintenanceTime = json['lastMaintenanceTime'];
    final baseLocation = json['baseLocation'];
    final specs = json['specs'];

    if (model is! String || model.trim().isEmpty) {
      throw const FormatException('model is required and must be a string');
    }
    if (manufacturer is! String || manufacturer.trim().isEmpty) {
      throw const FormatException(
        'manufacturer is required and must be a string',
      );
    }
    if (yearOfManufacture is! int) {
      throw const FormatException('yearOfManufacture must be an integer');
    }
    if (emptyWeightKg is! num) {
      throw const FormatException('emptyWeightKg must be a number');
    }
    if (status is! String) {
      throw const FormatException('status must be a string');
    }
    if (role is! String) {
      throw const FormatException('role must be a string');
    }
    if (firstFlightDate is! String) {
      throw const FormatException('firstFlightDate must be a string');
    }
    if (lastMaintenanceTime is! String) {
      throw const FormatException('lastMaintenanceTime must be a string');
    }
    if (baseLocation is! Map<String, dynamic>) {
      throw const FormatException('baseLocation must be an object');
    }
    if (specs is! Map<String, dynamic>) {
      throw const FormatException('specs must be an object');
    }

    final tagsValue = json['tags'];
    final conflictsValue = json['conflicts'];
    final metadataValue = json['metadata'];
    final manualArchiveValue = json['manualArchive'];

    final tags = tagsValue == null
        ? <String>[]
        : List<String>.from(tagsValue as List<dynamic>);
    final conflicts = conflictsValue == null
        ? <Map<String, dynamic>>[]
        : List<Map<String, dynamic>>.from(conflictsValue as List<dynamic>);
    final metadata = metadataValue == null
        ? <String, String>{}
        : Map<String, String>.from(metadataValue as Map<dynamic, dynamic>);
    final manualArchive = manualArchiveValue == null
        ? null
        : List<int>.from(manualArchiveValue as List<dynamic>);

    return CreateAircraftRequest(
      model: model,
      manufacturer: manufacturer,
      serialNumber: json['serialNumber'] as String?,
      yearOfManufacture: yearOfManufacture,
      priceMillions: json['priceMillions'].toString(),
      emptyWeightKg: emptyWeightKg.toDouble(),
      status: status,
      role: role,
      tags: tags,
      firstFlightDate: firstFlightDate,
      lastMaintenanceTime: lastMaintenanceTime,
      baseLocation: baseLocation,
      specs: specs,
      conflicts: conflicts,
      metadata: metadata,
      estimatedUnitsProduced: json['estimatedUnitsProduced'] as int?,
      estimatedActiveUnits: json['estimatedActiveUnits'] as int?,
      photoUrl: json['photoUrl'] as String?,
      manualArchive: manualArchive,
    );
  }
}
