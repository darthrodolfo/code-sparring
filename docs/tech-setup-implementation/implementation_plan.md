# AircraftV2 Comprehensive Entity Plan

## Goal
Design the [AircraftV2] entity and its corresponding DTOs to test the maximum number of data types, structures, and native features of C#. This will serve as a rich foundation for the "Sparring" in other languages.

**Critical Learning Requirement:** The USER is studying and requires guidance and reference for efficient learning. The AI must NOT write the final code automatically. The process must be step-by-step, allowing the USER to copy/type the code manually to maximize knowledge retention and study efficiency. The AI will provide snippets, explanations, and review the USER's implementation.

## Proposed Changes

### 1. Enums & Value Objects

We will start by defining the supporting types.

#### `AircraftRole` (Enum)
Tests strongly typed enums and JSON string serialization.
- `Fighter`, `Bomber`, `Transport`, `Trainer`, `Drone`, `Reconnaissance`

#### `AircraftStatus` (Enum)
Tests enum as state machine replacement for bool flags. Richer and more extensible than `IsActive: bool`.
- `Active`, `Maintenance`, `Retired`, `Stored`

#### `GeoLocation` (Record)
Tests nested objects representing coordinates.
- `Latitude` (double)
- `Longitude` (double)

#### `AircraftSpecs` (Record)
Tests nested complex types, nullable value types, and duration.
- `MaxSpeedKmh` (int)
- `WingspanMeters` (double)
- `RangeKm` (int)
- `MaxAltitudeMeters` (int?) — nullable value type (`Nullable<int>`), distinct from nullable reference types
- `FlightEndurance` (TimeSpan) — duration type, native to .NET; serializes as ISO 8601 (e.g., `PT14H30M`); each language handles this differently

#### `ConflictHistory` (Record)
Tests lists of complex objects (1:N simulation). Note: `Duration` as `string` was intentionally avoided — that would be stringly-typed design, which violates the project's core principles.
- `Name` (string)
- `StartYear` (int)
- `EndYear` (int)

### 2. The Entity / DTO: AircraftV2

Here are the fields covering all major C# data types and concepts:

| Field                 | C# Type                      | Concept Tested                                              |
| :-------------------- | :--------------------------- | :---------------------------------------------------------- |
| `Id`                  | `Guid`                       | Unique Identifiers (UUIDv4/v7 — use `Guid.CreateVersion7()` in .NET 9) |
| `Model`               | `string`                     | Basic text, Trim validation, Max Length                     |
| `Manufacturer`        | `string`                     | Regex validation, Non-empty                                 |
| `SerialNumber`        | `string?`                    | Nullable reference type (string), optional field            |
| `YearOfManufacture`   | `int`                        | Integer ranges (e.g., 1903 to current year)                 |
| `PriceMillions`       | `decimal`                    | High precision numbers (financial/money — never use double) |
| `EmptyWeightKg`       | `double`                     | Floating point numbers                                      |
| `Status`              | `AircraftStatus` (Enum)      | Enum as state — replaces stringly-typed bool flags          |
| `Role`                | `AircraftRole` (Enum)        | Enum serialization/deserialization                          |
| `Tags`                | `IReadOnlyList<string>`      | Immutable collection on domain model; request DTO uses `List<string>` |
| `FirstFlightDate`     | `DateOnly`                   | Date without time — .NET-specific type                      |
| `LastMaintenanceTime` | `DateTimeOffset`             | Date and Time with Timezone awareness (vs `DateTime`)       |
| `BaseLocation`        | `GeoLocation`                | Nested record/value object                                  |
| `Specs`               | `AircraftSpecs`              | Nested complex object with nullable value type + TimeSpan   |
| `Conflicts`           | `List<ConflictHistory>`      | List of complex objects (1:N simulation)                    |
| `Metadata`            | `Dictionary<string, string>` | Hash maps, unstructured/dynamic JSON structures             |
| `PhotoUrl`            | `Uri?`                       | Built-in URI type + nullable reference type                 |
| `ManualArchive`       | `byte[]?`                    | Binary payloads (for testing Multipart in later rounds)     |

### Type Concept Coverage Summary

| Category                  | C# Concept                    | Where It Appears                        |
| :------------------------ | :---------------------------- | :-------------------------------------- |
| Integers                  | `int`                         | YearOfManufacture, MaxSpeedKmh, etc.    |
| Nullable value type       | `int?` (`Nullable<int>`)      | MaxAltitudeMeters                       |
| Floating point            | `double`                      | EmptyWeightKg, WingspanMeters           |
| High-precision decimal    | `decimal`                     | PriceMillions                           |
| Text                      | `string`                      | Model, Manufacturer                     |
| Nullable reference type   | `string?`, `Uri?`, `byte[]?`  | SerialNumber, PhotoUrl, ManualArchive   |
| Boolean                   | `bool`                        | Not used — replaced by AircraftStatus enum (intentional design decision) |
| Unique identifier         | `Guid`                        | Id                                      |
| Enum                      | `AircraftRole`, `AircraftStatus` | Role, Status                         |
| Date without time         | `DateOnly`                    | FirstFlightDate                         |
| Date + time + timezone    | `DateTimeOffset`              | LastMaintenanceTime                     |
| Duration                  | `TimeSpan`                    | FlightEndurance (in AircraftSpecs)      |
| Immutable collection      | `IReadOnlyList<string>`       | Tags (domain model)                     |
| Mutable collection        | `List<T>`                     | Conflicts                               |
| Hash map                  | `Dictionary<K, V>`            | Metadata                                |
| Nested record             | `GeoLocation`, `AircraftSpecs`, `ConflictHistory` | Nested types     |
| Built-in type             | `Uri`                         | PhotoUrl                                |
| Binary                    | `byte[]`                      | ManualArchive                           |

> **Note on `bool`:** `IsActive: bool` was intentionally removed. A single boolean can never represent a real status lifecycle (Active → Maintenance → Retired). `AircraftStatus` enum is the correct design. This is a deliberate teaching point about avoiding primitive obsession.

### 3. Implementation Steps in `backend-csharp`

1. Define all types at the bottom of [Program.cs](../../backend-csharp/Program.cs) (acceptable for Phase 0.5; can be split into files in Round 1+).
2. Create a `POST /aircraft-v2` endpoint that accepts a `CreateAircraftV2Request` DTO (without `Id` — server generates it using `Guid.CreateVersion7()`).
3. Add basic Minimal API validation flow.
4. Return `201 Created` echoing the parsed and constructed `AircraftV2` object.

### 4. What the Entity Does NOT Cover (Language Behavior)

The entity covers **data types**. C# fundamentals are also exercised through *behavior*. These are tested via how the endpoints and storage are implemented:

| C# Feature             | Where It Gets Exercised                                    |
| :--------------------- | :--------------------------------------------------------- |
| `async/await`          | Making endpoints async (`async Task<IResult>`)             |
| LINQ                   | Filtering/searching the in-memory list (Round 5)           |
| Pattern matching       | `switch` expressions in validation or status transitions   |
| Generics               | Generic result wrappers, generic validation helpers        |
| Interfaces             | Defining a storage contract (`IAircraftRepository`)        |
| Exception handling     | Global error middleware, `try/catch` in validation         |
| Extension methods      | Custom validators, fluent helpers on collections           |
| `ConcurrentDictionary` | Thread-safe in-memory storage (Round 1)                    |
| `record` vs `class`    | Domain model (`record`) vs request DTO (`class`) distinction |

## Verification Plan

### Manual Verification
1. Run `dotnet run`.
2. Send a `POST /aircraft-v2` request using `curl` or the IDE's HTTP client with a full JSON payload covering all fields.
3. Verify the API returns `201 Created` and echoes the data back correctly — pay attention to:
   - Enum serialized as string (not int)
   - `DateOnly` as `"YYYY-MM-DD"`
   - `DateTimeOffset` as ISO 8601 with offset
   - `TimeSpan` as ISO 8601 duration (e.g., `"PT14H30M"`)
   - `Uri` as string
   - `Guid` as lowercase UUID
   - `byte[]` as Base64 string
