import { test } from 'node:test'
import * as assert from 'node:assert'
import { build } from '../helper'

const payload = {
  model: 'F-22 Raptor',
  manufacturer: 'Lockheed Martin',
  serialNumber: 'LM-2005-001',
  yearOfManufacture: 2005,
  priceMillionUSD: '143.8',
  emptyWeightKg: 19700,
  status: 'Active',
  role: 'Fighter',
  tags: ['stealth', 'air-superiority'],
  firstFlightDate: '1997-09-07',
  lastMaintenanceTime: '2024-06-01T08:00:00Z',
  baseLocation: { latitude: 30.4, longitude: -86.5 },
  specs: {
    maxSpeedKmh: 1960,
    wingspanMeters: 13.56,
    rangeKm: 2960,
    maxAltitudeMeters: 19812,
    flightEndurance: 'PT2H30M'
  },
  conflictHistory: [
    { name: 'Operation Iraqi Freedom', startYear: 2003, endYear: 2011, roleInConflict: 'Fighter' }
  ],
  metadata: { classification: '5th-gen' },
  estimatedUnitsProduced: 195,
  estimatedActiveUnits: 183,
  photoUrl: null,
  manualArchive: null
}

test('GET /aircraft returns wrapped list contract', async (t) => {
  const app = await build(t)

  const res = await app.inject({
    method: 'GET',
    url: '/aircraft'
  })

  assert.equal(res.statusCode, 200)
  assert.deepStrictEqual(JSON.parse(res.payload), {
    data: [],
    meta: {
      traceId: res.headers['x-trace-id'],
      page: 1,
      pageSize: 0,
      totalCount: 0,
      totalPages: 1
    }
  })
})

test('POST /aircraft creates and wraps resource contract', async (t) => {
  const app = await build(t)

  const res = await app.inject({
    method: 'POST',
    url: '/aircraft',
    payload
  })

  const body = JSON.parse(res.payload)

  assert.equal(res.statusCode, 201)
  assert.equal(body.meta.traceId, res.headers['x-trace-id'])
  assert.equal(typeof body.data.id, 'string')
  assert.equal(body.data.model, payload.model)
  assert.equal(body.data.estimatedActiveUnits, payload.estimatedActiveUnits)
})

test('GET /aircraft/:id returns structured not found error', async (t) => {
  const app = await build(t)

  const res = await app.inject({
    method: 'GET',
    url: '/aircraft/missing-id'
  })

  assert.equal(res.statusCode, 404)
  assert.deepStrictEqual(JSON.parse(res.payload), {
    error: {
      code: 'NOT_FOUND',
      message: 'Aircraft not found',
      traceId: res.headers['x-trace-id']
    }
  })
})

test('DELETE /aircraft/:id returns 204 with no body', async (t) => {
  const app = await build(t)

  const created = await app.inject({
    method: 'POST',
    url: '/aircraft',
    payload
  })

  const createdBody = JSON.parse(created.payload)

  const res = await app.inject({
    method: 'DELETE',
    url: `/aircraft/${createdBody.data.id}`
  })

  assert.equal(res.statusCode, 204)
  assert.equal(res.payload, '')
})


test('POST /aircraft rejects invalid payload with VALIDATION_ERROR', async (t) => {
  const app = await build(t);

  const invalidPayload = {
    ...payload,
    model: '',
    tags: [],
    yearOfManufacture: 1800
  }

  const res = await app.inject({
    method: "POST",
    url: "/aircraft",
    payload: invalidPayload
  })

  const body = JSON.parse(res.payload)

  assert.equal(res.statusCode, 400)
  assert.equal(body.error.code, "VALIDATION_ERROR")
  assert.equal(body.error.traceId, res.headers["x-trace-id"])
  assert.equal(Array.isArray(body.error.details), true)
  assert.equal(body.error.details.length > 0, true)
})

test('POST /aircraft rejects future firstFlightDate', async (t) => {
  const app = await build(t)

  const invalidPayload = {
    ...payload,
    firstFlightDate: '2999-01-01'
  }

  const res = await app.inject({
    method: 'POST',
    url: '/aircraft',
    payload: invalidPayload
  })

  const body = JSON.parse(res.payload)

  assert.equal(res.statusCode, 400)
  assert.equal(body.error.code, 'VALIDATION_ERROR')
  assert.equal(body.error.details.some((detail: { field: string }) => detail.field === 'firstFlightDate'), true)
})

test('POST /aircraft rejects priceMillionUSD equal to zero', async (t) => {
  const app = await build(t)

  const invalidPayload = {
    ...payload,
    priceMillionUSD: '0'
  }

  const res = await app.inject({
    method: 'POST',
    url: '/aircraft',
    payload: invalidPayload
  })

  const body = JSON.parse(res.payload)

  assert.equal(res.statusCode, 400)
  assert.equal(body.error.code, 'VALIDATION_ERROR')
  assert.equal(body.error.details.some((detail: { field: string }) => detail.field === 'priceMillionUSD'), true)
})

test('POST /aircraft rejects conflictHistory with startYear greater than endYear', async (t) => {
  const app = await build(t)

  const invalidPayload = {
    ...payload,
    conflictHistory: [
      {
        name: 'Invalid Conflict',
        startYear: 2020,
        endYear: 2010,
        roleInConflict: 'Fighter'
      }
    ]
  }

  const res = await app.inject({
    method: 'POST',
    url: '/aircraft',
    payload: invalidPayload
  })

  const body = JSON.parse(res.payload)

  assert.equal(res.statusCode, 400)
  assert.equal(body.error.code, 'VALIDATION_ERROR')
  assert.equal(
    body.error.details.some((detail: { field: string }) => detail.field === 'conflictHistory.0.startYear'),
    true
  )
})

test('POST /aircraft rejects duplicate serialNumber', async (t) => {
  const app = await build(t)

  await app.inject({
    method: 'POST',
    url: '/aircraft',
    payload
  })

  const duplicatedPayload = {
    ...payload,
    model: 'F-22 Raptor Copy'
  }

  const res = await app.inject({
    method: 'POST',
    url: '/aircraft',
    payload: duplicatedPayload
  })

  const body = JSON.parse(res.payload)

  assert.equal(res.statusCode, 400)
  assert.equal(body.error.code, 'VALIDATION_ERROR')
  assert.equal(body.error.details.some((detail: { field: string }) => detail.field === 'serialNumber'), true)
})
