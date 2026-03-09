# Go Backend — Progress Tracker

> Stack: Go stdlib (net/http) | Port: 8080 | Status: Round 1 COMPLETE

---

## Completed Phases

### Phase 0 — Lift-off

- Minimal `net/http` server on `:8080`
- `GET /decolamos` — health check
- `GET /aircraft` — list endpoint (in-memory slice)
- `POST /aircraft` — create with validation
- Error handling: explicit `if err != nil`
- Method routing inside handlers (switch statement)

### Phase 0.5 — Rich Entity

- **AircraftV2 struct:** 20 fields fully implemented
- **Type system:**
  - Enums: `AircraftRole`, `AircraftStatus` (string type aliases + const)
  - Nested: `GeoLocation`, `AircraftSpecs`, `ConflictHistory`
  - Nullable: `*string`, `*int`, `[]byte`
  - Collections: `[]string`, `[]ConflictHistory`, `map[string]string`
  - External types: `uuid.UUID`, `decimal.Decimal`, `time.Time`, `time.Duration`
- **Endpoints:** `GET /aircraft-v2`, `POST /aircraft-v2` with full validation
- **Serialization verified:**
  - UUIDs -> lowercase hyphenated strings
  - Decimals -> string (precision-safe)
  - Times -> RFC3339 (ISO 8601)
  - Durations -> nanoseconds
  - Bytes -> Base64
  - Nulls -> explicit JSON null

### Round 1 — Full CRUD + SQLite

- **Endpoints:**
  - GET /aircraft-v2 — list all with nested tags/conflicts
  - GET /aircraft-v2/{id} — fetch single by UUID
  - POST /aircraft-v2 — 201 Created + Location header + transactions
  - PUT /aircraft-v2/{id} — update with validation + transactions
  - DELETE /aircraft-v2/{id} — delete with CASCADE cleanup
- **SQLite persistence:**
  - 3-table schema: aircraft_v2, aircraft_tags, aircraft_conflicts
  - Foreign keys with ON DELETE CASCADE
  - Transactions (BeginTx/Commit/Rollback) on all writes
  - Tag normalization + full validation

---

## Dependencies

| Package | Purpose |
|---------|---------|
| github.com/google/uuid | UUID v4 generation |
| github.com/shopspring/decimal | High-precision decimal for financial data |
| github.com/mattn/go-sqlite3 | SQLite driver (CGO) |

Run `go mod tidy` to ensure all dependencies are downloaded.

---

## Design Decisions (Locked)

1. **No framework:** stdlib `net/http` only (completed before framework-first rule was established)
2. **Type aliases for enums:** `type AircraftRole string` + `const` — idiomatic Go
3. **Pointers for nullable:** `*string`, `*int` — Go's way of optional types
4. **Slices for collections:** `[]T` not arrays `[N]T`
5. **JSON tags: snake_case** — matches C# and Python conventions
6. **Explicit error handling:** No exceptions, `if err != nil` everywhere
7. **Struct-based handlers:** Functions taking `http.ResponseWriter` and `*http.Request`

---

## Go vs C# Trade-offs

| Aspect | Go | C# Equivalent |
|--------|-----|---------------|
| Enums | Type alias + const (no native enum) | `enum` keyword with `[JsonStringEnumConverter]` |
| Optional types | Pointers (`*T`) | `Nullable<T>` / `T?` |
| Validation | Manual in handler | Data annotations / FluentValidation |
| JSON serialization | Struct tags | `System.Text.Json` attributes |
| Duration serialization | Nanoseconds (int64) | ISO 8601 (TimeSpan) |
| Error handling | `if err != nil` (explicit) | try/catch (exception-based) |
| DI | Manual / function params | Built-in DI container |

---

## Key Files

| File | Purpose |
|------|---------|
| `backend-go/main.go` | All code (monolithic for safari) |
| `backend-go/go.mod`, `go.sum` | Dependencies |
| `backend-go/requests.http` | API test cases |
| `backend-go/GO_README.md` | User-facing guide |

---

## Learning Notes

- **stdlib net/http is powerful** but verbose for rich CRUD APIs compared to .NET Minimal API
- **Error handling:** Explicit `if err != nil` checks make flow clear but add boilerplate (vs try/catch)
- **Type system:** Strict but flexible via interfaces and type aliases
- **JSON marshaling:** Struct tags are the contract between code and JSON
- **No generics for Option\<T\>:** Pointers are the idiomatic Go way for optional values
- **Concurrency primitives:** Goroutines and channels available but not used in safari scope
