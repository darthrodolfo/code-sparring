from fastapi import FastAPI, Response
from uuid import UUID, uuid4
from pydantic import BaseModel, Field
from datetime import datetime, date, timedelta
from enum import Enum
from decimal import Decimal

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

@app.get("/aircraft-v2")
async def list_aircraft_v2() -> list[AircraftV2]:
    return list(aircraft_v2_store.values())

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
async def create_aircraft_v2(request: CreateAircraftV2Request, response: Response) -> AircraftV2:
    aircraft = AircraftV2(
        id=uuid4(),
        model=request.model.strip(),
        manufacturer=request.manufacturer.strip(),
        serial_number=request.serial_number,
        year_of_manufacture=request.year_of_manufacture,
        price_millions=request.price_millions,
        empty_weight_kg=request.empty_weight_kg,
        status=request.status,
        role=request.role,
        tags=request.tags,
        first_flight_date=request.first_flight_date,
        last_maintenance_time=request.last_maintenance_time,
        base_location=request.base_location,
        specs=request.specs,
        conflicts=request.conflicts,
        metadata=request.metadata,
        estimated_units_produced=request.estimated_units_produced,
        estimated_active_units=request.estimated_active_units,
        photo_url=request.photo_url,
        manual_archive=request.manual_archive,
    )
    aircraft_v2_store[aircraft.id] = aircraft
    response.headers["Location"] = f"/aircraft-v2/{aircraft.id}"
    return aircraft


    