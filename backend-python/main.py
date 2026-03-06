from fastapi import FastAPI, Response
from uuid import UUID, uuid4
from pydantic import BaseModel, Field
from datetime import datetime

app = FastAPI()

class CreateAircraftRequest(BaseModel):
    model: str = Field(min_length=1, max_length=80)
    manufacturer: str = Field(min_length=1, max_length=80)
    year: int = Field(ge=1903, le=datetime.now().year + 1)

class Aircraft(BaseModel):
    id: UUID
    model: str
    manufacturer: str
    year: int


from fastapi.responses import PlainTextResponse

aircraft_store: dict[UUID, Aircraft] = {}

@app.get("/aircraft")
def list_aircraft() -> list[Aircraft]:
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


    