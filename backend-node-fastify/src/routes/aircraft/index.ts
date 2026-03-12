import { FastifyPluginAsync } from "fastify";
import { v4 as uuidv4 } from "uuid";
import { AircraftV2 } from "./aircraft.types.js";

// ── DB row → AircraftV2 ───────────────────────────────────
function rowToAircraft(row: Record<string, unknown>): AircraftV2 {
  return {
    id: row.id as string,
    model: row.model as string,
    manufacturer: row.manufacturer as string,
    serialNumber: row.serial_number as string | null,
    yearOfManufacture: row.year_of_manufacture as number,
    priceMillionUSD: row.price_million_usd as string,
    emptyWeightKg: row.empty_weight_kg as number,
    status: row.status as AircraftV2["status"],
    role: row.role as AircraftV2["role"],
    tags: JSON.parse(row.tags as string),
    firstFlightDate: row.first_flight_date as string,
    lastMaintenanceTime: row.last_maintenance_time as string,
    baseLocation: JSON.parse(row.base_location as string),
    specs: JSON.parse(row.specs as string),
    conflictHistory: JSON.parse(row.conflict_history as string),
    metadata: JSON.parse(row.metadata as string),
    estimatedUnitsProduced: row.estimated_units_produced as number | null,
    estimatedActiveUnits: row.estimated_active_units as number | null,
    photoUrl: row.photo_url as string | null,
    manualArchive: row.manual_archive as string | null,
  };
}

const aircraft: FastifyPluginAsync = async (fastify): Promise<void> => {
  fastify.get("/decolamos", async () => {
    return { message: "Decolamos Fastify!" };
  });

  fastify.get("/", async () => {
    const rows = fastify.db.prepare("SELECT * FROM aircraft").all();
    return rows.map((row) => rowToAircraft(row as Record<string, unknown>));
  });

  fastify.post<{ Body: Omit<AircraftV2, "id"> }>(
    "/",
    async (request, reply) => {
      const aircraftId = uuidv4();
      const b = request.body;
      fastify.db
        .prepare(
          `
      INSERT INTO aircraft VALUES (
        ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
      )
    `,
        )
        .run(
          aircraftId,
          b.model,
          b.manufacturer,
          b.serialNumber,
          b.yearOfManufacture,
          b.priceMillionUSD,
          b.emptyWeightKg,
          b.status,
          b.role,
          JSON.stringify(b.tags),
          b.firstFlightDate,
          b.lastMaintenanceTime,
          JSON.stringify(b.baseLocation),
          JSON.stringify(b.specs),
          JSON.stringify(b.conflictHistory),
          JSON.stringify(b.metadata),
          b.estimatedUnitsProduced,
          b.estimatedActiveUnits,
          b.photoUrl,
          b.manualArchive,
        );
      const created = fastify.db
        .prepare("SELECT * FROM aircraft WHERE id = ?")
        .get(aircraftId);
      return reply
        .code(201)
        .send(rowToAircraft(created as Record<string, unknown>));
    },
  );

  fastify.get<{ Params: { id: string } }>("/:id", async (request, reply) => {
    const row = fastify.db
      .prepare("SELECT * FROM aircraft WHERE id = ?")
      .get(request.params.id);
    if (!row) return reply.code(404).send({ message: "Aircraft not found" });
    return rowToAircraft(row as Record<string, unknown>);
  });

  fastify.put<{ Params: { id: string }; Body: Omit<AircraftV2, "id"> }>(
    "/:id",
    async (request, reply) => {
      const existing = fastify.db
        .prepare("SELECT id FROM aircraft WHERE id = ?")
        .get(request.params.id);
      if (!existing)
        return reply.code(404).send({ message: "Aircraft not found" });
      const b = request.body;
      fastify.db
        .prepare(
          `
      UPDATE aircraft SET
        model=?, manufacturer=?, serial_number=?, year_of_manufacture=?,
        price_million_usd=?, empty_weight_kg=?, status=?, role=?, tags=?,
        first_flight_date=?, last_maintenance_time=?, base_location=?, specs=?,
        conflict_history=?, metadata=?, estimated_units_produced=?,
        estimated_active_units=?, photo_url=?, manual_archive=?
      WHERE id=?
    `,
        )
        .run(
          b.model,
          b.manufacturer,
          b.serialNumber,
          b.yearOfManufacture,
          b.priceMillionUSD,
          b.emptyWeightKg,
          b.status,
          b.role,
          JSON.stringify(b.tags),
          b.firstFlightDate,
          b.lastMaintenanceTime,
          JSON.stringify(b.baseLocation),
          JSON.stringify(b.specs),
          JSON.stringify(b.conflictHistory),
          JSON.stringify(b.metadata),
          b.estimatedUnitsProduced,
          b.estimatedActiveUnits,
          b.photoUrl,
          b.manualArchive,
          request.params.id,
        );
      const updated = fastify.db
        .prepare("SELECT * FROM aircraft WHERE id = ?")
        .get(request.params.id);
      return rowToAircraft(updated as Record<string, unknown>);
    },
  );

  fastify.delete<{ Params: { id: string } }>("/:id", async (request, reply) => {
    const existing = fastify.db
      .prepare("SELECT id FROM aircraft WHERE id = ?")
      .get(request.params.id);
    if (!existing)
      return reply.code(404).send({ message: "Aircraft not found" });
    fastify.db
      .prepare("DELETE FROM aircraft WHERE id = ?")
      .run(request.params.id);
    return reply.code(200).send({ message: "Aircraft deleted" });
  });
};
export default aircraft;
