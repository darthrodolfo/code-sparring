import { Inject, Injectable } from '@nestjs/common';
import Database from 'better-sqlite3';
import { DB_TOKEN } from '../database/database.module';
import { AircraftV2 } from '../entities/aircraft-v2.entity';

@Injectable()
export class AircraftRepository {
  constructor(@Inject(DB_TOKEN) private readonly db: Database.Database) {}

  initSchema(): void {
    this.db.exec(`
      CREATE TABLE IF NOT EXISTS aircraft_v2 (
        id TEXT PRIMARY KEY,
        model TEXT NOT NULL,
        manufacturer TEXT NOT NULL,
        year INTEGER NOT NULL,
        status TEXT NOT NULL,
        category TEXT NOT NULL,
        max_speed_kph REAL,
        ceiling_meters REAL,
        range_km REAL,
        engine_count INTEGER NOT NULL,
        engine_model TEXT,
        wingspan_meters REAL,
        empty_weight_kg REAL,
        max_takeoff_weight_kg REAL,
        country_of_origin TEXT NOT NULL,
        description TEXT,
        first_flight_date TEXT,
        is_stealth_capable INTEGER NOT NULL DEFAULT 0
      );

      CREATE TABLE IF NOT EXISTS aircraft_tags (
        aircraft_id TEXT NOT NULL REFERENCES aircraft_v2(id) ON DELETE CASCADE,
        tag TEXT NOT NULL
      );

      CREATE TABLE IF NOT EXISTS aircraft_conflicts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        aircraft_id TEXT NOT NULL REFERENCES aircraft_v2(id) ON DELETE CASCADE,
        conflict_name TEXT NOT NULL,
        start_year INTEGER NOT NULL,
        end_year INTEGER,
        role TEXT NOT NULL
      );
    `);
  }

  findAll(): AircraftV2[] {
    const rows = this.db.prepare('SELECT * FROM aircraft_v2').all() as any[];
    return rows.map((row) => this.hydrate(row));
  }

  findById(id: string): AircraftV2 | undefined {
    const row = this.db
      .prepare('SELECT * FROM aircraft_v2 WHERE id = ?')
      .get(id) as any;
    if (!row) return undefined;
    return this.hydrate(row);
  }

  create(aircraft: AircraftV2): AircraftV2 {
    const insert = this.db.transaction((a: AircraftV2) => {
      this.db
        .prepare(
          `
        INSERT INTO aircraft_v2 (
          id, model, manufacturer, year, status, category,
          max_speed_kph, ceiling_meters, range_km, engine_count, engine_model,
          wingspan_meters, empty_weight_kg, max_takeoff_weight_kg,
          country_of_origin, description, first_flight_date, is_stealth_capable
        ) VALUES (
          @id, @model, @manufacturer, @year, @status, @category,
          @maxSpeedKpm, @ceilingMeters, @rangeKm, @engineCount, @engineModel,
          @wingspanMeters, @emptyWeightKg, @maxTakeoffWeightKg,
          @countryOfOrigin, @description, @firstFlightDate, @isStealthCapable
        )
      `,
        )
        .run({ ...a, isStealthCapable: a.isStealthCapable ? 1 : 0 });

      for (const tag of a.tags) {
        this.db
          .prepare('INSERT INTO aircraft_tags (aircraft_id, tag) VALUES (?, ?)')
          .run(a.id, tag);
      }

      for (const c of a.conflictHistory) {
        this.db
          .prepare(
            `
          INSERT INTO aircraft_conflicts (aircraft_id, conflict_name, start_year, end_year, role)
          VALUES (?, ?, ?, ?, ?)
        `,
          )
          .run(a.id, c.conflictName, c.startYear, c.endYear ?? null, c.role);
      }
    });

    insert(aircraft);
    return this.findById(aircraft.id)!;
  }

  update(id: string, partial: Partial<AircraftV2>): AircraftV2 | undefined {
    const existing = this.findById(id);
    if (!existing) return undefined;
    const merged = { ...existing, ...partial };

    const updateTx = this.db.transaction(() => {
      this.db
        .prepare(
          `
        UPDATE aircraft_v2 SET
          model = @model, manufacturer = @manufacturer, year = @year,
          status = @status, category = @category, max_speed_kph = @maxSpeedKpm,
          ceiling_meters = @ceilingMeters, range_km = @rangeKm,
          engine_count = @engineCount, engine_model = @engineModel,
          wingspan_meters = @wingspanMeters, empty_weight_kg = @emptyWeightKg,
          max_takeoff_weight_kg = @maxTakeoffWeightKg, country_of_origin = @countryOfOrigin,
          description = @description, first_flight_date = @firstFlightDate,
          is_stealth_capable = @isStealthCapable
        WHERE id = @id
      `,
        )
        .run({ ...merged, isStealthCapable: merged.isStealthCapable ? 1 : 0 });

      this.db
        .prepare('DELETE FROM aircraft_tags WHERE aircraft_id = ?')
        .run(id);
      for (const tag of merged.tags) {
        this.db
          .prepare('INSERT INTO aircraft_tags (aircraft_id, tag) VALUES (?, ?)')
          .run(id, tag);
      }

      this.db
        .prepare('DELETE FROM aircraft_conflicts WHERE aircraft_id = ?')
        .run(id);
      for (const c of merged.conflictHistory) {
        this.db
          .prepare(
            `
          INSERT INTO aircraft_conflicts (aircraft_id, conflict_name, start_year, end_year, role)
          VALUES (?, ?, ?, ?, ?)
        `,
          )
          .run(id, c.conflictName, c.startYear, c.endYear ?? null, c.role);
      }
    });

    updateTx();
    return this.findById(id)!;
  }

  delete(id: string): boolean {
    const result = this.db
      .prepare('DELETE FROM aircraft_v2 WHERE id = ?')
      .run(id);
    return result.changes > 0;
  }

  private hydrate(row: any): AircraftV2 {
    const tags = this.db
      .prepare('SELECT tag FROM aircraft_tags WHERE aircraft_id = ?')
      .all(row.id) as { tag: string }[];

    const conflicts = this.db
      .prepare('SELECT * FROM aircraft_conflicts WHERE aircraft_id = ?')
      .all(row.id) as any[];

    return {
      id: row.id,
      model: row.model,
      manufacturer: row.manufacturer,
      year: row.year,
      status: row.status,
      category: row.category,
      maxSpeedKpm: row.max_speed_kph ?? undefined,
      ceilingMeters: row.ceiling_meters ?? undefined,
      rangeKm: row.range_km ?? undefined,
      engineCount: row.engine_count,
      engineModel: row.engine_model ?? undefined,
      wingspanMeters: row.wingspan_meters ?? undefined,
      emptyWeightKg: row.empty_weight_kg ?? undefined,
      maxTakeoffWeightKg: row.max_takeoff_weight_kg ?? undefined,
      countryOfOrigin: row.country_of_origin,
      description: row.description ?? undefined,
      firstFlightDate: row.first_flight_date ?? undefined,
      isStealthCapable: row.is_stealth_capable === 1,
      tags: tags.map((t) => t.tag),
      conflictHistory: conflicts.map((c) => ({
        conflictName: c.conflict_name,
        startYear: c.start_year,
        endYear: c.end_year ?? undefined,
        role: c.role,
      })),
    };
  }
}
