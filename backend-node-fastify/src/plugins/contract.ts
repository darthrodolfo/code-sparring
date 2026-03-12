import fp from 'fastify-plugin'
import { FastifyError } from 'fastify'

type ErrorDetail = {
  field?: string
  message: string
}

type SuccessMeta = {
  traceId: string
  page?: number
  pageSize?: number
  totalCount?: number
  totalPages?: number
}

type SuccessEnvelope<T> = {
  data: T
  meta: SuccessMeta
}

type ErrorEnvelope = {
  error: {
    code: string
    message: string
    details?: ErrorDetail[]
    traceId: string
  }
}

function isObject(value: unknown): value is Record<string, unknown> {
  return value !== null && typeof value === 'object' && !Array.isArray(value)
}

function isSuccessEnvelope(payload: unknown): payload is SuccessEnvelope<unknown> {
  return isObject(payload) && 'data' in payload && 'meta' in payload
}

function isErrorEnvelope(payload: unknown): payload is ErrorEnvelope {
  return isObject(payload) && 'error' in payload
}

function buildSuccessMeta(payload: unknown, traceId: string): SuccessMeta {
  if (Array.isArray(payload)) {
    return {
      traceId,
      page: 1,
      pageSize: payload.length,
      totalCount: payload.length,
      totalPages: 1
    }
  }

  return { traceId }
}

function resolveErrorCode(error: FastifyError): string {
  if (error.validation) {
    return 'VALIDATION_ERROR'
  }

  if (error.statusCode === 404) {
    return 'NOT_FOUND'
  }

  if (error.statusCode && error.statusCode >= 500) {
    return 'INTERNAL_SERVER_ERROR'
  }

  return 'BAD_REQUEST'
}

function normalizeError(error: unknown): FastifyError {
  return error as FastifyError
}

export default fp(async (fastify) => {
  fastify.addHook('onRequest', async (request, reply) => {
    reply.header('x-trace-id', request.id)
  })

  fastify.addHook('preSerialization', async (request, reply, payload) => {
    if (reply.statusCode === 204 || payload === undefined) {
      return payload
    }

    if (isSuccessEnvelope(payload) || isErrorEnvelope(payload)) {
      return payload
    }

    return {
      data: payload,
      meta: buildSuccessMeta(payload, request.id)
    }
  })

  fastify.setErrorHandler(async (error, request, reply) => {
    const normalizedError = normalizeError(error)
    const statusCode = normalizedError.statusCode && normalizedError.statusCode >= 400
      ? normalizedError.statusCode
      : 500
    const details = normalizedError.validation?.map((item: { instancePath: string, message?: string }) => ({
      field: item.instancePath ? item.instancePath.replace(/^\//, '').replace(/\//g, '.') : undefined,
      message: item.message ?? 'Invalid value'
    }))

    return reply.code(statusCode).send({
      error: {
        code: resolveErrorCode(normalizedError),
        message: normalizedError.message,
        details: details && details.length > 0 ? details : undefined,
        traceId: request.id
      }
    })
  })

  fastify.setNotFoundHandler(async (request, reply) => {
    return reply.code(404).send({
      error: {
        code: 'NOT_FOUND',
        message: `Route ${request.method} ${request.url} not found`,
        traceId: request.id
      }
    })
  })
})
