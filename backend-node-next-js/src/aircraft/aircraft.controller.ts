import {
  Body,
  Controller,
  Delete,
  Get,
  HttpCode,
  NotFoundException,
  Param,
  Post,
  Put,
} from '@nestjs/common';
import { AircraftService } from './aircraft.service';
import { CreateAircraftV2Request } from '../dto/create-aircraft-v2.dto';
import { UpdateAircraftV2Request } from '../dto/update-aircraft-v2.dto';

@Controller('aircraft')
export class AircraftController {
  constructor(private readonly aircraftService: AircraftService) {}

  @Get()
  findAll() {
    return this.aircraftService.findAll();
  }

  @Get(':id')
  findById(@Param('id') id: string) {
    return this.aircraftService.findById(id);
  }

  @Post()
  create(@Body() dto: CreateAircraftV2Request) {
    return this.aircraftService.create(dto);
  }

  @Put(':id')
  update(@Param('id') id: string, @Body() dto: UpdateAircraftV2Request) {
    return this.aircraftService.update(id, dto);
  }

  @Delete(':id')
  @HttpCode(204)
  delete(@Param('id') id: string): void {
    this.aircraftService.delete(id);
  }
}
