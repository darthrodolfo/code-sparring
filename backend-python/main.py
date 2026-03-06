from fastapi import FastAPI, Response
from uuid import UUID, uuid4
from pydantic import BaseModel, Field
from datetime import datetime, date, timedelta
from enum import Enum
from typing import Optional

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
    # max_altitude: Optional[int] = None
    max_altitude: int | None = None

class Aircraft(BaseModel):
    id: UUID
    model: str
    manufacturer: str
    year: int
    
aircraft_store: dict[UUID, Aircraft] = {}

@app.get("/aircraft")
async def list_aircraft() -> list[Aircraft]:
    return list(aircraft_store.values())

@app.get("/decolamos")
def health():
    return "Decolamos"

@app.post("/aircraft", status_code=201)
def create_aircraft(request: CreateAircraftRequest, response: Response) -> Aircraft:
    aircraft = Aircraft(
        id = uuid4(),
        model = request.model.strip(),
        manufacturer = request.manufacturer.strip(),
        year = request.year,
    )
    aircraft_store[aircraft.id] = aircraft
    response.headers["Location"] = f"/aircraft/{aircraft.id}"

    return aircraft


    