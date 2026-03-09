# Backend Go — Phase 0 Lift-off

## Overview
Go stdlib (`net/http`) implementation of AeroStack backend. Phase 0 focuses on minimal viable API structure and correct HTTP handling patterns.

## Stack
- **Language:** Go 1.21+ (or latest)
- **HTTP:** stdlib `net/http` (Phase 0), Gin later
- **Serialization:** stdlib `encoding/json`

## Running

```bash
cd backend-go
go run main.go
```

Server starts on `:8080`.

## Phase 0 Contract

### GET /decolamos
Server alive check.
- **Status:** 200 OK
- **Body:** `"Decolamos"` (text/plain)

### GET /aircraft
List all aircraft (Phase 0: empty list).
- **Status:** 200 OK
- **Content-Type:** application/json
- **Body:** `[]` or populated array

### POST /aircraft
Create aircraft.
- **Status:** 201 Created
- **Body:** Created `Aircraft` object with server-generated UUID

**Payload:**
```json
{
  "model": "F-16C",
  "manufacturer": "Lockheed Martin",
  "year": 1991
}
```

**Validation Rules:**
- `model`: required, trimmed, 1–80 chars
- `manufacturer`: required, trimmed, 1–80 chars
- `year`: required, 1903 ≤ year ≤ (currentYear + 1)

## Testing

Use `requests.http` with REST Client extension, or curl:

```bash
curl http://localhost:8080/decolamos
curl http://localhost:8080/aircraft
curl -X POST http://localhost:8080/aircraft \
  -H "Content-Type: application/json" \
  -d '{"model":"F-16C","manufacturer":"Lockheed Martin","year":1991}'
```

## Architecture (Phase 0)

```go
// Handler registration
http.HandleFunc("/decolamos", decolamosHandler)

// Method routing in handler
func decolamosHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
    // response
}
```

Key pattern: `http.HandleFunc` with manual method routing inside the handler.

## Phase 0.5 — Rich Entity (COMPLETE)

### AircraftV2 Struct
- **20 fields** covering all major Go types
- **Type aliases for enums:** `AircraftRole`, `AircraftStatus` (string-based, not `iota`)
- **Nested structs:** `GeoLocation`, `AircraftSpecs` (with `*int` nullable field), `ConflictHistory`
- **Pointers for nullable:** `*string`, `*int` (Go way of optional types)
- **Collections:** `[]string` (tags), `[]ConflictHistory` (conflicts), `map[string]string` (metadata)
- **Special types:** `uuid.UUID` (google/uuid), `decimal.Decimal` (shopspring/decimal), `time.Time`, `time.Duration`
- **Endpoints:** `GET /aircraft-v2`, `POST /aircraft-v2` with full validation

### Verified Serialization
- `uuid.UUID` → JSON string (lowercase hyphenated)
- `decimal.Decimal` → JSON string (preserves precision)
- `time.Time` → JSON RFC3339 (ISO 8601 with timezone)
- `time.Duration` → JSON number (nanoseconds)
- `[]byte` → JSON string (Base64 encoded)
- Null handling: `nil` pointers → `null` in JSON

## Round 1 — Full CRUD (COMPLETE)

### Endpoints
- **GET /aircraft-v2** — list all aircraft
- **GET /aircraft-v2/{id}** — fetch by UUID
- **POST /aircraft-v2** — create with validation
- **PUT /aircraft-v2/{id}** — update with validation
- **DELETE /aircraft-v2/{id}** — delete (cascades)

### Storage
- SQLite with 3 tables: aircraft_v2, aircraft_tags, aircraft_conflicts
- Transactions on write operations
- Foreign key constraints with CASCADE delete

## Next Phases (Optional)
- **Round 3:** Query features (filtering, pagination, sorting)
- **Round 4+:** Caching, advanced concurrency, Gin framework migration

## Standards
- Enum serialization: string-based type aliases (Go idiom)
- UUID generation: `github.com/google/uuid`
- Decimal: `github.com/shopspring/decimal`
- Date/Time: `time.Time` (RFC3339 JSON)
- Error handling: explicit `if err != nil`, JSON error responses
- JSON tags: `snake_case` for consistency across all stacks
