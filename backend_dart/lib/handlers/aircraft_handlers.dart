import 'dart:io';

import 'package:backend_dart/models/aircraft.dart';
import 'package:backend_dart/repositories/aircraft_repository.dart';
import 'package:dart_frog/dart_frog.dart';
import 'package:decimal/decimal.dart';
import 'package:uuid/uuid.dart';

const _uuid = Uuid();

Future<Response> onAircraftCollection(
  RequestContext context, {
  required String basePath,
}) {
  switch (context.request.method) {
    case HttpMethod.get:
      return Future.value(_getAll(context));
    case HttpMethod.post:
      return _create(context, basePath: basePath);
    default:
      return Future.value(Response(statusCode: HttpStatus.methodNotAllowed));
  }
}

Future<Response> onAircraftItem(RequestContext context, String id) {
  switch (context.request.method) {
    case HttpMethod.get:
      return Future.value(_getById(context, id));
    case HttpMethod.put:
      return _update(context, id);
    case HttpMethod.delete:
      return Future.value(_delete(context, id));
    default:
      return Future.value(Response(statusCode: HttpStatus.methodNotAllowed));
  }
}

Response _getAll(RequestContext context) {
  final store = context.read<AircraftRepository>();
  return Response.json(
    body: store.getAll().map((a) => a.toJson()).toList(),
  );
}

Response _getById(RequestContext context, String id) {
  final store = context.read<AircraftRepository>();
  final aircraft = store.getById(id);
  if (aircraft == null) {
    return _jsonError(HttpStatus.notFound, 'Aircraft not found');
  }
  return Response.json(body: aircraft.toJson());
}

Future<Response> _create(
  RequestContext context, {
  required String basePath,
}) async {
  final store = context.read<AircraftRepository>();
  final dtoOrError = await _readDto(context);
  if (dtoOrError.error != null) {
    return _jsonError(HttpStatus.badRequest, dtoOrError.error!);
  }

  final aircraftOrError = _fromDto(_uuid.v7(), dtoOrError.dto!);
  if (aircraftOrError.error != null) {
    return _jsonError(HttpStatus.badRequest, aircraftOrError.error!);
  }

  final aircraft = aircraftOrError.aircraft!;
  store.add(aircraft);
  return Response.json(
    statusCode: HttpStatus.created,
    headers: {'Location': '$basePath/${aircraft.id}'},
    body: aircraft.toJson(),
  );
}

Future<Response> _update(RequestContext context, String id) async {
  final store = context.read<AircraftRepository>();
  if (store.getById(id) == null) {
    return _jsonError(HttpStatus.notFound, 'Aircraft not found');
  }

  final dtoOrError = await _readDto(context);
  if (dtoOrError.error != null) {
    return _jsonError(HttpStatus.badRequest, dtoOrError.error!);
  }

  final updatedOrError = _fromDto(id, dtoOrError.dto!);
  if (updatedOrError.error != null) {
    return _jsonError(HttpStatus.badRequest, updatedOrError.error!);
  }

  final updated = updatedOrError.aircraft!;
  store.update(id, updated);
  return Response.json(body: updated.toJson());
}

Response _delete(RequestContext context, String id) {
  final store = context.read<AircraftRepository>();
  if (!store.delete(id)) {
    return _jsonError(HttpStatus.notFound, 'Aircraft not found');
  }
  return Response(statusCode: HttpStatus.noContent);
}

Future<_DtoParseResult> _readDto(RequestContext context) async {
  final Object? parsedBody;
  try {
    parsedBody = await context.request.json();
  } catch (_) {
    return const _DtoParseResult(error: 'Invalid JSON body');
  }

  if (parsedBody is! Map<String, dynamic>) {
    return const _DtoParseResult(error: 'JSON body must be an object');
  }

  try {
    return _DtoParseResult(dto: CreateAircraftRequest.fromJson(parsedBody));
  } on FormatException catch (e) {
    return _DtoParseResult(error: e.message);
  } on ArgumentError catch (e) {
    return _DtoParseResult(error: e.message?.toString() ?? e.toString());
  } catch (_) {
    return const _DtoParseResult(error: 'Invalid request body shape');
  }
}

_AircraftMapResult _fromDto(String id, CreateAircraftRequest dto) {
  final status = AircraftStatus.values.asNameMap()[dto.status];
  if (status == null) {
    return const _AircraftMapResult(
      error: 'status must be one of: active, maintenance, retired, stored',
    );
  }

  final role = AircraftRole.values.asNameMap()[dto.role];
  if (role == null) {
    return const _AircraftMapResult(
      error:
          'role must be one of: fighter, bomber, transport, reconnaissance, trainer, drone',
    );
  }

  final priceMillions = Decimal.tryParse(dto.priceMillions);
  if (priceMillions == null) {
    return const _AircraftMapResult(error: 'priceMillions must be decimal');
  }

  DateTime firstFlightDate;
  DateTime lastMaintenanceTime;
  try {
    firstFlightDate = DateTime.parse(dto.firstFlightDate);
    lastMaintenanceTime = DateTime.parse(dto.lastMaintenanceTime);
  } catch (_) {
    return const _AircraftMapResult(
      error: 'firstFlightDate/lastMaintenanceTime must be ISO-8601 date-time',
    );
  }

  try {
    return _AircraftMapResult(
      aircraft: Aircraft(
        id: id,
        model: dto.model,
        manufacturer: dto.manufacturer,
        serialNumber: dto.serialNumber,
        yearOfManufacture: dto.yearOfManufacture,
        priceMillions: priceMillions,
        emptyWeightKg: dto.emptyWeightKg,
        status: status,
        role: role,
        tags: dto.tags,
        firstFlightDate: firstFlightDate,
        lastMaintenanceTime: lastMaintenanceTime,
        baseLocation: GeoLocation.fromJson(dto.baseLocation),
        specs: AircraftSpecs.fromJson(dto.specs),
        conflicts: dto.conflicts.map(ConflictHistory.fromJson).toList(),
        metadata: dto.metadata,
        estimatedUnitsProduced: dto.estimatedUnitsProduced,
        estimatedActiveUnits: dto.estimatedActiveUnits,
        photoUrl: dto.photoUrl,
        manualArchive: dto.manualArchive,
      ),
    );
  } on FormatException catch (e) {
    return _AircraftMapResult(error: e.message);
  } on ArgumentError catch (e) {
    return _AircraftMapResult(error: e.message?.toString() ?? e.toString());
  } catch (_) {
    return const _AircraftMapResult(error: 'Invalid nested object shape');
  }
}

Response _jsonError(int statusCode, String error) {
  return Response.json(
    statusCode: statusCode,
    body: {'error': error},
  );
}

class _DtoParseResult {
  const _DtoParseResult({this.dto, this.error});
  final CreateAircraftRequest? dto;
  final String? error;
}

class _AircraftMapResult {
  const _AircraftMapResult({this.aircraft, this.error});
  final Aircraft? aircraft;
  final String? error;
}
