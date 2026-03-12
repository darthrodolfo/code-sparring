import fp from "fastify-plugin";
import Database from "better-sqlite3";
import { join } from "node:path";

declare module "fastify" {
  interface FastifyInstance {
    db: Database.Database;
  }
}

export default fp(async (fastify) => {
  const db = new Database(join(process.cwd(), "aircraft.db"));

  db.exec(`
    CREATE TABLE IF NOT EXISTS aircraft (
      id                      TEXT PRIMARY KEY,
      model                   TEXT NOT NULL,
      manufacturer            TEXT NOT NULL,
      serial_number           TEXT,
      year_of_manufacture     INTEGER NOT NULL,
      price_million_usd       TEXT NOT NULL,
      empty_weight_kg         REAL NOT NULL,
      status                  TEXT NOT NULL,
      role                    TEXT NOT NULL,
      tags                    TEXT NOT NULL DEFAULT '[]',
      first_flight_date       TEXT NOT NULL,
      last_maintenance_time   TEXT NOT NULL,
      base_location           TEXT NOT NULL,
      specs                   TEXT NOT NULL,
      conflict_history        TEXT NOT NULL DEFAULT '[]',
      metadata                TEXT NOT NULL DEFAULT '{}',
      estimated_units_produced INTEGER,
      estimated_active_units  INTEGER,
      photo_url               TEXT,
      manual_archive          TEXT
    )
  `);

  fastify.decorate("db", db);

  fastify.addHook("onClose", async () => {
    db.close();
  });
});
