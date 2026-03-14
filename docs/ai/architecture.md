# AeroStack Lab — Architecture

## Backend Safari Plan

The backend safari is a controlled sequence of implementations of the **same domain + CRUD + SQLite** across multiple stacks.

### Stack Order

1. **C# / .NET 9 Minimal API**
2. **Python (FastAPI / Pydantic v2)**
3. **Go (stdlib)**
4. **Node.js puro + Fastify**
5. **Java / Spring Boot**
6. **Dart backend (Dart Frog)**
7. Return to **C# / .NET** for the premium final backend

Legacy reference kept in the repo:
- **Node.js / NestJS + Express** — archived safari snapshot, not part of the active roadmap

### Node.js Stack Decision

The active Node.js/TypeScript stack for premium work and AI integrations is **Fastify** (pure, no NestJS wrapper). The existing NestJS backend remains only as an archived safari reference in the repo.

- **Fastify (pure)** — Tier 1, plugin-based, schema-driven, used for AI integrations
- **NestJS + Express** — safari complete, stop-point reached, no further expansion

---

## Anti-Pattern: Near-Duplicate Stacks

**NEVER suggest two stacks that only differ by HTTP adapter or minor infrastructure swap.**

- :x: `NestJS + Express` AND `NestJS + Fastify` — same mental model, trivial adapter swap
- :white_check_mark: `NestJS + Express` AND `Node puro + Fastify` — genuinely different approaches

Before suggesting any stack variation, ask: **"Does this teach a meaningfully different mental model, or just swap an infrastructure detail?"** If the latter, do not suggest it.

---

## Technologies in Scope

### Main Stacks

| Stack | Framework | Persistence | Status |
|-------|-----------|-------------|--------|
| C# / .NET 9 | Minimal API | SQLite | Active |
| Python | FastAPI / Pydantic v2 | SQLite (aiosqlite) | Safari complete |
| Go | stdlib (net/http) | SQLite (go-sqlite3) | Safari complete |
| Node.js | Fastify (puro) | SQLite (better-sqlite3) | **Tier 1 — active focus for AI integrations** |
| Node.js | NestJS + Express | SQLite (better-sqlite3) | Archived safari reference only |
| Java | Spring Boot | SQLite (planned) | Planned |
| Dart | Dart Frog | SQLite (sqlite3) | Safari complete |

### Common Persistence Baseline

**SQLite** is the common persistence layer across all safari stacks. Future conceptual exploration may include PostgreSQL, MariaDB, Redis, MongoDB, DynamoDB, CockroachDB, Supabase, ClickHouse, and pgvector — but these are **not** part of the mandatory safari stop-point.

### Final Premium Topics (C# / .NET + Node.js / Fastify)

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

Full entity specification: [implementation-plan.md](../study-docs/general/implementation-plan.md)
