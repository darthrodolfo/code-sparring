# agent.md — AeroStack Lab (Code-Sparring)

> Bootstrap entry point for AI agents. Start here, then load modules as needed.

## Quick Overview

AeroStack Lab is a polyglot backend training ground. The same rich **Aircraft** domain (20-field entity) is implemented across multiple stacks (C#, Python, Go, Node.js/Fastify, Dart, Java) to compare language ergonomics, framework DX, and persistence patterns. An older NestJS safari snapshot remains in the repo only as a legacy reference, not as an active focus stack.

This is a **learning and experimentation project**, not a production system.

## AI Documentation Index

| File | Responsibility |
|------|---------------|
| [context.md](context.md) | Project vision, learning objectives, user profile, domain strategy |
| [architecture.md](architecture.md) | Stack plan, tech choices, anti-patterns, domain model overview |
| [services.md](services.md) | Service registry (directories, ports, commands, key files) |
| [phases.md](phases.md) | Roadmap, current state per stack, cross-stack conclusions |
| [conventions.md](conventions.md) | AI interaction rules, coding standards, mentoring protocol, **agent exclusion list** |

## Entry Point for Claude Code

See [claude.md](claude.md) for Claude Code-specific behavioral rules.

**Important:** Before exploring the codebase, review the **Agent Exclusion List** in [conventions.md](conventions.md) — it lists all generated/build directories that agents must never read or write to (e.g., `bin/`, `obj/`, `node_modules/`, `build/`, `dist/`, `__pycache__/`, `.dart_tool/`, `target/`).

## Study Documentation (reference only)

Personal study notes, learning trackers, and reference materials — not AI context files:

### General
- [implementation-plan.md](../study-docs/general/implementation-plan.md) — AircraftV2 entity specification
- [roadmap-linguagem-zero-ao-avancado.md](../study-docs/general/roadmap-linguagem-zero-ao-avancado.md) — Generic language learning roadmap
- [mentoring-progress.md](../study-docs/general/mentoring-progress.md) — Legacy progress tracker

### Go Backend
- [loops-quick-ref.md](../study-docs/go/loops-quick-ref.md) — Go syntax quick reference
- [progress.md](../study-docs/go/progress.md) — Go backend progress tracker
