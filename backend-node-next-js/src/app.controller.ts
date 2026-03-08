import { randomUUID } from 'crypto';
import {
  Body,
  Controller,
  Get,
  Post,
  Param,
  NotFoundException,
  Delete,
  Put,
} from '@nestjs/common';
import { AppService } from './app.service';
import { CreateAircraftDto } from './dto/create-aircraft.dto';
import type { AircraftV2 } from './entities/aircraft-v2.entity';
import { CreateAircraftV2Request } from './dto/create-aircraft-v2.dto';
import { UpdateAircraftV2Request } from './dto/update-aircraft-v2.dto';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  private aircraftList: AircraftV2[] = [];

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Get('decolamos')
  decolamos(): string {
    return 'Decolamos!';
  }

  @Get('aircraft')
  getAircraft(): any[] {
    return [];
  }

  @Get('aircraft/:id')
  getAircraftById(@Param('id') id: string): AircraftV2 {
    const aircraft = this.aircraftList.find((a) => a.id === id);
    if (!aircraft) {
      throw new NotFoundException(`Aircraft with ID ${id} not found`);
    }
    return aircraft;
  }

  @Delete('aircraft/:id')
  deleteAircraft(@Param('id') id: string): void {
    const index = this.aircraftList.findIndex((a) => a.id === id);
    if (index === -1) throw new NotFoundException(`Aircraft ${id} not found`);

    this.aircraftList.splice(index, 1);
  }

  @Post('aircraft')
  createAircraft(@Body() dto: CreateAircraftDto): object {
    return {
      id: randomUUID(),
      ...dto,
    };
  }

  @Put('aircraft/:id')
  updateAircraft(
    @Param('id') id: string,
    @Body() dto: UpdateAircraftV2Request,
  ): AircraftV2 {
    const index = this.aircraftList.findIndex((a) => a.id === id);
    if (index === -1) throw new NotFoundException(`Aircraft ${id} not found`);

    this.aircraftList[index] = { ...this.aircraftList[index], ...dto };

    return this.aircraftList[index];
  }
}
