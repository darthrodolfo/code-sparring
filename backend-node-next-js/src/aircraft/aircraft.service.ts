import { Injectable, NotFoundException, OnModuleInit } from '@nestjs/common';
import { AircraftRepository } from './aircraft.repository';
import { AircraftV2 } from '../entities/aircraft-v2.entity';
import { CreateAircraftV2Request } from '../dto/create-aircraft-v2.dto';
import { UpdateAircraftV2Request } from '../dto/update-aircraft-v2.dto';
import { randomUUID } from 'crypto';

@Injectable()
export class AircraftService implements OnModuleInit {
  constructor(private readonly repo: AircraftRepository) {}

  onModuleInit(): void {
    this.repo.initSchema();
  }

  findAll(): AircraftV2[] {
    return this.repo.findAll();
  }

  findById(id: string): AircraftV2 {
    const aircraft = this.repo.findById(id);
    if (!aircraft) throw new NotFoundException(`Aircraft ${id} not found`);
    return aircraft;
  }

  create(dto: CreateAircraftV2Request): AircraftV2 {
    const aircraft: AircraftV2 = { id: randomUUID(), ...dto };
    return this.repo.create(aircraft);
  }

  update(id: string, dto: UpdateAircraftV2Request): AircraftV2 {
    const updated = this.repo.update(id, dto);
    if (!updated) throw new NotFoundException(`Aircraft ${id} not found`);
    return updated;
  }

  delete(id: string): void {
    const deleted = this.repo.delete(id);
    if (!deleted) throw new NotFoundException(`Aircraft ${id} not found`);
  }
}
