# AeroStack Lab — Service Registry

## Directory Map

| Directory | Stack | Port | Status |
|-----------|-------|------|--------|
| `backend-csharp/` | .NET 9 Minimal API | 5202 | Round 1 Complete |
| `backend-python/` | FastAPI + Pydantic v2 | 8000 | Round 1 Complete |
| `backend-go/` | Go stdlib (net/http) | 8080 | Round 1 Complete |
| `backend-node-next-js/` | NestJS + Express + TypeScript | 3000 | Round 1 Complete |
| `backend-dart/` | Dart Frog | 8080 | In Progress |

### Planned (not yet scaffolded)

- Node.js puro + Fastify
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

### Node.js / NestJS

```bash
cd backend-node-next-js && npm run start:dev
cd backend-node-next-js && npm run build
```

### Dart

```bash
cd backend-dart && dart_frog dev
cd backend-dart && dart_frog build
```

---

## Tech Stack Specifics

### C#

- .NET 9 Minimal API, System.Text.Json
- ConcurrentDictionary for in-memory storage (early phases)
- Guid.CreateVersion7() for server-generated IDs (.NET 9)
- JsonStringEnumConverter for enum-as-string serialization

### Python

- FastAPI 0.135.1, Pydantic v2, Uvicorn
- Enums: `(str, Enum)` pattern (serializes as string natively)
- Decimal serializes as string, timedelta as ISO 8601 (PT format)
- Persistence: aiosqlite + FastAPI Depends DI

### Go

- stdlib: net/http, encoding/json, database/sql
- External: google/uuid, shopspring/decimal, mattn/go-sqlite3
- Enums: type alias + const pattern (Go has no native enum)
- Nullable: pointers (`*string`, `*int`)
- Serialization: Decimal as string, time.Time as RFC3339, time.Duration as nanoseconds

### Node.js / NestJS

- NestJS + Express, TypeScript strict mode
- experimentalDecorators + emitDecoratorMetadata enabled
- Validation: class-validator + class-transformer (`@IsString`, `@IsEnum`, `@ValidateNested`, `@Type`)
- PartialType from @nestjs/mapped-types for update DTO
- DB: better-sqlite3 (synchronous), transactions via db.transaction()
- Architecture: DatabaseModule (@Global) + AircraftModule (Controller/Service/Repository)
- OnModuleInit lifecycle hook for schema initialization
- Entity: interface (not class) with optional fields via `?`

### Dart

- Dart Frog framework (Very Good Ventures)
- File-based routing (routes/ directory)
- Hot reload via `dart_frog dev`
- dart_frog CLI for scaffolding (`dart_frog create`, `dart_frog new route`)
- SQLite via `sqlite3` package (dart:ffi based)

---

## Key Files

| File | Purpose |
|------|---------|
| `backend-csharp/Program.cs` | Main C# API file (monolithic for early phases) |
| `docs/study-docs/general/implementation_plan.md` | AircraftV2 entity plan + navigation model |
| `docs/study-docs/general/roadmap-linguagem-zero-ao-avancado.md` | Generic language learning roadmap |
| `docs/study-docs/go/GO_LOOPS_QUICK_REF.md` | Go loop/range quick reference |
| `docs/study-docs/go/PROGRESS.md` | Go backend progress tracker |
| `backend-go/GO_README.md` | Go backend documentation |
| `docs/study-docs/general/MENTORING_PROGRESS.md` | Legacy progress tracker (content migrated to docs/ai/) |
| Each stack: `requests.http` | API test cases (VS Code REST Client) |

---

## General Repo Principle

Start monolithic in early phases when it helps learning. Split files/folders only when the monolith starts obscuring responsibilities.
