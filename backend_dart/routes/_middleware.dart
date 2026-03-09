import 'package:backend_dart/store/aircraft_store.dart';
import 'package:dart_frog/dart_frog.dart';

final _store = AircraftStore();

Handler middleware(Handler handler) {
  return handler.use(provider<AircraftStore>((_) => _store));
}
