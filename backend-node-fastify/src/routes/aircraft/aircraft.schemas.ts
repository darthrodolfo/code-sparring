import { AircraftRole, AircraftStatus } from './aircraft.types'

const aircraftRoleValues = Object.values(AircraftRole);
const aircraftStatusValues = Object.values(AircraftStatus);

const geolocationSchema = {
  type: 'object',
  additionalProperties: false,
  required: ['latitude', 'longitude'],
  properties: {
    latitude: { type: 'number', minimum: -90, maximum: 90 },
    longitude: { type: 'number', minimum: -180, maximum: 180 }
  }
} as const

const aircraftSpecsSchema = {
  type: 'object',
  additionalProperties: false,
  required: [
    'maxSpeedKmh',
    'wingspanMeters',
    'rangeKm',
    'maxAltitudeMeters',
    'flightEndurance'
  ],
  properties: {
    maxSpeedKmh: { type: 'number', exclusiveMinimum: 0 },
    wingspanMeters: { type: 'number', exclusiveMinimum: 0 },
    rangeKm: { type: 'number', exclusiveMinimum: 0 },
    maxAltitudeMeters: {
      type: ['number', 'null'],
      exclusiveMinimum: 0
    },
    flightEndurance: {
      type: 'string',
      minLength: 2,
      maxLength: 50,
      pattern: '^P.+'
    }
  }
} as const

const conflictHistoryItemSchema = {
  type: 'object',
  additionalProperties: false,
  required: ['name', 'startYear', 'endYear', 'roleInConflict'],
  properties: {
    name: { type: 'string', minLength: 1, maxLength: 100 },
    startYear: { type: 'integer', minimum: 1900, maximum: 2030 },
    endYear: { type: 'integer', minimum: 1900, maximum: 2030 },
    roleInConflict: {
      type: 'string',
      enum: aircraftRoleValues
    }
  }
} as const

export const aircraftIdParamsSchema = {
  type: 'object',
  additionalProperties: false,
  required: ['id'],
  properties: {
    id: {
      type: 'string',
      minLength: 1
    }
  }
} as const

export const aircraftBodySchema = {
  type: 'object',
  additionalProperties: false,
  required: [
    'model',
    'manufacturer',
    'serialNumber',
    'yearOfManufacture',
    'priceMillionUSD',
    'emptyWeightKg',
    'status',
    'role',
    'tags',
    'firstFlightDate',
    'lastMaintenanceTime',
    'baseLocation',
    'specs',
    'conflictHistory',
    'metadata',
    'estimatedUnitsProduced',
    'estimatedActiveUnits',
    'photoUrl',
    'manualArchive'
  ],
  properties: {
    model: { type: 'string', minLength: 2, maxLength: 100 },
    manufacturer: { type: 'string', minLength: 2, maxLength: 100 },
    serialNumber: {
      type: ['string', 'null'],
      minLength: 1,
      maxLength: 100,
      pattern: '^[A-Za-z0-9-]+$'
    },
    yearOfManufacture: {
      type: 'integer',
      minimum: 1900,
      maximum: 2030
    },
    priceMillionUSD: {
      type: 'string',
      pattern: '^[0-9]+(?:\\.[0-9]{1,2})?$'
    },
    emptyWeightKg: {
      type: 'number',
      exclusiveMinimum: 0
    },
    status: {
      type: 'string',
      enum: aircraftStatusValues
    },
    role: {
      type: 'string',
      enum: aircraftRoleValues
    },
    tags: {
      type: 'array',
      minItems: 1,
      maxItems: 5,
      uniqueItems: true,
      items: {
        type: 'string',
        minLength: 1,
        maxLength: 20
      }
    },
    firstFlightDate: {
      type: 'string',
      format: 'date'
    },
    lastMaintenanceTime: {
      type: 'string',
      format: 'date-time'
    },
    baseLocation: geolocationSchema,
    specs: aircraftSpecsSchema,
    conflictHistory: {
      type: 'array',
      items: conflictHistoryItemSchema,
      maxItems: 5
    },
    metadata: {
      type: 'object',
      maxProperties: 20,
      propertyNames: {
        type: 'string',
        maxLength: 50
      },
      additionalProperties: {
        type: 'string',
        maxLength: 500
      }
    },
    estimatedUnitsProduced: {
      type: ['integer', 'null'],
      minimum: 0
    },
    estimatedActiveUnits: {
      type: ['integer', 'null'],
      minimum: 0
    },
    photoUrl: {
      type: ['string', 'null'],
      format: 'uri',
      maxLength: 2048
    },
    manualArchive: {
      type: ['string', 'null'],
      maxLength: 14000000
    }
  }
} as const