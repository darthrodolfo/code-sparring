import { AircraftStatus } from '../enums/aircraft-status.enum';
import { AircraftCategory } from '../enums/aircraft-category.enum';

export interface ConflictHistory {
  conflictName: string;
  startYear: number;
  endYear?: number;
  role: string;
}

export interface AircraftV2 {
  id: string;
  model: string;
  manufacturer: string;
  year: number;
  status: AircraftStatus;
  category: AircraftCategory;
  maxSpeedKpm?: number;
  ceilingMeters?: number;
  rangeKm?: number;
  engineCount: number;
  engineModel?: string;
  wingspanMeters?: number;
  emptyWeightKg?: number;
  maxTakeoffWeightKg?: number;
  countryOfOrigin: string;
  description?: string;
  firstFlightDate?: string;
  isStealthCapable: boolean;
  tags: string[];
  conflictHistory: ConflictHistory[];
}
