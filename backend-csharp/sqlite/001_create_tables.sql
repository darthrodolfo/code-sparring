CREATE TABLE IF NOT EXISTS aircraft_v2 (
    id TEXT PRIMARY KEY,
    model TEXT NOT NULL,
    manufacturer TEXT NOT NULL,
    serial_number TEXT,
    year_of_manufacture INTEGER NOT NULL,
    price_millions TEXT NOT NULL,
    empty_weight_kg REAL NOT NULL,
    status TEXT NOT NULL,
    role TEXT NOT NULL,
    first_flight_date TEXT NOT NULL,
    last_maintenance_time TEXT NOT NULL,
    base_latitude REAL NOT NULL,
    base_longitude REAL NOT NULL,
    max_speed_kmh INTEGER NOT NULL,
    wingspan_meters REAL NOT NULL,
    range_km INTEGER NOT NULL,
    max_altitude_meters INTEGER,
    flight_endurance TEXT NOT NULL,
    metadata TEXT NOT NULL,
    estimated_units_produced INTEGER,
    estimated_active_units INTEGER,
    photo_url TEXT,
    manual_archive BLOB
);

CREATE TABLE IF NOT EXISTS aircraft_tags (
    aircraft_id TEXT NOT NULL,
    tag TEXT NOT NULL,
    PRIMARY KEY (aircraft_id, tag),
    FOREIGN KEY (aircraft_id) REFERENCES aircraft_v2(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS aircraft_conflicts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    aircraft_id TEXT NOT NULL,
    name TEXT NOT NULL,
    start_year INTEGER NOT NULL,
    end_year INTEGER NOT NULL,
    FOREIGN KEY (aircraft_id) REFERENCES aircraft_v2(id) ON DELETE CASCADE
);