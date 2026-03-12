# AeroStack Lab — Phases & Current State

## Target Scope Per Safari Stack

Every non-final stack should stop after the same comparable milestone:

### Phase 0 — Boot
- Project scaffold, run locally, health endpoint, first request working

### Phase 1 — Language Fundamentals in Context
- Data models, request/response JSON, basic validation, control flow, main language idioms, error-handling model

### Phase 2 — CRUD
- GET list, GET by id, POST, PUT, DELETE

### Phase 3 — Persistence
- SQLite integration, schema/migration baseline, repository/data-access layer, DB-to-domain mapping

### Stop Rule

After **CRUD + SQLite**, stop expanding that stack. Do **not** keep growing each implementation into a huge production system.

Topics that wait for the Tier 1 stacks (.NET + Node.js/Fastify):
- Authentication / authorization
- Caching
- Background jobs / queues
- Observability
- Advanced testing layers
- AI integrations (vector search, RAG, agents)

---

## Current State

| Stack | Status | Notes |
|-------|--------|-------|
| C# (.NET 9) | Phase 0.5 + Round 1 **COMPLETE** | Full CRUD + SQLite |
| Python (FastAPI) | Phase 0 + 0.5 + Round 1 **COMPLETE** | Full CRUD + SQLite |
| Go (stdlib) | Phase 0 + 0.5 + Round 1 **COMPLETE** | Full CRUD + SQLite |
| Node.js / NestJS | Phase 0 + 0.5 + Round 1 **COMPLETE** | Full CRUD + SQLite |
| Dart (Dart Frog) | Phase 0 + 0.5 + Round 1 **COMPLETE** | Full CRUD + SQLite |
| Node.js puro + Fastify | Phase 0 + 0.5 + Round 1 **COMPLETE** | Full CRUD + SQLite |
| Java / Spring Boot | Delayed/Postponed | — |

### Next Actions

1. Safari complete for all primary stacks (Java / Spring Boot postponed)
2. Next: return to C# / .NET + Node.js / Fastify as Tier 1 for pre-AI foundations + AI integrations
3. NestJS stays at stop-point — safari reference only, no further expansion

---

## Capability Rounds (Premium .NET Backend)

These rounds apply to the **final C# / .NET** premium backend after the safari completes:

| Round | Focus |
|-------|-------|
| 1 | Core CRUD + in-memory storage → SQLite |
| 1.5 | Types & validation discipline (DTO separation, enums, tag normalization) |
| 2 | Nested types + numeric constraints |
| 3 | Image uploads (multipart + binary) |
| 4 | Audio uploads + streaming + Range support |
| 5 | Query features (filtering, searching, pagination, sorting) |
| 6 | Concurrency, idempotency, basic security (API key → JWT) |
| 7 | Persistence (Postgres + migrations), Redis caching, background processing, observability |

---

## Cross-Stack Conclusions

Working observations, not permanent dogma.

### C# / .NET

**Strengths:** Excellent minimal API ergonomics, strong productivity for business APIs, strong integrated DX, elegant request/response handling, good balance of structure and speed.

### Python

**Strengths:** Very fast to stand up APIs, great learning feedback loop, concise syntax, very good prototyping and API ergonomics.

### Go

**Strengths:** Explicit, operationally attractive, clear runtime model, good for learning low-magic backend plumbing.

**Pain points:** More verbose for business CRUD, more manual HTTP/JSON/error handling, more manual mapping and helper functions, less ergonomic than .NET minimal API for rich business endpoints.

### Node.js / NestJS

**Strengths:** Decorator-based routing is clean and readable, modular architecture with DI out of the box, TypeScript gives strong typing over JS runtime, class-validator decorator pattern similar to C# data annotations, PartialType avoids DTO duplication, better-sqlite3 sync API is simple and predictable.

**Pain points:** Module/provider/controller wiring has a learning curve (DI errors at runtime, not compile-time), decorator metadata requires extra packages (reflect-metadata, class-transformer), more boilerplate for nested validation compared to C# automatic model binding.

These conclusions should inform future comparisons, not bias them blindly.

---

## Tier 1 Premium Stacks — C# / .NET + Node.js / Fastify

After the safari, **C# / .NET** and **Node.js / Fastify** advance together as Tier 1 stacks.

Expected premium topics (both stacks):
- Richer architecture
- Validation
- Authentication / authorization
- Caching
- Background jobs
- Observability
- Cleaner project structure
- Testing
- AI integrations (embeddings, RAG, agents, MCP, streaming)
- Premium / gold-standard backend quality

Both stacks deliver the same capabilities with different architectural approaches:
- **.NET** — opiniated, convention-based, enterprise patterns (Polly, Serilog, EF Core)
- **Fastify** — plugin-based, schema-driven, composition over inheritance (Pino, Ajv, native hooks)
