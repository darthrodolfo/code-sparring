# 🥊 AeroStack Lab (Code-Sparring)

Welcome to **AeroStack Lab**, a polyglot progress tracker and backend engineering training ground.

This repository serves as a "code-sparring" environment where the goal is to master backend fundamentals and real-world I/O across multiple technology stacks, and eventually connect them to modern frontend clients.

## 🌟 The North Star

The main objective is to build the exact same business capabilities—managing an `Aircraft` domain with complex properties, validations, and media uploads—across different stacks to learn their idioms and best practices.

### 🏗️ Backend Stacks
- [x] **C# / .NET 9 Minimal API** (Current Focus)
- [ ] **Go** (stdlib / Gin)
- [ ] **Python** (FastAPI / Pydantic v2)
- [ ] **Node.js** (NestJS / Fastify)

### 📱 Frontend Clients (Later Phases)
- [ ] **Next.js** (PWA with offline caching and camera input)
- [ ] **Flutter** (Mobile app with multipart uploads and playback)

**Golden Rule:** All backend APIs must share a single contract (OpenAPI) and adhere to the same behavior and test suite.

## 🚀 Iterative Capability Rounds

Development happens in structured "rounds", introducing new architectural and API design challenges progressively:

- **Phase 0 — Lift-off:** Minimal project scaffold, health endpoint (`GET /decolamos`), and basic entity serialization.
- **Round 1 — Core CRUD & Validation:** In-memory storage, strongly typed models (no stringly-typed design), explicit validation, and basic endpoints (`GET` and `POST` for `Aircraft`).
- **Round 1.5 — Types & Validation Discipline:** Separate Domain vs DTOs, Enums for state/roles, normalized collections (tags).
- **Round 2 — Nested Types:** Numeric constraints, optional fields, and defaulting rules.
- **Round 3 — Image Uploads:** Binary and multipart payloads, size limits, and content-type validation.
- **Round 4 — Audio:** Audio-note uploads, streaming, and HTTP Range support.
- **Round 5 — Querying:** Filtering, searching, pagination, sorting.
- **Round 6 — Resiliency & Security:** Idempotency keys, rate limiting, and basic auth (API key to JWT).
- **Round 7 — Production Readiness:** Postgres database with migrations, Redis caching, background processing, and observability.

## 🧠 Core Engineering Principles

- **No Spaghetti Code:** Keep entry points (like `Program.cs`) clean and structured.
- **Strong Typing & Encapsulation:** Use language primitives correctly, enforce immutability where possible, and never expose internal mutable state.
- **Strict Validation:** Input is always normalized and validated early via dedicated DTOs.
- **Consistent API Behavior:** Correct use of HTTP status codes, headers (like `Location` for created resources), and consistent error shapes.

---
*This is an active mentoring and training project under a challenge-first "sparring" approach.*
