# AeroStack Lab — Mentoring Progress Tracker

> Historical record of implementation progress across all safari stacks.
> For current state and roadmap, see [docs/ai/phases.md](../../ai/phases.md).

---

## Per-Stack Implementation Log

### C# (.NET 9) — Round 1 COMPLETE

- AircraftV2 entity with 20 fields covering all major C# types
- CreateAircraftV2Request DTO (no Id — server generates Guid v7)
- Enums: AircraftRole, AircraftStatus (replaces bool flags)
- Records: GeoLocation, AircraftSpecs (with `int?` + `TimeSpan`), ConflictHistory
- POST /aircraft-v2 — 201 Created, JsonStringEnumConverter configured
- GET /aircraft-v2 (list all), GET /aircraft-v2/{id}, PUT /aircraft-v2/{id}, DELETE /aircraft-v2/{id}
- SQLite persistence for all CRUD operations
- Verified: enum as string, DateOnly, DateTimeOffset with offset, TimeSpan, nullable value types, nested objects, `byte[]` as Base64

### Python (FastAPI / Pydantic v2) — Round 1 COMPLETE

- venv + FastAPI 0.135.1 + Pydantic v2 + Uvicorn
- Enums: AircraftRole, AircraftStatus (`(str, Enum)` pattern — serializes as string natively)
- Nested models: GeoLocation, AircraftSpecs (timedelta), ConflictHistory
- AircraftV2 entity: 20 fields covering all major Python types
- CreateAircraftV2Request DTO with Field validators (`ge`, `le`, `gt`, `min_length`, `max_length`)
- Full CRUD: GET list, GET by id, POST (201 + Location), PUT, DELETE
- SQLite persistence via aiosqlite + FastAPI `Depends` DI
- Verified: Decimal as string, date as ISO 8601, datetime as UTC, timedelta as PT format, bytes as Base64, null explicit

### Go (stdlib) — Round 1 COMPLETE

- net/http stdlib (no framework)
- Enums: AircraftRole, AircraftStatus (type alias + const pattern)
- Nested structs: GeoLocation, AircraftSpecs (with `*int` nullable), ConflictHistory
- AircraftV2 struct: 20 fields, UUID + Decimal external types
- Full CRUD with SQLite persistence:
  - 3-table schema (aircraft_v2, aircraft_tags, aircraft_conflicts)
  - Foreign keys with ON DELETE CASCADE
  - Transactions (BeginTx/Commit/Rollback) on all writes
  - Tag normalization + full validation
- Verified: uuid.UUID, decimal.Decimal, time.Time (RFC3339), time.Duration (nanoseconds), `[]byte` (Base64)

### Node.js / NestJS — Round 1 COMPLETE

- `nest new .` scaffold + TypeScript + ESLint + Prettier
- Enums: AircraftStatus, AircraftCategory (TypeScript native enum with string values)
- Nested validation: ConflictHistoryDto with `@ValidateNested` + `@Type` (class-transformer)
- AircraftV2 entity: interface with optional fields (`?` syntax)
- CreateAircraftV2Request DTO with decorators: `@IsString`, `@IsEnum`, `@IsInt`, `@Min`, `@Max`, `@IsOptional`, `@IsArray`, `@ArrayMaxSize`
- UpdateAircraftV2Request via PartialType (all fields optional, inherits validation rules)
- Full CRUD with SQLite persistence (better-sqlite3):
  - 3-table schema with FK CASCADE
  - Transactions via `db.transaction()` on all writes
  - `hydrate()` method for DB row to entity mapping (snake_case to camelCase)
  - Boolean stored as INTEGER (0/1)
- Architecture: DatabaseModule (@Global) + AircraftModule (Controller/Service/Repository)
- OnModuleInit lifecycle hook for schema initialization
- Exception-based error handling (NotFoundException -> 404 automatic)

### Dart (Dart Frog) — Round 1 COMPLETE

- Framework: Dart Frog (Very Good Ventures) — file-based routing, CLI scaffolding
- Boot completed (project scaffold, local run, first routes)
- Phase 1 fundamentals completed (rich model/DTO flow, JSON handling, basic validation and request/response mapping)
- Current checkpoint: paused before Phase 2 (full CRUD) and Phase 3 (SQLite)
- Port: 8080

---

## Capability Rounds (Premium .NET Backend)

These rounds define the progression for the final C# / .NET premium backend after the safari:

### Round 1 — Aircraft Core (Minimal)

- Model: Aircraft with in-memory storage (ConcurrentDictionary)
- Endpoints: GET /aircraft (list), POST /aircraft (create)
- Constraints: no spaghetti Program.cs, no stringly-typed design, modern C# features (records, typed results)

### Round 1.5 — Types and Validation Discipline

- Separate domain model vs request DTO (Aircraft vs CreateAircraftRequest)
- Enums: Role (Fighter, Bomber, Transport...), Status (Active, Maintenance, Retired...)
- Tags: `string[]` with normalization (trim, remove empty, distinct case-insensitive, max 10, reject > 24 chars)
- Program.cs structured in 4 blocks: types/records/enums, store, app build, endpoints

### Round 2 — Nested Types and Numeric Constraints

- Specs object (range, speed, wingspan, etc.)
- Optional fields + defaulting rules

### Round 3 — Image Upload/Download

- `POST /aircraft/{id}/photo`, `GET /aircraft/{id}/photo`
- Multipart + binary, size limits, content-type validation

### Round 4 — Audio Upload/Stream

- Audio-note upload + download
- HTTP Range support

### Round 5 — Query Features

- Filtering, searching, pagination, sorting
- Query param validation

### Round 6 — Concurrency, Idempotency, Basic Security

- Idempotency key
- Basic auth boundary (API key -> JWT later)
- Rate limiting and request limits

### Round 7 — Persistence, Caching, Jobs

- Postgres + migrations
- Redis caching
- Background processing (thumb/waveform)
- Observability (metrics/tracing)

---

## Key Decisions (Locked)

- `Id`: Guid (Guid.CreateVersion7() in .NET 9)
- `Year`: int with range validation (1903..currentYear+1)
- Domain model: `record` (immutability by default)
- `AircraftStatus` enum replaces `IsActive: bool` (anti primitive-obsession)

---

## Review Checklist (Per Round)

- Types: enums, ranges, patterns, no accidental string typing
- Validation: clear, consistent, normalized input
- Encapsulation: storage not exposed, no leaking internal list refs
- Thread-safety: minimal correctness (lock/concurrent collection)
- API correctness: status codes + Created/Location behavior
- Maintainability: file structure, naming, separation of DTO/domain
