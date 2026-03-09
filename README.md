# AeroStack Lab (Code-Sparring)

Polyglot backend engineering training ground. The same rich **Aircraft** domain (20-field entity with enums, nested types, collections, decimals, dates, and binary payloads) is implemented across multiple technology stacks to compare language idioms, framework DX, persistence patterns, and architectural clarity.

This is a **learning and experimentation project**, not a production system. Development follows a challenge-first "sparring" approach with AI-assisted mentoring.

## Project Structure

```
code-sparring/
├── backend-csharp/          # .NET 9 Minimal API
├── backend-python/          # FastAPI + Pydantic v2
├── backend-go/              # Go stdlib (net/http)
├── backend-node-next-js/    # NestJS + Express + TypeScript
├── backend_dart/            # Dart Frog
├── docs/
│   ├── ai/                  # AI agent context (start at agent.md)
│   └── study-docs/          # Personal study notes and references
└── README.md
```

## Backend Safari

Each stack goes through the same progression — **scaffold → rich entity → CRUD → SQLite** — then stops. The goal is comparative learning, not building N production systems.

| Stack | Framework | Port | Status |
|-------|-----------|------|--------|
| C# / .NET 9 | Minimal API | 5202 | Round 1 COMPLETE (CRUD + SQLite) |
| Python | FastAPI / Pydantic v2 | 8000 | Round 1 COMPLETE (CRUD + SQLite) |
| Go | stdlib (net/http) | 8080 | Round 1 COMPLETE (CRUD + SQLite) |
| Node.js | NestJS + Express | 3000 | Round 1 COMPLETE (CRUD + SQLite) |
| Dart | Dart Frog | 8080 | Phase 1 COMPLETE (paused before CRUD) |
| Node.js puro | Fastify | — | Planned |
| Java | Spring Boot | — | Planned |

### Safari Phases

| Phase | Scope |
|-------|-------|
| Phase 0 | Project scaffold, health endpoint, first request |
| Phase 1 | Data models, JSON serialization, validation, language idioms |
| Phase 2 | Full CRUD (GET all, GET by id, POST, PUT, DELETE) |
| Phase 3 | SQLite persistence, schema, repository layer |

After the safari completes, C# / .NET returns as the **flagship premium backend** with advanced rounds (auth, caching, AI integrations, observability, Postgres).

## Domain Model — AircraftV2

A deliberately rich entity (20 fields) chosen to stress-test each language's type system:

- **Primitives:** integers, floats, high-precision decimals, booleans
- **Strings & Enums:** `AircraftRole`, `AircraftStatus` (replaces primitive bool flags)
- **Dates:** date-only, datetime with timezone, durations
- **Nullable types:** both value and reference types
- **Collections:** lists (tags), maps
- **Nested types:** `GeoLocation`, `AircraftSpecs`, `ConflictHistory`
- **Binary:** image/audio payloads, URIs, UUIDs

Full specification: [implementation-plan.md](docs/study-docs/general/implementation-plan.md)

## Engineering Principles

- **Strong Typing & Encapsulation** — use language primitives correctly, enforce immutability, never expose internal mutable state
- **Strict Validation** — input normalized and validated early via dedicated DTOs
- **Consistent API Behavior** — correct HTTP status codes, `Location` headers, consistent error shapes
- **No Stringly-Typed Design** — proper enums and typed fields over raw strings
- **Monolith-First in Early Phases** — split files only when readability demands it

## Running a Stack

Each backend has a `requests.http` file (VS Code REST Client) for API testing.

```bash
# C#
cd backend-csharp && dotnet run

# Python
cd backend-python && uvicorn main:app --reload --port 8000

# Go
cd backend-go && go run .

# Node.js / NestJS
cd backend-node-next-js && npm run start:dev

# Dart
cd backend_dart && dart_frog dev
```

## For AI Agents

AI context files live in [`docs/ai/`](docs/ai/). Start at [`agent.md`](docs/ai/AGENT.md) — it bootstraps all other context modules (architecture, services, phases, conventions). Claude Code-specific rules are in [`claude.md`](docs/ai/CLAUDE.md).

Key rules for agents working in this repo:
- **Do not write code into the user's files** — the user types everything manually
- **Do not read or modify generated/build artifacts** — see the full exclusion list in [`conventions.md`](docs/ai/conventions.md)
- **Framework-first** — always recommend frameworks over stdlib-level approaches
- **Respect stop-points** — do not expand a safari stack beyond CRUD + SQLite

---

*Active mentoring and training project. Author: Rodolfo Venancio.*
