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

## Next Phases
- **Phase 0.5:** AircraftV2 (20 fields), enums, nested models
- **Round 1:** CRUD endpoints, in-memory storage (encapsulated)
- **Round 2+:** SQLite, Gin framework, etc.

## Standards
- Enum serialization: string (native in Go)
- UUID generation: `google.golang.org/uuid` or stdlib `crypto/rand`
- Date/Time: `time.Time` (RFC3339 JSON)
- Error handling: consistent, typed responses
