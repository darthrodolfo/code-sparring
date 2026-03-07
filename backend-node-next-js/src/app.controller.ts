import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';

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
}
