import 'package:backend_dart/handlers/aircraft_handlers.dart';
import 'package:dart_frog/dart_frog.dart';

Future<Response> onRequest(RequestContext context) async {
  return onAircraftCollection(context, basePath: '/aircraft');
}
