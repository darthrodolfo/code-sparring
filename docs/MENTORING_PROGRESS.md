# 🥊 CODE-SPARRING — AeroStack Lab (Polyglot Progress Tracker)

> Mentoring mode: **Sparring** (challenge-first, no full solutions).
> Goal: master backend fundamentals + real-world I/O across stacks, then build Flutter + PWA clients.
> Primary stack (Phase 1): **.NET 9 Minimal API (C# 12/13+)**

---

## ✅ North Star
Build the same business capabilities across:
- /backend-csharp (.NET 9 Minimal API) **CURRENT**
- /backend-go (stdlib / Gin)
- /backend-python (FastAPI / Pydantic v2)
- /backend-nodejs (NestJS / Fastify)
Later:
- /frontend-nextjs (Next.js PWA)
- /frontend-flutter (Mobile)

**Rule:** One shared contract (OpenAPI) + contract tests. Each stack must match behavior.

---

## 📍 Current Status

### C# (.NET 9) — Phase 0.5 COMPLETE, Round 1 COMPLETE
- AircraftV2 entity with 20 fields covering all major C# types
- CreateAircraftV2Request DTO (no Id — server generates Guid v7)
- Enums: AircraftRole, AircraftStatus (replaces bool flags)
- Records: GeoLocation, AircraftSpecs (with int? + TimeSpan), ConflictHistory
- POST /aircraft-v2 endpoint — 201 Created, JsonStringEnumConverter configured
- GET /aircraft-v2 (list all)
- GET /aircraft-v2/{id} (by id)
- PUT /aircraft-v2/{id} (update)
- DELETE /aircraft-v2/{id} (delete)
- **SQLite persistence added for all CRUD operations** (cross-stack test data preservation)
- requests.http with full payload + nullable payload + PUT/DELETE test cases
- Verified: enum as string, DateOnly, DateTimeOffset with offset, TimeSpan, nullable value types, nested objects, byte[] as Base64

### Python (FastAPI / Pydantic v2) — Phase 0 COMPLETE, Phase 0.5 COMPLETE, Round 1 COMPLETE
- venv + FastAPI 0.135.1 + Pydantic v2 + Uvicorn scaffolded
- GET /decolamos, GET /aircraft, POST /aircraft with validation
- Enums: AircraftRole, AircraftStatus (str, Enum — serializes as string natively)
- Nested models: GeoLocation, AircraftSpecs (timedelta), ConflictHistory
- AircraftV2 entity: 20 fields covering all major Python types
- CreateAircraftV2Request DTO with Field validators (ge, le, gt, min_length, max_length)
- POST /aircraft-v2 — 201 Created + Location header
- GET /aircraft-v2 — list all (CRUD)
- GET /aircraft-v2/{id} — get by id (CRUD)
- PUT /aircraft-v2/{id} — full update (CRUD)
- DELETE /aircraft-v2/{id} — remove by id (CRUD)
- **SQLite persistence added for all CRUD operations** (shared DB setup)
- **Dependency Injection** via FastAPI `Depends` for `aiosqlite` HTTP requests
- Verified: Decimal as string, date as ISO 8601, datetime as UTC, timedelta as PT format, bytes as Base64, null explicit
- requests.http with full payload + nullable fields + PUT/DELETE test cases

### Next: Checkpoint Actions
- Open Go Phase 0 (next stack baseline)

---

## 🎯 Round 1 — Aircraft Core (Minimal)
### Goal
Prove clean Minimal API structure + strong typing + validation + encapsulated state.

### Scope
- Model: `Aircraft`
- Storage: in-memory (encapsulated, thread-safe enough)
- Endpoints:
  - `GET /aircraft` (list)
  - `POST /aircraft` (create)

### Constraints
- No spaghetti `Program.cs`
- No stringly-typed nonsense
- Prefer modern C# features (records, typed results where it helps)
- Validation must be explicit and consistent

---

## ➕ Round 1.5 — “Types & Validation Discipline” (NEXT)
### Additions
1) Separate domain vs request DTO:
   - `Aircraft` (domain)
   - `CreateAircraftRequest` (DTO)

2) Enums:
   - `Role` (e.g., Fighter, Bomber, Transport, Recon, Trainer, Drone, etc.)
   - `Status` (Active, Maintenance, Retired, Stored, etc.)

3) Collections:
   - `Tags: string[]`
     - normalization: trim, remove empty, distinct (case-insensitive), max 10
     - reject tags > N chars (e.g., 24)

### Definition of Done (DoD)
- POST validates payload and returns consistent errors
- Tags normalized on write (not on read)
- Program.cs stays structured in 4 blocks:
  1) types/records/enums
  2) store
  3) app build
  4) endpoints

---

## 🧠 Key Decisions (Locked-in unless we revisit)
- `Id`: Guid (upgrade to Guid v7 when available/decided)
- `Year`: int with range validation (e.g., 1903..currentYear+1)
- Domain model: `record` (immutability by default)

---

## 🧪 Review Checklist (Jarvis will grade this every round)
- Types: enums, ranges, patterns, no accidental string typing
- Validation: clear, consistent, normalized input
- Encapsulation: storage not exposed; no leaking internal list refs
- Thread-safety: minimal correctness (lock/concurrent collection)
- API correctness: status codes + Created/Location behavior
- Maintainability: file structure, naming, separation of DTO/domain

---

## 🗺️ Roadmap (Capability Rounds)
### Round 2 — Nested types + numeric constraints
- Specs object (range, speed, wingspan, etc.)
- Optional fields + defaulting rules

### Round 3 — Image upload/download (multipart + binary)
- `POST /aircraft/{id}/photo`
- `GET /aircraft/{id}/photo`
- size limits + content-type validation

### Round 4 — Audio upload/stream
- audio-note upload + download
- later: Range support

### Round 5 — Query features
- filtering, searching, pagination, sorting
- validation of query params

### Round 6 — Concurrency, idempotency, basic security
- idempotency key
- basic auth boundary (API key → JWT later)
- rate limiting & request limits

### Round 7 — Persistence + caching + jobs
- Postgres + migrations
- Redis caching
- background processing (thumb/waveform)
- observability (metrics/tracing)

---

## 🧩 Frontends (After backend can serve real I/O)
Start after Round 4:
- Next.js PWA: camera input, offline caching
- Flutter: multipart uploads, file system, playback

---

## 📝 Next Action
Complete Round 1 CRUD endpoints with ConcurrentDictionary storage, then add SQLite persistence.
After C# checkpoint is solid, open Phase 0-1 in next stack (Go, Python, or Node).

# 🥊 AeroStack Lab — IDE AI Context Pack (Phase 0: Lift-off)

## What is happening
I’m doing a polyglot “startup from zero” training.  
The goal is to bootstrap **multiple backend stacks** from scratch and ensure each one:
1) builds
2) runs locally
3) answers a simple health endpoint
4) can serialize at least one real entity (Aircraft)

This is the **Phase 0: Lift-off**. After all stacks “take off”, we will incrementally add validations, DTO discipline, enums, lists, file uploads (images/audio), etc.

## Repo structure
- /backend-csharp   (.NET 9 Minimal API)  ✅ current focus
- /backend-go       (stdlib / Gin later)
- /backend-python   (FastAPI / Pydantic v2) **Phase 0 + 0.5 + Round 1 COMPLETE**
- /backend-nodejs   (NestJS / Fastify)
- /frontend-nextjs  (Next.js PWA - later)
- /frontend-flutter (Flutter - later)

## Phase 0 — Lift-off Contract (MVP for every stack)
### Endpoint A — Server Alive
GET /decolamos
- 200 OK
- Body: "Decolamos"
- Prefer text/plain (JSON also acceptable)

### Entity (v0)
Aircraft
- id: string (UUID)        // generated by server
- model: string            // required, trimmed, 1..80
- manufacturer: string     // required, trimmed, 1..80
- year: integer            // required, range: 1903..(currentYear+1)

### DTO (v0)
CreateAircraftRequest
- model: string
- manufacturer: string
- year: integer

### Endpoint B — Uses the entity (minimum)
GET /aircraft
- 200 OK
- Content-Type: application/json
- Body: [] (empty list is fine for Phase 0)

### Optional (only if fast)
POST /aircraft
- Validates payload using the rules above
- Generates id (UUID)
- Returns 201 Created + created Aircraft

## Key rules (Sparring / Senior discipline)
- Avoid spaghetti in the main entry file (Program.cs / main.go / app.py / main.ts).
- Avoid stringly-typed design. Use correct primitives and validate input.
- Keep state encapsulated (don’t expose a global mutable list directly).
- Keep Phase 0 minimal: do NOT add role/status/tags/media yet.

## Example JSON payload (POST /aircraft)
{
  "model": "F-16C",
  "manufacturer": "Lockheed Martin",
  "year": 1991
}

## What I need from the IDE AI
- Help generate the minimal project scaffold for the stack.
- Implement the Phase 0 contract exactly.
- Keep code clean and minimal (but not sloppy).
- Prefer standard library patterns first; avoid unnecessary dependencies.

## Current task for me (human)
Start with backend-csharp:
- Ensure `dotnet new web` runs
- Implement GET /decolamos and GET /aircraft
- (Optional) POST /aircraft
- Run locally and test with curl or HTTP client