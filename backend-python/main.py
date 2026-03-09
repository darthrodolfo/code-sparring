from fastapi import FastAPI, Response, HTTPException, Depends
from uuid import UUID, uuid4
from pydantic import BaseModel, Field
from datetime import datetime, date, timedelta
from enum import Enum
from decimal import Decimal
import aiosqlite
import json

DB_PATH = "d:/DevStuff/code-sparring/backend-csharp/aerostack.db"

app = FastAPI()

class GeoLocation(BaseModel):
    latitude: float
    longitude: float

class AircraftRole(str, Enum):
    Fighter = "Fighter"
    Bomber = "Bomber"
    Transport = "Transport"
    Trainer = "Trainer"
    Drone = "Drone"
    Reconnaissance = "Reconnaissance"

class AircraftStatus(str, Enum):
    Active = "Active"
    Maintenance = "Maintenance"
    Retired = "Retired"
    Stored = "Stored"

class AircraftSpecs(BaseModel):
    max_speed_kmh: int
    wingspan_meters: float
    range_km: int
    max_altitude_meters: int | None = None
    flight_endurance: timedelta

class ConflictHistory(BaseModel):
    name: str
    start_year: int
    end_year: int

class CreateAircraftRequest(BaseModel):
    model: str = Field(min_length=1, max_length=80)
    manufacturer: str = Field(min_length=1, max_length=80)
    year: int = Field(ge=1903, le=datetime.now().year + 1)

class CreateAircraftV2Request(BaseModel):
    model: str = Field(min_length=1, max_length=80)
    manufacturer: str = Field(min_length=1, max_length=80)
    serial_number: str | None = None
    year_of_manufacture: int = Field(ge=1903, le=datetime.now().year + 1)
    price_millions: Decimal = Field(gt=0)
    empty_weight_kg: float = Field(gt=0)
    status: AircraftStatus
    role: AircraftRole
    tags: list[str] = []
    first_flight_date: date
    last_maintenance_time: datetime
    base_location: GeoLocation
    specs: AircraftSpecs
    conflicts: list[ConflictHistory] = []
    metadata: dict[str,str] = {}
    estimated_units_produced: int | None = None
    estimated_active_units: int | None = None
    photo_url: str | None = None
    manual_archive: bytes | None = None

class Aircraft(BaseModel):
    id: UUID
    model: str
    manufacturer: str
    year: int

class AircraftV2(BaseModel):
    id: UUID
    model: str
    manufacturer: str
    serial_number: str | None = None
    year_of_manufacture: int
    price_millions: Decimal
    empty_weight_kg: float
    status: AircraftStatus
    role: AircraftRole
    tags: list[str]
    first_flight_date: date
    last_maintenance_time: datetime
    base_location: GeoLocation
    specs: AircraftSpecs
    conflicts: list[ConflictHistory]
    metadata: dict[str, str]
    estimated_units_produced: int | None = None
    estimated_active_units: int | None = None
    photo_url: str | None = None
    manual_archive: bytes | None = None
    
aircraft_store: dict[UUID, Aircraft] = {}
aircraft_v2_store: dict[UUID, AircraftV2] = {}

async def get_db():
    async with aiosqlite.connect(DB_PATH) as db:
        db.row_factory = aiosqlite.Row
        yield db

@app.get("/aircraft-v2")
async def list_aircraft_v2(db: aiosqlite.Connection = Depends(get_db)) -> list[AircraftV2]:

    # 1. Buscar todas as aeronaves
    async with db.execute("SELECT * FROM aircraft_v2") as cursor:
        rows = await cursor.fetchall()
    
    aircraft_list = []
    
    for row in rows:
        aircraft_id = row["id"]
        
        # 2. Buscar Tags para esta aeronave específica
        async with db.execute("SELECT tag FROM aircraft_tags WHERE aircraft_id = ?", (aircraft_id,)) as t_cursor:
            tags = [t[0] for t in await t_cursor.fetchall()]
        
        # 3. Buscar Conflitos para esta aeronave específica
        async with db.execute("SELECT name, start_year, end_year FROM aircraft_conflicts WHERE aircraft_id = ?", (aircraft_id,)) as c_cursor:
            conflicts = [dict(c) for c in await c_cursor.fetchall()]
        
        # 4. Montar o objeto AircraftV2
        aircraft_list.append(AircraftV2(
            id=UUID(row["id"]),
            model=row["model"],
            manufacturer=row["manufacturer"],
            serial_number=row["serial_number"],
            year_of_manufacture=row["year_of_manufacture"],
            price_millions=Decimal(row["price_millions"]),
            empty_weight_kg=row["empty_weight_kg"],
            status=row["status"],
            role=row["role"],
            tags=tags,
            first_flight_date=date.fromisoformat(row["first_flight_date"]),
            last_maintenance_time=datetime.fromisoformat(row["last_maintenance_time"]),
            base_location=GeoLocation(latitude=row["base_latitude"], longitude=row["base_longitude"]),
            specs=AircraftSpecs(
                max_speed_kmh=row["max_speed_kmh"],
                wingspan_meters=row["wingspan_meters"],
                range_km=row["range_km"],
                max_altitude_meters=row["max_altitude_meters"],
                flight_endurance=row["flight_endurance"]
            ),
            conflicts=[ConflictHistory(**c) for c in conflicts],
            metadata=json.loads(row["metadata"]),
            estimated_units_produced=row["estimated_units_produced"],
            estimated_active_units=row["estimated_active_units"],
            photo_url=row["photo_url"],
            manual_archive=row["manual_archive"]
        ))
        
    return aircraft_list


@app.get("/aircraft-v2/{id}")
async def get_aircraft_v2(id: UUID, db: aiosqlite.Connection = Depends(get_db)):     
    # Query principal
    async with db.execute("SELECT * FROM aircraft_v2 WHERE id = ?", (str(id),)) as cursor:
        row = await cursor.fetchone()
        if not row:
            raise HTTPException(404, "Not found")

    # Buscar Tags
    async with db.execute("SELECT tag FROM aircraft_tags WHERE aircraft_id = ?", (str(id),)) as cursor:
        tags = [r[0] for r in await cursor.fetchall()]

    # Buscar Conflitos
    async with db.execute("SELECT name, start_year, end_year FROM aircraft_conflicts WHERE aircraft_id = ?", (str(id),)) as cursor:
        conflicts = [dict(r) for r in await cursor.fetchall()]

    return AircraftV2(
        id=UUID(row["id"]),
        model=row["model"],
        manufacturer=row["manufacturer"],
        serial_number=row["serial_number"],
        year_of_manufacture=row["year_of_manufacture"],
        price_millions=Decimal(row["price_millions"]),
        empty_weight_kg=row["empty_weight_kg"],
        status=row["status"],
        role=row["role"],
        tags=tags,
        first_flight_date=date.fromisoformat(row["first_flight_date"]),
        last_maintenance_time=datetime.fromisoformat(row["last_maintenance_time"]),
        base_location=GeoLocation(latitude=row["base_latitude"], longitude=row["base_longitude"]),
        specs=AircraftSpecs(
            max_speed_kmh=row["max_speed_kmh"],
            wingspan_meters=row["wingspan_meters"],
            range_km=row["range_km"],
            max_altitude_meters=row["max_altitude_meters"],
            flight_endurance=row["flight_endurance"]
        ),
        conflicts=[ConflictHistory(**c) for c in conflicts],
        metadata=json.loads(row["metadata"]),
        estimated_units_produced=row["estimated_units_produced"],
        estimated_active_units=row["estimated_active_units"],
        photo_url=row["photo_url"],
        manual_archive=row["manual_archive"]
    )


@app.get("/decolamos")
def health():
    return "Decolamos"

@app.post("/aircraft", status_code=201)
async def create_aircraft(request: CreateAircraftRequest, response: Response) -> Aircraft:
    aircraft = Aircraft(
        id = uuid4(),
        model = request.model.strip(),
        manufacturer = request.manufacturer.strip(),
        year = request.year,
    )
    aircraft_store[aircraft.id] = aircraft
    response.headers["Location"] = f"/aircraft/{aircraft.id}"

    return aircraft

@app.post("/aircraft-v2", status_code=201)
async def create_aircraft_v2(request: CreateAircraftV2Request, response: Response):
    new_id = uuid4()

    async with aiosqlite.connect(DB_PATH) as db:
        await db.execute("""
            INSERT INTO aircraft_v2 (
                id, model, manufacturer, serial_number, year_of_manufacture,
                price_millions, empty_weight_kg, status, role, 
                first_flight_date, last_maintenance_time,
                base_latitude, base_longitude,
                max_speed_kmh, wingspan_meters, range_km, max_altitude_meters, flight_endurance,
                metadata, estimated_units_produced, estimated_active_units, photo_url
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        """, (
            str(new_id), request.model.strip(), request.manufacturer.strip(), request.serial_number,
            request.year_of_manufacture, str(request.price_millions), request.empty_weight_kg,
            request.status.value, request.role.value,
            request.first_flight_date.isoformat(), request.last_maintenance_time.isoformat(),
            request.base_location.latitude, request.base_location.longitude,
            request.specs.max_speed_kmh, request.specs.wingspan_meters, request.specs.range_km,
            request.specs.max_altitude_meters, str(request.specs.flight_endurance),
            json.dumps(request.metadata), request.estimated_units_produced,
            request.estimated_active_units, request.photo_url
        ))

        # 2. Inserir Tags (Loop)
        for tag in request.tags:
            await db.execute("INSERT INTO aircraft_tags (aircraft_id, tag) VALUES (?, ?)", (str(new_id), tag))

        # 3. Inserir Conflitos (Loop)
        for c in request.conflicts:
            await db.execute("""
                INSERT INTO aircraft_conflicts (aircraft_id, name, start_year, end_year)
                VALUES (?, ?, ?, ?)
            """, (str(new_id), c.name, c.start_year, c.end_year))
        await db.commit()
    
    response.headers["Location"] = f"/aircraft-v2/{new_id}"
    return {**request.model_dump(), "id": new_id}

@app.put("/aircraft-v2/{id}")
async def update_aircraft_v2(id: UUID, request: CreateAircraftV2Request):
    async with aiosqlite.connect(DB_PATH) as db:
        # 1. Verificar existência
        async with db.execute("SELECT id FROM aircraft_v2 WHERE id = ?", (str(id),)) as cursor:
            if not await cursor.fetchone():
                raise HTTPException(status_code=404, detail="Aircraft not found")

        # 2. Update dos campos principais
        await db.execute("""
            UPDATE aircraft_v2 SET 
                model = ?, manufacturer = ?, serial_number = ?, year_of_manufacture = ?,
                price_millions = ?, empty_weight_kg = ?, status = ?, role = ?, 
                first_flight_date = ?, last_maintenance_time = ?,
                base_latitude = ?, base_longitude = ?,
                max_speed_kmh = ?, wingspan_meters = ?, range_km = ?, max_altitude_meters = ?, flight_endurance = ?,
                metadata = ?, estimated_units_produced = ?, estimated_active_units = ?, photo_url = ?
            WHERE id = ?
        """, (
            request.model.strip(), request.manufacturer.strip(), request.serial_number,
            request.year_of_manufacture, str(request.price_millions), request.empty_weight_kg,
            request.status.value, request.role.value,
            request.first_flight_date.isoformat(), request.last_maintenance_time.isoformat(),
            request.base_location.latitude, request.base_location.longitude,
            request.specs.max_speed_kmh, request.specs.wingspan_meters, request.specs.range_km,
            request.specs.max_altitude_meters, str(request.specs.flight_endurance),
            json.dumps(request.metadata), request.estimated_units_produced,
            request.estimated_active_units, request.photo_url,
            str(id)
        ))

        # 3. Atualizar Tags (Deletar e Reinserir)
        await db.execute("DELETE FROM aircraft_tags WHERE aircraft_id = ?", (str(id),))
        for tag in request.tags:
            await db.execute("INSERT INTO aircraft_tags (aircraft_id, tag) VALUES (?, ?)", (str(id), tag))

        # 4. Atualizar Conflitos (Deletar e Reinserir)
        await db.execute("DELETE FROM aircraft_conflicts WHERE aircraft_id = ?", (str(id),))
        for c in request.conflicts:
            await db.execute("""
                INSERT INTO aircraft_conflicts (aircraft_id, name, start_year, end_year)
                VALUES (?, ?, ?, ?)
            """, (str(id), c.name, c.start_year, c.end_year))

        await db.commit()
    
    return {**request.model_dump(), "id": id}


@app.delete("/aircraft-v2/{id}", status_code=204)
async def delete_aircraft_v2(id: UUID):
    async with aiosqlite.connect(DB_PATH) as db:
        # Tenta deletar (As FKs devem cuidar do resto se houver CASCADE)
        cursor = await db.execute("DELETE FROM aircraft_v2 WHERE id = ?", (str(id),))
        await db.commit()
        
        if cursor.rowcount == 0:
            raise HTTPException(status_code=404, detail="Aircraft not found")
            
    return Response(status_code=204)



    