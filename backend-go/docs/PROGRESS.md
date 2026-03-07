# Go Backend — Progress Tracker

## Current Status

**Phase 0.5 COMPLETE** — Round 1 CRUD implementation in progress, moving to SQLite integration.

---

## Completed Work

### Phase 0 — Lift-off ✅
- Minimal `net/http` server on `:8080`
- `GET /decolamos` — health check
- `GET /aircraft` — list endpoint (in-memory slice)
- `POST /aircraft` — create with validation
- Error handling: explicit `if err != nil`
- Method routing inside handlers (switch statement)

### Phase 0.5 — Rich Entity ✅
- **AircraftV2 struct:** 20 fields fully implemented
- **Type system:**
  - Enums: `AircraftRole`, `AircraftStatus` (string type aliases + const)
  - Nested: `GeoLocation`, `AircraftSpecs`, `ConflictHistory`
  - Nullable: `*string`, `*int`, `[]byte`
  - Collections: `[]string`, `[]ConflictHistory`, `map[string]string`
  - External types: `uuid.UUID`, `decimal.Decimal`, `time.Time`, `time.Duration`
- **Endpoints:**
  - `GET /aircraft-v2` — list all
  - `POST /aircraft-v2` — create with full validation
- **Serialization tested:**
  - UUIDs → lowercase hyphenated strings
  - Decimals → string (precision-safe)
  - Times → RFC3339 (ISO 8601)
  - Durations → nanoseconds
  - Bytes → Base64
  - Nulls → explicit JSON null

---

## Dependencies

- `github.com/google/uuid` — UUID v4 generation
- `github.com/shopspring/decimal` — High-precision decimal for financial data

Run `go mod tidy` to ensure these are downloaded.

---

## Design Decisions (Locked for Go)

1. **No framework:** Stdlib `net/http` only (stays true to Go's philosophy)
2. **Type aliases for enums:** `type AircraftRole string` + `const` — idiomatic Go
3. **Pointers for nullable:** `*string`, `*int` — Go's way of optional types
4. **Slices for collections:** `[]T` not arrays `[N]T`
5. **JSON tags: `snake_case`** — matches C# and Python conventions
6. **Explicit error handling:** No exceptions, `if err != nil` everywhere
7. **Struct-based handlers:** Functions that take `http.ResponseWriter` and `*http.Request`

---

## Known Limitations & Trade-offs

| Aspect | Go Approach | Why |
|---|---|---|
| Enums | Type alias + const | Go has no native enum; this is idiomatic |
| Optional types | Pointers | `*T` is the Go way; no generics for `Option<T>` |
| Validation | Manual in handler | No validation framework; explicit control |
| JSON serialization | Struct tags | Go's standard approach |
| Duration serialization | Nanoseconds | `time.Duration` is `int64` nanos internally |

---

## Next Actions (for next session)

### Round 1 — Full CRUD + Storage

Add handlers for:

1. **GET /aircraft-v2/{id}** — fetch single aircraft by UUID
   - Extract ID from URL path
   - Query in-memory store
   - Return 404 if not found

2. **PUT /aircraft-v2/{id}** — update aircraft
   - Validate request payload (same rules as POST)
   - Update in store
   - Return updated object

3. **DELETE /aircraft-v2/{id}** — delete aircraft
   - Remove from store
   - Return 204 No Content

4. **Encapsulate store:** Move `aircraftV2Store` to a separate package or type to prevent direct access

### Round 2 — SQLite Persistence

Use `database/sql` with SQLite driver (`github.com/mattn/go-sqlite3`):
- Create schema for aircraft_v2 + related tables (tags, conflicts)
- Replace in-memory map with database queries
- Migrations strategy

---

## Mentoring Guardrails (Go)

- Explain new Go syntax as soon as it appears in snippets:
  - `for range` semantics by collection type
  - `_` discard identifier
  - pointer/address operator `&`
  - short declaration `:=`
  - `if init; condition` form
- Keep explanations short and practical, with C# analogies when useful.
- Reference: `docs/GO_LOOPS_QUICK_REF.md` for loop/range quick lookup.

---

## Testing

Use `requests.http` in VS Code (REST Client extension):

```bash
### Get all AircraftV2
GET http://localhost:8080/aircraft-v2

### Create AircraftV2
POST http://localhost:8080/aircraft-v2
Content-Type: application/json

{
  "model": "F-22A Raptor",
  "manufacturer": "Lockheed Martin",
  ...
}
```

---

## Key Files

- `main.go` — All code (Phase 0–0.5)
- `go.mod`, `go.sum` — Dependencies
- `requests.http` — API test cases
- `docs/README.md` — User-facing guide
- `docs/PROGRESS.md` — This file
- `docs/GO_LOOPS_QUICK_REF.md` — Loop/range cheat sheet

---

## Go Learning Notes

- **stdlib net/http is powerful:** No framework overhead for Phase 0–1
- **Error handling:** Explicit checks make flow clear (vs try/catch)
- **Type system:** Strict but flexible via interfaces and type aliases
- **JSON marshaling:** Struct tags are the contract between code and JSON
- **Concurrency primitives:** Goroutines and channels are available if needed (not yet)

---

**Last Updated:** March 7, 2026
**Status:** Ready for handoff to next contributor
**Confidence Level:** High — All Phase 0.5 requirements met, code compiles and runs.
