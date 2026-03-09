import { IsInt, IsString, Max, MaxLength, Min } from 'class-validator';

export class CreateAircraftDto {
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
}
