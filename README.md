# AeroStack Lab

**Polyglot backend architecture comparison + AI-integrated engineering** — the same rich domain implemented across multiple stacks to evaluate language ergonomics, framework DX, and persistence patterns, then extended with production-grade AI capabilities on the flagship C# / .NET backend.

One domain. Multiple ecosystems. Same contract. Every backend exposes the same API shape over a deliberately complex **Aircraft** entity (20 fields spanning enums, nested types, nullable values, decimals, dates, durations, collections, and binary payloads). The entity surfaces real-world friction that trivial CRUD demos never reveal: serialization edge cases, DTO separation, null-safety discipline, and persistence mapping. The .NET implementation then goes further — integrating LLM orchestration, RAG pipelines, and AI-powered features as first-class architectural components.

## Implemented Stacks

| Stack | Framework | Persistence | Port |
|-------|-----------|-------------|------|
| **C# / .NET 9** | Minimal API | SQLite | 5202 |
| **Python** | FastAPI + Pydantic v2 | SQLite (aiosqlite) | 8000 |
| **Go** | stdlib (net/http) | SQLite (go-sqlite3) | 8080 |
| **Node.js / TypeScript** | NestJS + Express | SQLite (better-sqlite3) | 3000 |
| **Dart** | Dart Frog | SQLite (sqlite3/dart:ffi) | 8080 |

Each stack follows the same progression: **scaffold → rich entity modeling → full CRUD → SQLite persistence** with schema design, transactions, and domain-to-DB mapping.

## Cross-Stack Comparison

| Dimension | C# / .NET 9 | Python | Go | Node.js / NestJS |
|-----------|-------------|--------|----|-------------------|
| **Validation** | Minimal API model binding | Pydantic v2 field validators | Manual struct validation | class-validator decorators |
| **Nullability** | `T?` nullable types | `str \| None` (3.10+) | Pointers (`*string`, `*int`) | `?` optional fields |
| **Enum strategy** | Native enum + `JsonStringEnumConverter` | `(str, Enum)` dual-inherit | Type alias + `const` iota | TypeScript `enum` (string values) |
| **DI / Wiring** | Built-in DI container | FastAPI `Depends` | Manual (no framework DI) | `@Module` / `@Injectable` |
| **DB access** | Raw SQL + Microsoft.Data.Sqlite | aiosqlite (async) | `database/sql` + transactions | better-sqlite3 (sync) |
| **DTO separation** | Record types (request/response) | Pydantic models (in/out) | Separate structs | class-validator classes + `PartialType` |
| **ID generation** | `Guid.CreateVersion7()` | `uuid4()` | `google/uuid` | `crypto.randomUUID()` |
| **Decimal handling** | Native `decimal` | `Decimal` → string serialization | `shopspring/decimal` → string | `number` (IEEE 754) |

## Domain Model — AircraftV2

A reference entity engineered to expose typing quality, serialization friction, and persistence complexity across ecosystems:

- **Enums:** `AircraftRole`, `AircraftStatus` — replaces primitive `bool` flags (anti-primitive-obsession)
- **Nested types:** `GeoLocation`, `AircraftSpecs`, `ConflictHistory` (with `StartYear`/`EndYear`, not stringly-typed `Duration`)
- **Nullable semantics:** both value and reference type nullability per language
- **High-precision numerics:** `decimal` / `Decimal` for financial-grade fields
- **Temporal types:** date-only, datetime with timezone, durations (ISO 8601 / TimeSpan)
- **Collections:** `IReadOnlyList<string>` (domain) vs `List<string>` (DTO) — immutability at the boundary
- **Binary:** image/audio payload paths, URIs, UUIDs

## Architecture Principles

- **Strong typing & encapsulation** — language primitives used correctly, immutability enforced, internal mutable state never exposed
- **Strict input validation** — normalized and validated early via dedicated DTOs, not domain models
- **Consistent API surface** — correct HTTP status codes, `Location` headers for created resources, uniform error shapes
- **No stringly-typed design** — proper enums and typed nested records over raw strings and maps
- **Monolith-first** — single-file entry points in early phases, split only when readability demands it

## Running Any Stack

Each backend includes a `requests.http` file ([VS Code REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)) with pre-built API test cases.

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

## Project Structure

```
aerostack-lab/
├── backend-csharp/            .NET 9 Minimal API — CRUD + SQLite
├── backend-python/            FastAPI + Pydantic v2 — CRUD + SQLite
├── backend-go/                Go stdlib — CRUD + SQLite
├── backend-node-next-js/      NestJS + Express + TypeScript — CRUD + SQLite
├── backend_dart/              Dart Frog — CRUD + SQLite
├── docs/
│   ├── ai/                    AI agent context modules
│   └── study-docs/            Architecture specs and reference materials
└── README.md
```

## AI Integration Layer (C# / .NET)

The flagship .NET backend extends beyond CRUD into a full **AI-integrated architecture**, embedding intelligent capabilities as first-class system components — not bolted-on API wrappers.

| Capability | Approach |
|------------|----------|
| **Chat Completion** | OpenAI / Azure OpenAI SDK — structured prompt management, conversation context, streaming responses (SSE) |
| **Embeddings & Vector Search** | Text-to-embedding pipelines, vector storage (pgvector), cosine similarity retrieval |
| **RAG Pipelines** | Retrieval-Augmented Generation — document ingestion, chunking strategies, context-aware LLM responses |
| **AI Orchestration** | Semantic Kernel / Microsoft.Extensions.AI — multi-step agent workflows, function calling, tool use |
| **Structured Output** | Type-safe LLM responses mapped to C# models — no stringly-typed prompt parsing |
| **Content Processing** | Summarization, classification, and metadata extraction over domain entities |

These integrations follow the same engineering discipline as the rest of the codebase: strong typing, clean separation of concerns, proper DI wiring, and testable abstractions over external AI services.

## Roadmap

| Phase | Focus |
|-------|-------|
| **Advanced API patterns** | Image/audio uploads (multipart + streaming), HTTP Range, pagination, filtering, sorting |
| **Security & resilience** | Idempotency keys, rate limiting, JWT authentication |
| **Infrastructure** | PostgreSQL + migrations, Redis caching, background processing, observability |
| **Cloud deployment** | CI/CD pipelines, containerization, cloud-native architecture |

## For AI Agents

Context modules for AI-assisted development live in [`docs/ai/`](docs/ai/). Start at [`AGENT.md`](docs/ai/AGENT.md) for the bootstrap entry point. See [`conventions.md`](docs/ai/conventions.md) for interaction rules and the **agent exclusion list** (generated directories that must not be read or modified).

---

*Rodolfo Venancio — Senior Software Engineer*
