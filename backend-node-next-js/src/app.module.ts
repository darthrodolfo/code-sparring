import { Module } from '@nestjs/common';
import { AircraftModule } from './aircraft/aircraft.module';
import { DatabaseModule } from './database/database.module';

@Module({
  imports: [AircraftModule, DatabaseModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
