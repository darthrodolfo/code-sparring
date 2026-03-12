# AircraftV2 — Entity Specification

> Comprehensive entity design covering the maximum number of data types, structures, and native features across languages.
> This serves as the reference blueprint for the polyglot backend safari.

---

## Navigation Model

**Background:** 15+ years of .NET/C# experience. Strong transferable fundamentals across stacks.

**The invariant (non-negotiable):**

> Phase 0-1 in all requested stacks before asymmetric advancement in any single one.

This ensures a comparative baseline exists across all languages when sparring begins.

**After Phase 0-1 across all stacks — free navigation:**

- Jump from C# Phase 1 to Phase 5-6 if the momentum is there
- Come back and do Phase 2 across all stacks later
- Switch to Go for a "rest" session, then return to C# — all valid
- No mandatory linear progression per stack after the baseline is set

**Rhythm:** Intense by design. Multiple stacks running simultaneously like parallel "games". The polyglot comparison is the training — not just the individual stack depth.

---

## Supporting Types

### AircraftRole (Enum)

Tests strongly typed enums and JSON string serialization.

Values: `Fighter`, `Bomber`, `Transport`, `Trainer`, `Drone`, `Reconnaissance`

### AircraftStatus (Enum)

Tests enum as state machine replacement for bool flags. Richer and more extensible than `IsActive: bool`.

Values: `Active`, `Maintenance`, `Retired`, `Stored`

### GeoLocation (Record/Struct)

Tests nested objects representing coordinates.

| Field | Type |
|-------|------|
| Latitude | double/float64 |
| Longitude | double/float64 |

### AircraftSpecs (Record/Struct)

Tests nested complex types, nullable value types, and duration.

| Field | Type | Notes |
|-------|------|-------|
| MaxSpeedKmh | int | |
| WingspanMeters | double | |
| RangeKm | int | |
| MaxAltitudeMeters | int? | Nullable value type |
| FlightEndurance | TimeSpan/Duration | ISO 8601 duration (e.g., `PT14H30M`); varies per language |

### ConflictHistory (Record/Struct)

Tests lists of complex objects (1:N simulation). `Duration` as `string` was intentionally avoided (anti stringly-typed design).

| Field | Type | Notes |
|-------|------|-------|
| Name | string | |
| StartYear | int | |
| EndYear | int | |
| RoleInConflict | AircraftRole | Enum reuse inside nested type — enables role-based filtering |

---

## AircraftV2 Entity — Field Specification

| Field | C# Type | Concept Tested |
|-------|---------|----------------|
| Id | Guid | Unique identifiers (UUIDv4/v7) |
| Model | string | Basic text, trim validation, max length |
| Manufacturer | string | Regex validation, non-empty |
| SerialNumber | string? | Nullable reference type, optional field |
| YearOfManufacture | int | Integer ranges (1903 to current year) |
| PriceMillions | decimal | High precision numbers (financial — never use double) |
| EmptyWeightKg | double | Floating point numbers |
| Status | AircraftStatus | Enum as state — replaces bool flags |
| Role | AircraftRole | Enum serialization/deserialization |
| Tags | IReadOnlyList\<string\> | Immutable collection (domain); List\<string\> (DTO) |
| FirstFlightDate | DateOnly | Date without time (.NET-specific) |
| LastMaintenanceTime | DateTimeOffset | Date + time + timezone awareness |
| BaseLocation | GeoLocation | Nested record/value object |
| Specs | AircraftSpecs | Nested complex object with nullable + duration |
| Conflicts | List\<ConflictHistory\> | List of complex objects (1:N simulation) |
| Metadata | Dictionary\<string, string\> | Hash maps, unstructured JSON structures |
| EstimatedUnitsProduced | int? | Nullable value type — unknown for classified aircraft |
| EstimatedActiveUnits | int? | Nullable value type — operational count |
| PhotoUrl | Uri? | Built-in URI type + nullable reference |
| ManualArchive | byte[]? | Binary payloads (multipart in later rounds) |

---

## Type Coverage Summary

| Category | C# Concept | Where It Appears |
|----------|-----------|------------------|
| Integers | int | YearOfManufacture, MaxSpeedKmh |
| Nullable value type | int? (Nullable\<int\>) | MaxAltitudeMeters |
| Floating point | double | EmptyWeightKg, WingspanMeters |
| High-precision decimal | decimal | PriceMillions |
| Text | string | Model, Manufacturer |
| Nullable reference type | string?, Uri?, byte[]? | SerialNumber, PhotoUrl, ManualArchive |
| Boolean | bool | Not used — replaced by AircraftStatus enum |
| Unique identifier | Guid | Id |
| Enum | AircraftRole, AircraftStatus | Role, Status |
| Date without time | DateOnly | FirstFlightDate |
| Date + time + timezone | DateTimeOffset | LastMaintenanceTime |
| Duration | TimeSpan | FlightEndurance (in AircraftSpecs) |
| Immutable collection | IReadOnlyList\<string\> | Tags (domain model) |
| Mutable collection | List\<T\> | Conflicts |
| Hash map | Dictionary\<K, V\> | Metadata |
| Nested record | GeoLocation, AircraftSpecs, ConflictHistory | Nested types |
| Built-in type | Uri | PhotoUrl |
| Binary | byte[] | ManualArchive |

> **Design note on bool:** `IsActive: bool` was intentionally removed. A single boolean cannot represent a real status lifecycle (Active -> Maintenance -> Retired). `AircraftStatus` enum is the correct design. This teaches avoiding primitive obsession.

---

## Behavioral Features (Not Covered by Entity)

The entity covers data types. Language fundamentals are exercised through behavior:

| Feature | Where Exercised |
|---------|----------------|
| async/await | Async endpoints |
| LINQ / functional | Filtering/searching (Round 5) |
| Pattern matching | Switch expressions in validation |
| Generics | Generic result wrappers |
| Interfaces | Storage contract (IAircraftRepository) |
| Exception handling | Global error middleware |
| Extension methods | Custom validators, fluent helpers |
| ConcurrentDictionary | Thread-safe in-memory storage |
| record vs class | Domain model (record) vs DTO (class) |

---

## Implementation Steps (C# Reference)

1. Define all types (enums, records, entity)
2. Create POST /aircraft-v2 with CreateAircraftV2Request DTO
3. Configure JsonStringEnumConverter for enum serialization
4. Add ConcurrentDictionary in-memory storage
5. Add GET /aircraft-v2 (list all)
6. Add GET /aircraft-v2/{id} (lookup by Guid)
7. Add PUT /aircraft-v2/{id} (update)
8. Add DELETE /aircraft-v2/{id} (delete)
9. Add SQLite persistence

---

## Verification Checklist

After implementation, verify the API returns correct serialization:

- Enum serialized as string (not int)
- DateOnly as `"YYYY-MM-DD"`
- DateTimeOffset as ISO 8601 with offset
- TimeSpan as ISO 8601 duration (e.g., `"PT14H30M"`)
- Uri as string
- Guid as lowercase UUID
- byte[] as Base64 string
- Null fields as explicit JSON `null`
