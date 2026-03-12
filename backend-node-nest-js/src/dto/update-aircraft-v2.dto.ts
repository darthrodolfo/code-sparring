import { PartialType } from '@nestjs/mapped-types';
import { CreateAircraftV2Request } from './create-aircraft-v2.dto';

export class UpdateAircraftV2Request extends PartialType(
  CreateAircraftV2Request,
) {}
