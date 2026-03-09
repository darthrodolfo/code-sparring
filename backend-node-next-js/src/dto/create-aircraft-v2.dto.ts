import {
  IsString,
  IsInt,
  IsEnum,
  IsBoolean,
  IsOptional,
  IsNumber,
  IsArray,
  ValidateNested,
  MaxLength,
  Min,
  Max,
  ArrayMaxSize,
} from 'class-validator';
import { Type } from 'class-transformer';
import { AircraftStatus } from '../enums/aircraft-status.enum';
import { AircraftCategory } from '../enums/aircraft-category.enum';

export class ConflictHistoryDto {
  @IsString()
  conflictName!: string;

  @IsInt()
  @Min(1900)
  startYear!: number;

  @IsOptional()
  @IsInt()
  endYear?: number;

  @IsString()
  role!: string;
}

export class CreateAircraftV2Request {
  @IsString()
  @MaxLength(80)
  model!: string;

  @IsString()
  @MaxLength(80)
  manufacturer!: string;

  @IsInt()
  @Min(1903)
  @Max(new Date().getFullYear() + 1)
  year!: number;

  @IsEnum(AircraftStatus)
  status!: AircraftStatus;

  @IsEnum(AircraftCategory)
  category!: AircraftCategory;

  @IsOptional()
  @IsNumber()
  maxSpeedKph?: number;

  @IsOptional()
  @IsNumber()
  ceilingMeters?: number;

  @IsOptional()
  @IsNumber()
  rangeKm?: number;

  @IsInt()
  @Min(1)
  @Max(10)
  engineCount!: number;

  @IsOptional()
  @IsString()
  engineModel?: string;

  @IsOptional()
  @IsNumber()
  wingspanMeters?: number;

  @IsOptional()
  @IsNumber()
  emptyWeightKg?: number;

  @IsOptional()
  @IsNumber()
  maxTakeoffWeightKg?: number;

  @IsString()
  @MaxLength(60)
  countryOfOrigin!: string;

  @IsOptional()
  @IsString()
  description?: string;

  @IsOptional()
  @IsString()
  firstFlightDate?: string;

  @IsBoolean()
  isStealthCapable!: boolean;

  @IsArray()
  @IsString({ each: true })
  @ArrayMaxSize(20)
  tags!: string[];

  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => ConflictHistoryDto)
  conflictHistory!: ConflictHistoryDto[];
}
