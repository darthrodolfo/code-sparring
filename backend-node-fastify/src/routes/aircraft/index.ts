import { FastifyPluginAsync } from "fastify";
import { v4 as uuidv4 } from "uuid";
import { AircraftV2 } from "./aircraft.types.js";
import { aircraftStore } from "./aircraft.store.js";

const aircraft: FastifyPluginAsync = async (fastify): Promise<void> => {
  fastify.get("/decolamos", async () => {
    return { message: "Decolamos Fastify!" };
  });

  fastify.get("/", async () => {
    return Array.from(aircraftStore.values());
  });

  fastify.post<{ Body: Omit<AircraftV2, "id"> }>(
    "/",
    async (request, reply) => {
      const aircraft: AircraftV2 = {
        id: uuidv4(),
        ...request.body,
      };
      aircraftStore.set(aircraft.id, aircraft);
      return reply.code(201).send(aircraft);
    },
  );

  fastify.get<{ Params: { id: string } }>("/:id", async (request, reply) => {
    const aircraft = aircraftStore.get(request.params.id);
    if (!aircraft)
      return reply.code(404).send({ message: "Aircraft not found" });
    return aircraft;
  });

  fastify.put<{ Params: { id: string }; Body: Omit<AircraftV2, "id"> }>(
    "/:id",
    async (request, reply) => {
      if (!aircraftStore.has(request.params.id))
        return reply.code(404).send({ message: "Aircraft not found" });
      const aircraftUpdated: AircraftV2 = {
        id: request.params.id,
        ...request.body,
      };
      aircraftStore.set(request.params.id, aircraftUpdated);
      return aircraftUpdated;
    },
  );

  fastify.delete<{ Params: { id: string } }>("/:id", async (request, reply) => {
    if (!aircraftStore.has(request.params.id))
      return reply.code(404).send({ message: "Aircraft not found" });
    aircraftStore.delete(request.params.id);
    return reply.code(200).send({ message: "Aircraft deleted" });
  });
};
export default aircraft;
