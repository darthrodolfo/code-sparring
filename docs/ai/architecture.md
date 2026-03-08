# AeroStack Lab — Architecture

## Backend Safari Plan

The backend safari is a controlled sequence of implementations of the **same domain + CRUD + SQLite** across multiple stacks.

### Stack Order

1. **C# / .NET 9 Minimal API**
2. **Python (FastAPI / Pydantic v2)**
3. **Go (stdlib)**
4. **Node.js / NestJS + Express**
5. **Node.js puro + Fastify**
6. **Java / Spring Boot**
7. **Dart backend (Dart Frog)**
8. Return to **C# / .NET** for the premium final backend

### Node.js Clarification

If a future chat mentions **Node.js / Next.js** in backend context, clarify whether the real target is **NestJS**. The intended backend stack is **NestJS**, not Next.js, unless the user explicitly changes direction.

For Node exploration, preserve learning contrast:
- `NestJS + Express` — framework-heavy, opinionated path
- `Node puro + Fastify` — lower-level, lighter-weight path
- Do not suggest both as separate safari stacks unless the user explicitly wants an HTTP-adapter comparison

---

## Anti-Pattern: Near-Duplicate Stacks

**NEVER suggest two stacks that only differ by HTTP adapter or minor infrastructure swap.**

- :x: `NestJS + Express` AND `NestJS + Fastify` — same mental model, trivial adapter swap
- :white_check_mark: `NestJS + Express` AND `Node puro + Fastify` — genuinely different approaches

Before suggesting any stack variation, ask: **"Does this teach a meaningfully different mental model, or just swap an infrastructure detail?"** If the latter, do not suggest it.

---

## Technologies in Scope

### Main Stacks

| Stack | Framework | Persistence |
|-------|-----------|-------------|
| C# / .NET 9 | Minimal API | SQLite |
| Python | FastAPI / Pydantic v2 | SQLite (aiosqlite) |
| Go | stdlib (net/http) | SQLite (go-sqlite3) |
| Node.js | NestJS + Express | SQLite (better-sqlite3) |
| Node.js | Fastify (puro) | SQLite (planned) |
| Java | Spring Boot | SQLite (planned) |
| Dart | Dart Frog | SQLite (planned) |

### Common Persistence Baseline

**SQLite** is the common persistence layer across all safari stacks. Future conceptual exploration may include PostgreSQL, MariaDB, Redis, MongoDB, DynamoDB, CockroachDB, Supabase, ClickHouse, and pgvector — but these are **not** part of the mandatory safari stop-point.

### Final Premium Topics (C# / .NET only)

- AI integrations (embeddings, semantic search, RAG, agents)
- Authentication / authorization
- Caching, background jobs, observability
- Premium backend architecture
- Testing layers

---

## Domain Model

The AircraftV2 entity covers 20 fields spanning all major type categories per language:

- Integers, floats, high-precision decimals
- Nullable value types and reference types
- Enums: `AircraftRole`, `AircraftStatus` (replaces bool flags — intentional anti-primitive-obsession)
- Dates (date-only, datetime+timezone), durations
- Collections (lists, maps)
- Nested records/structs: `GeoLocation`, `AircraftSpecs`, `ConflictHistory`
- Binary payloads, URIs, UUIDs

Full entity specification: [implementation_plan.md](../study-docs/general/implementation_plan.md)
