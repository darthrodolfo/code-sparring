export enum AircraftRole {
  Fighter = "Fighter",
  Bomber = "Bomber",
  Transport = "Transport",
  Reconnaissance = "Reconnaissance",
  Drone = "Drone",
  Trainer = "Trainer",
}

export enum AircraftStatus {
  Active = "Active",
  Retired = "Retired",
  Maintenance = "Maintenance",
  Stored = "Stored",
}

export interface Geolocation {
  latitude: number;
  longitude: number;
}
export interface AircraftSpecs {
  maxSpeedKmh: number; // in km/h
  wingspanMeters: number; // in km
  rangeKm: number; // in meters
  maxAltitudeMeters: number | null;
  flightEndurance: string; // ISO 8601 duration ex: "PT14H30M"
}

export interface ConflictHistory {
  name: string;
  startYear: number; // ISO 8601 date string
  endYear: number; // ISO 8601 date string
  roleInConflict: AircraftRole;
}

export interface AircraftV2 {
  id: string;
  model: string;
  manufacturer: string;
  serialNumber: string | null;
  yearOfManufacture: number;
  priceMillionUSD: string; // decimal serializado como string
  emptyWeightKg: number;
  status: AircraftStatus;
  firstFlightDate: string; // ISO 8601 date: "YYYY-MM-DD"
  lastMaintenanceTime: string; // ISO 8601 datetime: "2024-01-01T00:00:00Z"
  baseLocation: Geolocation;
  specs: AircraftSpecs;
  role: AircraftRole;
  tags: string[];
  conflictHistory: ConflictHistory[];
  metadata: Record<string, string>;
  estimatedUnitsProduced: number | null;
  estimatedActivateUnits: number | null;
  photoUrl: string | null;
  manualArchive: string | null; //base64 encoded document
}
