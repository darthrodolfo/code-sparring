import 'dart:io';
import 'package:backend_dart/models/aircraft.dart';
import 'package:backend_dart/store/aircraft_store.dart';
import 'package:dart_frog/dart_frog.dart';
import 'package:decimal/decimal.dart';

Future<Response> onRequest(RequestContext context, String id) async {
  return switch (context.request.method) {
    HttpMethod.get => _getById(context, id),
    HttpMethod.put => await _update(context, id),
    HttpMethod.delete => _delete(context, id),
    _ => Response(statusCode: HttpStatus.methodNotAllowed),
  };
}

Response _getById(RequestContext context, String id) {
  final store = context.read<AircraftStore>();
  final aircraft = store.getById(id);
  if (aircraft == null) {
    return Response.json(
      statusCode: HttpStatus.notFound,
      body: {'error': 'Aircraft not found'},
    );
  }
  return Response.json(body: aircraft.toJson());
}

Future<Response> _update(RequestContext context, String id) async {
  final store = context.read<AircraftStore>();

  if (store.getById(id) == null) {
    return Response.json(
      statusCode: HttpStatus.notFound,
      body: {'error': 'Aircraft not found'},
    );
  }

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

  final updated = Aircraft(
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

  store.update(id, updated);
  return Response.json(body: updated.toJson());
}

Response _delete(RequestContext context, String id) {
  final store = context.read<AircraftStore>();
  if (!store.delete(id)) {
    return Response.json(
      statusCode: HttpStatus.notFound,
      body: {'error': 'Aircraft not found'},
    );
  }
  return Response(statusCode: HttpStatus.noContent);
}
