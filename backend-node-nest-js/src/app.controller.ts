import { Controller, Get } from '@nestjs/common';

@Controller()
export class AppController {
  constructor() {}

  @Get()
  getHello(): string {
    return 'Hello Brrrrrrrrrrrr!';
  }

  @Get('decolamos')
  decolamos(): string {
    return 'Decolamos!';
  }
}
