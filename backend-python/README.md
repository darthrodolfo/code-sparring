# backend-python — AeroStack Lab

Python backend for the AeroStack Lab polyglot training ground.
Same Aircraft domain, different stack.

## Stack

- **Python 3.13**
- **FastAPI** — async web framework (Minimal API equivalent)
- **Pydantic v2** — data validation and serialization (System.Text.Json + DataAnnotations equivalent)
- **Uvicorn** — ASGI server (Kestrel equivalent)

## Setup

```bash
# Create virtual environment
python -m venv .venv

# Activate (PowerShell)
.venv\Scripts\Activate.ps1

# Activate (bash/macOS/Linux)
source .venv/bin/activate

# Install dependencies
pip install -r requirements.txt
```

## Run

```bash
uvicorn main:app --reload --port 8000
```

`--reload` enables hot-reload on file changes (like `dotnet watch`).

## Phase Status

- **Phase 0 — Lift-off:** COMPLETE
  - `GET /decolamos` — health check
  - `GET /aircraft` — list all (empty initially)
  - `POST /aircraft` — create with validation

- **Phase 0.5 — Rich Entity:** IN PROGRESS
  - `AircraftV2` with 20 fields covering all major Python types
  - Enums, nested models, optional fields, collections

## API Testing

Use `requests.http` (REST Client / VS Code) or curl:

```bash
curl http://localhost:8000/decolamos
curl http://localhost:8000/aircraft
```
