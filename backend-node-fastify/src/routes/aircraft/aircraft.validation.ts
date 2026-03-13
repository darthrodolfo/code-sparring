import { AircraftV2 } from "./aircraft.types";

type ValidationDetail = {
  field: string;
  message: string;
};

function isFutureData(dateValue: string): boolean {
  const parsed = new Date(dateValue)

  if (Number.isNaN(parsed.getTime())) {
    return false
  }

  return parsed.getTime() > Date.now()
}

function hasPositiveDecimal(value: string): boolean {
  const parsed = Number(value)
  return Number.isFinite(parsed) && parsed > 0
}

export function validateAircraftBusinessRules(
  input: Omit<AircraftV2, "id">
): ValidationDetail[] {

  const details: ValidationDetail[] = []

  if (hasPositiveDecimal(input.priceMillionUSD) == false) {
    details.push({
      field: "priceMillionUSD",
      message: "Price must be greater than zero"
    })
  }

  if (isFutureData(input.firstFlightDate)) {
    details.push({
      field: "firstFlightDate",
      message: "Last flight date cannot be in the future"
    })
  }

  if (isFutureData(input.lastMaintenanceTime)) {
    details.push({
      field: "lastMaintenanceTime",
      message: "Last maintenance time cannot be in the future"
    })
  }

  for (let index = 0; index < input.conflictHistory.length; index++) {
    const conflict = input.conflictHistory[index];

    if (conflict.startYear > conflict.endYear) {
      details.push({
        field: `conflictHistory.${index}.startYear`,
        message: "Start year cannot be greater than end year"
      })
    }
  }

  return details
}