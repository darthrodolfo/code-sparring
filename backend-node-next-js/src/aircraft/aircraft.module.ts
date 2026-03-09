import { Module } from '@nestjs/common';
import { AircraftController } from './aircraft.controller';
import { AircraftService } from './aircraft.service';
import { AircraftRepository } from './aircraft.repository';

@Module({
  controllers: [AircraftController],
  providers: [AircraftService, AircraftRepository],
})
export class AircraftModule {}
