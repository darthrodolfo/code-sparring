import 'dart:io';
import 'package:backend_dart/models/aircraft.dart';
import 'package:backend_dart/store/aircraft_store.dart';
import 'package:dart_frog/dart_frog.dart';
import 'package:decimal/decimal.dart';
import 'package:uuid/uuid.dart';

const _uuid = Uuid();

Future<Response> onRequest(RequestContext context) async {
  return switch (context.request.method) {
    HttpMethod.get => _getAll(context),
    HttpMethod.post => await _create(context),
    _ => Response(statusCode: HttpStatus.methodNotAllowed),
  };
}

Response _getAll(RequestContext context) {
  final store = context.read<AircraftStore>();
  return Response.json(
    body: store.getAll().map((a) => a.toJson()).toList(),
  );
}

Future<Response> _create(RequestContext context) async {
  final store = context.read<AircraftStore>();

  final Map<String, dynamic> body;
  try {
    body = await context.request.json() as Map<String, dynamic>;
  } catch (_) {
    return Response.json(
      statusCode: HttpStatus.badRequest,
      body: {'error': 'Invalid JSON body'},
    );
  }

  final CreateAircraftRequest dto;
  try {
    dto = CreateAircraftRequest.fromJson(body);
  } catch (e) {
    return Response.json(
      statusCode: HttpStatus.badRequest,
      body: {'error': e.toString()},
    );
  }

  final aircraft = _fromDto(_uuid.v7(), dto);
  store.add(aircraft);

  return Response.json(
    statusCode: HttpStatus.created,
    headers: {'Location': '/aircraft/${aircraft.id}'},
    body: aircraft.toJson(),
  );
}

Aircraft _fromDto(String id, CreateAircraftRequest dto) {
  return Aircraft(
    id: id,
    model: dto.model,
    manufacturer: dto.manufacturer,
    serialNumber: dto.serialNumber,
    yearOfManufacture: dto.yearOfManufacture,
    priceMillions: Decimal.parse(dto.priceMillions),
    emptyWeightKg: dto.emptyWeightKg,
    status: AircraftStatus.values.asNameMap()[dto.status] ??
        (throw ArgumentError.value(dto.status, 'status', 'Invalid status')),
    role: AircraftRole.values.asNameMap()[dto.role] ??
        (throw ArgumentError.value(dto.role, 'role', 'Invalid role')),
    tags: dto.tags,
    firstFlightDate: DateTime.parse(dto.firstFlightDate),
    lastMaintenanceTime: DateTime.parse(dto.lastMaintenanceTime),
    baseLocation: GeoLocation.fromJson(dto.baseLocation),
    specs: AircraftSpecs.fromJson(dto.specs),
    conflicts: dto.conflicts.map(ConflictHistory.fromJson).toList(),
    metadata: dto.metadata,
    estimatedUnitsProduced: dto.estimatedUnitsProduced,
    estimatedActiveUnits: dto.estimatedActiveUnits,
    photoUrl: dto.photoUrl,
  );
}
