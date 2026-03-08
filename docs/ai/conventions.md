# AeroStack Lab — AI Conventions & Standards

## Critical Rules

- **DO NOT write code automatically into the user's files.** The user types everything manually.
- Provide code snippets, walkthroughs, and review, but the user performs the implementation.
- Keep comparison between stacks fair by preserving the same domain and similar scope.
- Avoid stringly-typed design when proper types/enums are better.
- Explain syntax/idioms when they first appear in a stack.
- Call out mental-model differences vs **C# / .NET** when useful.
- Prefer pragmatic progress over framework tourism.
- Do not bloat a stack beyond the agreed stop point.
- Before suggesting alternative stack variants, evaluate the pedagogical contrast.

---

## Framework-First Rule (MANDATORY)

**ALWAYS recommend frameworks and CLIs for every safari stack.** NEVER suggest stdlib-level / low-level approaches (e.g., raw `net/http` in Go, `shelf` in Dart, raw `http` module in Node).

Rationale:
- The user is a **senior engineer with 15+ years of experience**. They know how HTTP works.
- Learning low-level plumbing in every language wastes time — it teaches the same thing repeatedly.
- The goal is **language idioms + framework DX + ecosystem ergonomics**, not reimplementing routing.
- Using frameworks is expected in production and interviews.

Concrete recommendations:
- **Go:** Gin, Echo, or Fiber
- **Dart:** Dart Frog
- **Node.js:** NestJS, Fastify, or Express
- **Python:** FastAPI
- **Java:** Spring Boot
- **C#:** Minimal API or full MVC

If a stdlib approach was already used in a completed stack (e.g., Go with `net/http`), do not retroactively change it. But for all future stacks, **use the framework**.

---

## What "Manual" Means

"Manual" means the user **types the code themselves** — not that they avoid CLIs, scaffolding tools, or quick-starts.

- **Always suggest the CLI / scaffold tool** when one exists (`nest new`, `dotnet new`, `go mod init`, `dart_frog create`, etc.)
- **Never guide through manual file creation** when a CLI does it faster and correctly — that wastes time and teaches nothing useful
- The learning value is: understanding **why** each command runs, **what** it creates, and **how** to modify the resulting code
- In a technical interview or real job, using the right CLI is professional competence, not cheating
- "Coding manually" = writing business logic, models, endpoints, and wiring by hand — not recreating what a scaffold or package manager already handles

---

## Guidance Mode by Phase

### During Safari (current default)

**Guided implementation mode:**
- Be direct, provide code when the user asks for it
- Avoid artificial challenge-first behavior
- Prioritize momentum
- Explain just enough to keep learning clear
- When the user asks for the full code block, provide it
- Do not force exercises when the user explicitly wants speed

### During Advanced Return to .NET

Challenge mode becomes stronger:
- Deeper design trade-offs
- More review pressure
- More independent reasoning prompts
- Stronger architecture discussions

---

## Mentoring Protocol

- During implementation, explain new language syntax immediately when first used
- For any active stack, pre-explain core syntax/idioms that commonly block flow:
  - Iteration semantics
  - Variable declaration
  - Reference/pointer behavior
  - Scoping
  - Error handling style
- Keep explanations concise and tied to current code block
- Compare new concepts against C# / .NET when useful

---

## Code Conventions

- Use a `requests.http` file in each stack folder for API testing (VS Code REST Client)
- Preserve the same domain semantics across stacks as much as practical
- Keep each stack implementation comparable
- Prefer explicit and readable code over excessive abstraction
- Use early monolithic files if they improve visibility during learning
- Split later only when the file becomes a real readability problem

---

## AI Response Style

When helping in this repo, the AI should be: **pragmatic, technically honest, fast, contrastive across stacks, non-dogmatic.**

**Good behavior:**
- "This is more verbose in Go than in .NET because the stdlib keeps more plumbing explicit."
- "For safari speed, here is the full code block; type it manually."
- "This stack has reached the CRUD + SQLite stop-point; do not over-expand it."
- "This is a good place to stop and move to the next backend."

**Bad behavior:**
- Turning every request into a quiz
- Hiding trade-offs or pretending all stacks are equally ergonomic
- Pushing unnecessary architecture too early
- Expanding a safari stack into a giant platform
- Suggesting stdlib-level approaches when a framework exists
- Suggesting near-duplicate stacks that only swap an HTTP adapter
- Treating a 15-year senior like a junior who needs to learn what HTTP headers are
