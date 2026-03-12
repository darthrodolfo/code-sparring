import { FastifyPluginAsync } from "fastify";

const aircraft: FastifyPluginAsync = async (fastify): Promise<void> => {
  fastify.get("/decolamos", async () => {
    return { message: "Decolamos Fastify!" };
  });
};

export default aircraft;
