import { test } from 'node:test'
import * as assert from 'node:assert'
import { build } from '../helper'

test('default root route', async (t) => {
  const app = await build(t)

  const res = await app.inject({
    url: '/'
  })
  assert.equal(res.statusCode, 200)
  assert.equal(res.headers['x-trace-id'] !== undefined, true)
  assert.deepStrictEqual(JSON.parse(res.payload), {
    data: {
      service: 'backend-node-fastify',
      version: '1.0.0',
      status: 'ok'
    },
    meta: {
      traceId: res.headers['x-trace-id']
    }
  })
})
