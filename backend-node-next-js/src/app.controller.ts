import { randomUUID } from 'crypto';
import { Body, Controller, Get, Post } from '@nestjs/common';
import { AppService } from './app.service';
import { CreateAircraftDto } from './dto/create-aircraft.dto';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

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

  @Post('aircraft')
  createAircraft(@Body() dto: CreateAircraftDto): object {
    return {
      id: randomUUID(),
      ...dto,
    };
  }
}
