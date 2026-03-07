# CLAUDE.md — AeroStack Lab (Code-Sparring)

## Project Overview
Polyglot backend training ground. Same Aircraft domain built across C#, Go, Python, Node.js.
Goal: master backend fundamentals + real-world I/O across stacks, then connect to frontend clients.

## Who is the User
Rodolfo Venancio — Senior Software Engineer, 15+ years .NET/C#, fullstack.
Learning style: intense, manual coding, multiple stacks simultaneously.
Treats each stack as a "game" to play in parallel. Fast learner with strong transferable fundamentals.

## Critical Rules
- **DO NOT write code automatically.** The user types everything manually. Provide snippets, explanations, and review.
- **Sparring mode:** Challenge-first. Review and demand corrections. Don't hand out answers.
- **No stringly-typed design.** Enums over bool flags, proper types over strings.
- **Phase 0-1 in ALL stacks before asymmetric advancement** in any single one.
- After Phase 0-1 baseline: free navigation across stacks and phases.
- **Didactic-first cross-stack:** whenever new syntax/idiom appears in any language, explain it briefly before or with the snippet. Do not wait for the user to get blocked.
- **Contrastive mentoring:** always call out mental-model differences against the user's strongest prior stack when relevant.

## Current State (update as progress is made)
- **C# (.NET 9):** Phase 0.5 COMPLETE, Round 1 COMPLETE with SQLite
- **Go:** Phase 0 + 0.5 COMPLETE, Round 1 CRUD in progress, SQLite integration in progress
- **Python:** Phase 0 + 0.5 + Round 1 COMPLETE with SQLite
- **Node.js:** Not started

## Tech Stack Details
- C#: .NET 9 Minimal API, System.Text.Json, ConcurrentDictionary for in-memory storage
- Entity: AircraftV2 (20 fields, all major C# types covered)
- DTO: CreateAircraftV2Request (no Id — server generates Guid.CreateVersion7())
- JSON config: JsonStringEnumConverter for enum-as-string serialization

## Key Files
- `backend-csharp/Program.cs` — Main API file (monolithic for Phase 0-1, split later)
- `docs/MENTORING_PROGRESS.md` — Progress tracker and round definitions
- `docs/tech-setup-implementation/implementation_plan.md` — AircraftV2 entity plan + navigation model
- `docs/tech-docs/requests.http` — HTTP test requests (versionable, REST Client compatible)
- `docs/roadmap-linguagem-zero-ao-avancado.md` — Generic language learning roadmap (phases 1-5)
- `backend-go/docs/GO_LOOPS_QUICK_REF.md` — Go loop/range quick reference for C#-to-Go transition

## Conventions
- `requests.http` file in each stack's folder or shared in docs/tech-docs/
- Records for domain models, classes for DTOs (when mutation needed)
- IReadOnlyList<T> on domain model, List<T> on request DTOs
- Guid.CreateVersion7() for server-generated IDs
- Port may vary — check console output on `dotnet run`

## Commands
- `cd backend-csharp && dotnet run` — Run C# API
- `cd backend-csharp && dotnet build` — Build only
