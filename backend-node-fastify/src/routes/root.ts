import { FastifyPluginAsync } from 'fastify'

const root: FastifyPluginAsync = async (fastify): Promise<void> => {
  fastify.get('/', async () => {
    return {
      service: 'backend-node-fastify',
      version: '1.0.0',
      status: 'ok'
    }
  })
}

export default root
