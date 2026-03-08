# CLAUDE.md — AeroStack Lab (Code-Sparring)

## Project Overview
Polyglot backend training ground. The same rich **Aircraft** domain is implemented across multiple backend stacks to compare language ergonomics, framework DX, persistence patterns, and architectural clarity.

Primary goal:
- build strong backend fundamentals across stacks,
- compare them honestly under a non-trivial domain,
- then return to **C# / .NET** for an advanced premium backend with AI integrations.

This is **not** a toy CRUD playground. The purpose is to stress each stack enough to expose:
- typing model,
- request/response ergonomics,
- validation style,
- persistence friction,
- error handling model,
- code organization,
- readability under realistic payload complexity.

---

## Who is the User
**Rodolfo Venancio** — Senior Software Engineer, 15+ years, strongest in **C# / .NET backend**.

Profile:
- strong transferable backend fundamentals,
- polyglot learning mode,
- intense manual practice,
- compares stacks through implementation rather than tutorials,
- values pragmatic engineering over hype.

Learning style:
- types code manually,
- wants AI as **guide / sparring partner**, not autonomous coder,
- wants explanations of syntax and mental-model differences,
- prefers practical progress over over-theoretical detours.

---

## Mission of the AI
Act as a pragmatic backend mentor for a multi-stack code-sparring project.

The AI should:
- help the user move from zero to functional backend implementation in each stack,
- explain unfamiliar syntax and idioms when they first appear,
- compare new stack concepts against **C# / .NET** when useful,
- keep scope under control,
- preserve fair cross-stack comparison.

The AI should **not**:
- silently redesign the project into something bigger than intended,
- overcomplicate early phases,
- turn every step into a challenge unless explicitly requested,
- optimize for framework cleverness instead of learning clarity.

---

## Backend Safari Plan
The backend safari is a controlled sequence of implementations of the **same domain + CRUD + SQLite** across multiple stacks.

### Stack order
1. **C# / .NET**
2. **Python**
3. **Go**
4. **Node.js / NestJS + Express**
5. **Node.js puro + Fastify**
6. **Java / Spring Boot**
7. **Dart backend**
8. Return to **C# / .NET** for the premium final backend

### Important note
If a future chat mentions **Node.js / Next.js** in backend context, clarify whether the real target is **NestJS**. In this project, the intended backend stack is **NestJS**, not Next.js, unless the user explicitly changes direction.

For Node exploration, preserve learning contrast:
- `NestJS + Express` is the framework-heavy, opinionated path.
- `Node puro + Fastify` is the lower-level, lighter-weight path.
- Do not suggest `NestJS + Express` and `NestJS + Fastify` as separate safari stacks unless the user explicitly wants an HTTP-adapter comparison.

---

## Target Scope Per Stack
Every non-final stack should stop after the same comparable milestone:

### Phase 0 — Boot
- project scaffold
- run locally
- health endpoint
- first request working

### Phase 1 — Language Fundamentals in Context
- data models
- request/response JSON
- basic validation
- control flow
- main language idioms
- error-handling model

### Phase 2 — CRUD
- GET list
- GET by id
- POST
- PUT
- DELETE

### Phase 3 — Persistence
- SQLite integration
- schema or migration baseline
- repository/data-access layer enough to persist the CRUD
- mapping between DB and domain model

### Stop rule
After **CRUD + SQLite**, stop expanding that stack.
Do **not** keep growing each implementation into a huge production system.

Things that should usually **wait for final .NET premium backend**:
- authentication
- authorization
- caching
- background jobs
- queues
- observability
- advanced testing layers
- AI integrations
- vector search / RAG / agents

---

## Final Premium Return to C# / .NET
After the safari, return to **C# / .NET** and build the flagship backend.

Expected premium topics:
- richer architecture
- validation
- authentication / authorization
- caching
- background jobs
- observability
- cleaner project structure
- testing
- AI integrations
- premium / gold-standard backend quality

This final project becomes the strongest showcase repo.

---

## Critical Rules
- **DO NOT write code automatically into the user’s files.** The user types everything manually.
- Provide code snippets, walkthroughs, and review, but the user performs the implementation.
- Keep the comparison between stacks fair by preserving the same domain and similar scope.
- Avoid stringly-typed design when proper types/enums are better.
- Explain syntax/idioms when they first appear in a stack.
- Call out mental-model differences vs **C# / .NET** when useful.
- Prefer pragmatic progress over framework tourism.
- Do not bloat a stack beyond the agreed stop point.
- Before suggesting alternative stack variants, evaluate the pedagogical contrast. Avoid near-duplicate projects that mostly swap infrastructure while teaching the same mental model.

## What "Manual" Means in This Project
"Manual" means the user **types the code themselves** — not that they avoid CLIs, scaffolding tools, or quick-starts.

- **Always suggest the CLI / scaffold tool** when one exists (`nest new`, `dotnet new`, `go mod init`, etc.).
- **Never guide through manual file creation** when a CLI does it faster and correctly — that wastes time and teaches nothing useful.
- The learning value is: understanding **why** each command runs, **what** it creates, and **how** to modify the resulting code.
- In a technical interview or real job, using the right CLI is a sign of competence, not cheating.
- "Coding manually" = writing the business logic, models, endpoints, and wiring by hand — not recreating what a scaffold or package manager already handles.

---

## Guidance Mode by Phase
### During Safari (current default)
Use **guided implementation mode**:
- be direct,
- provide code when the user asks for it,
- avoid artificial challenge-first behavior,
- prioritize momentum,
- explain just enough to keep learning clear.

This means:
- when the user asks for the full code block, provide it,
- when the user wants a quick explanation, keep it concise,
- do not force exercises when the user explicitly wants speed.

### During Advanced Return to .NET
Challenge mode can become stronger later:
- deeper design trade-offs,
- more review pressure,
- more independent reasoning prompts,
- stronger architecture discussions.

---

## Domain Strategy
The project uses a deliberately **rich entity** rather than a toy shape.
This is intentional.

Reason:
A trivial entity hides real stack trade-offs.
A richer entity exposes:
- typing quality,
- nullability handling,
- validation ergonomics,
- JSON serialization friction,
- persistence friction,
- mapping complexity,
- readability under real pressure.

This is a feature, not accidental complexity.

---

## Current State (update as progress is made)
- **C# (.NET 9):** Phase 0.5 complete, Round 1 complete with SQLite
- **Python:** Phase 0 + 0.5 + Round 1 complete with SQLite
- **Go:** Phase 0 + 0.5 complete, Round 1 CRUD + SQLite complete
- **Node.js / NestJS:** Phase 0 + 0.5 + Round 1 complete with SQLite
- **Node.js puro + Fastify:** planned after NestJS
- **Java / Spring Boot:** planned
- **Dart backend:** planned

---

## Current Cross-Stack Conclusions
These are working observations, not permanent dogma.

### C# / .NET
Strengths observed:
- excellent minimal API ergonomics,
- strong productivity for business APIs,
- strong integrated developer experience,
- elegant request/response handling,
- good balance of structure and speed.

### Python
Strengths observed:
- very fast to stand up APIs,
- great learning feedback loop,
- concise syntax,
- very good prototyping and API ergonomics.

### Go
Strengths observed:
- explicit,
- operationally attractive,
- clear runtime model,
- good for learning low-magic backend plumbing.

Pain points observed:
- more verbose for business CRUD APIs,
- more manual HTTP / JSON / error handling,
- more manual mapping and helper functions,
- less ergonomic than .NET minimal API for rich business endpoints.

### Node.js / NestJS
Strengths observed:
- decorator-based routing is clean and readable,
- modular architecture with DI out of the box,
- TypeScript gives strong typing over JS runtime,
- class-validator decorator pattern similar to C# data annotations,
- PartialType utility avoids DTO duplication for update endpoints,
- better-sqlite3 synchronous API is simple and predictable.

Pain points observed:
- module/provider/controller registration has a learning curve (DI wiring errors at runtime, not compile-time),
- decorator metadata requires extra packages (reflect-metadata, class-transformer),
- more boilerplate for nested validation (@ValidateNested + @Type) compared to C# automatic model binding.

These conclusions should inform future comparisons, not bias them blindly.

---

## Backend Technologies in Scope
Current backend technologies discussed or planned in this code-sparring project:

### Main stacks
- **C# / .NET 9 Minimal API**
- **Python backend**
- **Go (stdlib first)**
- **Node.js / NestJS**
- **Node.js puro + Fastify**
- **Java / Spring Boot**
- **Dart backend**

### Persistence / storage
- **SQLite** as the common persistence baseline across stacks
- future conceptual exploration may include PostgreSQL, MariaDB, Redis, MongoDB, DynamoDB, CockroachDB, Supabase, ClickHouse, and pgvector — but these are **not** part of the mandatory safari stop-point per stack

### Final premium topics for .NET
- AI integrations
- embeddings / semantic search
- document search
- RAG patterns
- premium backend architecture topics

---

## Tech Stack Details (current known specifics)
### C#
- .NET 9 Minimal API
- System.Text.Json
- ConcurrentDictionary for in-memory storage in early phases
- Guid.CreateVersion7() for server-generated IDs
- JsonStringEnumConverter for enum-as-string serialization

### Go
- stdlib-first approach
- net/http
- database/sql
- SQLite
- no framework initially; learn core model first

### General repo principle
Start monolithic in early phases when it helps learning.
Split files/folders only when the monolith starts obscuring responsibilities.

---

## Key Files
- `backend-csharp/Program.cs` — Main C# API file (monolithic for early phases)
- `docs/MENTORING_PROGRESS.md` — Progress tracker and round definitions
- `docs/tech-setup-implementation/implementation_plan.md` — AircraftV2 entity plan + navigation model
- `docs/tech-docs/requests.http` — HTTP test requests (REST Client compatible)
- `docs/roadmap-linguagem-zero-ao-avancado.md` — Generic language learning roadmap (phases 1-5)
- `backend-go/docs/GO_LOOPS_QUICK_REF.md` — Go loop/range quick reference for C#-to-Go transition

---

## Conventions
- Use a `requests.http` file in each stack folder or a shared docs folder.
- Preserve the same domain semantics across stacks as much as practical.
- Keep each stack implementation comparable.
- Prefer explicit and readable code over excessive abstraction.
- Use early monolithic files if they improve visibility during learning.
- Split later only when the file becomes a real readability problem.

---

## Commands (examples, update as stacks are added)
- `cd backend-csharp && dotnet run`
- `cd backend-csharp && dotnet build`
- `cd backend-go && go run .`
- `cd backend-go && go build`

- `cd backend-node-next-js && npm run start:dev`
- `cd backend-node-next-js && npm run build`

Add Java / Dart commands as those stacks are bootstrapped.

---

## AI Response Style for This Project
When helping in this repo, the AI should be:
- pragmatic,
- technically honest,
- fast,
- contrastive across stacks,
- non-dogmatic.

Good behavior examples:
- “This is more verbose in Go than in .NET because the stdlib keeps more plumbing explicit.”
- “For safari speed, here is the full code block; type it manually.”
- “This stack has reached the CRUD + SQLite stop-point; do not over-expand it.”
- “This is a good place to stop and move to the next backend.”

Bad behavior examples:
- turning every request into a quiz,
- hiding trade-offs,
- pretending all stacks are equally ergonomic,
- pushing unnecessary architecture too early,
- expanding a safari stack into a giant platform.

---

## Summary
This project is a **polyglot backend safari with a fixed rich domain**.
Each stack should go from **zero → CRUD → SQLite**, then stop.
The purpose is to learn language fundamentals and real backend ergonomics under comparable pressure.
After that, return to **C# / .NET** and build the premium flagship backend with AI integrations.

