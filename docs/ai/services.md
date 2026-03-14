# AeroStack Lab - Service Registry

## Directory Map

| Directory | Stack | Port | Status |
|-----------|-------|------|--------|
| `backend-csharp/` | .NET 9 Minimal API | 5202 | Round 1 complete |
| `backend-python/` | FastAPI + Pydantic v2 | 8000 | Round 1 complete |
| `backend-go/` | Go stdlib (`net/http`) | 8080 | Round 1 complete |
| `backend-node-fastify/` | Fastify + TypeScript | 3001 | Round 1 complete - Tier 1 active focus for AI integrations |
| `backend-node-nest-js/` | NestJS + Express + TypeScript | 3000 | Archived safari reference only |
| `backend_dart/` | Dart Frog | 8080 | Round 1 complete |

### Planned

- Java / Spring Boot

---

## Commands

### C#

```bash
cd backend-csharp && dotnet run
cd backend-csharp && dotnet build
```

### Python

```bash
cd backend-python && uvicorn main:app --reload --port 8000
```

### Go

```bash
cd backend-go && go run .
cd backend-go && go build
```

### Node.js / Fastify (Tier 1 active focus)

```bash
cd backend-node-fastify && npm run dev
cd backend-node-fastify && npm run build:ts
cd backend-node-fastify && npm test
```

### Node.js / NestJS (archived safari reference)

```bash
cd backend-node-nest-js && npm run start:dev
cd backend-node-nest-js && npm run build
```

### Dart

```bash
cd backend_dart && dart_frog dev
cd backend_dart && dart_frog build
```

---

## Tech Stack Specifics

### C#

- .NET 9 Minimal API, System.Text.Json
- ConcurrentDictionary for in-memory storage in early phases
- `Guid.CreateVersion7()` for server-generated IDs (.NET 9)
- `JsonStringEnumConverter` for enum-as-string serialization

### Python

- FastAPI 0.135.1, Pydantic v2, Uvicorn
- Enums with `(str, Enum)` pattern
- Decimal serializes as string, `timedelta` as ISO 8601 PT format
- Persistence: `aiosqlite` + FastAPI `Depends`

### Go

- stdlib: `net/http`, `encoding/json`, `database/sql`
- External packages: `google/uuid`, `shopspring/decimal`, `mattn/go-sqlite3`
- Enums via type alias + const pattern
- Nullable values via pointers (`*string`, `*int`)
- Serialization: Decimal as string, `time.Time` as RFC3339, `time.Duration` as nanoseconds

### Node.js / Fastify (Tier 1 active focus - AI integrations)

- Fastify + TypeScript, plugin-based architecture
- AutoLoad plugin system with `src/plugins/` and `src/routes/`
- DB plugin in `src/plugins/database.ts` exposing `fastify.db`
- SQLite via `better-sqlite3`
- Manual `rowToAircraft()` mapper from snake_case DB rows to camelCase domain objects
- Types in `src/routes/aircraft/aircraft.types.ts`
- Runtime validation via JSON Schema + Ajv
- Response/error contract centralized in `src/plugins/contract.ts`
- Port: 3001

### Node.js / NestJS (archived safari reference - no further expansion)

- NestJS + Express, TypeScript strict mode
- `experimentalDecorators` + `emitDecoratorMetadata`
- Validation via `class-validator` + `class-transformer`
- `PartialType` from `@nestjs/mapped-types` for update DTOs
- SQLite via `better-sqlite3`
- Modular architecture: `DatabaseModule` + `AircraftModule`
- OnModuleInit lifecycle hook for schema initialization

### Dart

- Dart Frog framework
- File-based routing
- Hot reload via `dart_frog dev`
- SQLite via `sqlite3`

---

## Key Files

| File | Purpose |
|------|---------|
| `backend-csharp/Program.cs` | Main C# API file for early phases |
| `docs/study-docs/general/implementation-plan.md` | AircraftV2 entity plan + navigation model |
| `docs/study-docs/general/roadmap-linguagem-zero-ao-avancado.md` | Generic language learning roadmap |
| `docs/study-docs/go/loops-quick-ref.md` | Go loop/range quick reference |
| `docs/study-docs/go/progress.md` | Go backend progress tracker |
| `backend-go/GO_README.md` | Go backend documentation |
| `docs/study-docs/general/mentoring-progress.md` | Legacy progress tracker |
| `backend-node-fastify/requests.http` | Fastify API test cases |
| `backend-node-nest-js/` | Archived NestJS safari reference |

---

## General Repo Principle

Start monolithic in early phases when it helps learning. Split files and folders only when the monolith starts obscuring responsibilities.
