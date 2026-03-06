using System.Collections.Concurrent;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.Data.Sqlite;
using Microsoft.Extensions.Options;

var builder = WebApplication.CreateBuilder(args);

builder.Services.ConfigureHttpJsonOptions(options =>
{
    options.SerializerOptions.Converters.Add(
        new System.Text.Json.Serialization.JsonStringEnumConverter());
});
var connectionString = builder.Configuration.GetConnectionString("AeroStackDb");

var app = builder.Build();


using var initConnection = new SqliteConnection(connectionString);
initConnection.Open();

var sqlPath = Path.Combine(Directory.GetCurrentDirectory(), "sqlite", "001_create_tables.sql");
var schemaSql = File.ReadAllText(sqlPath);

using var schemaCmd = initConnection.CreateCommand();
schemaCmd.CommandText = schemaSql;
schemaCmd.ExecuteNonQuery();
var aircraftStore = new ConcurrentDictionary<Guid, AircraftV2>();

app.MapGet("/", () => "Hello AeroStack!");

app.MapGet("/decolamos", () => "Decolamos!");

app.MapGet("/aircraft", () => new List<AircraftV1>());

app.MapGet("/aircraft-v2", () => aircraftStore.Values.ToList());

app.MapGet("/aircraft-v2/{id:guid}", (Guid id) =>
        aircraftStore.TryGetValue(id, out var aircraft)
        ? Results.Ok(aircraft)
        : Results.NotFound());

app.MapDelete("/aircraft-v2/{id:guid}", (Guid id) =>
    aircraftStore.TryRemove(id, out _)
    ? Results.NoContent()
    : Results.NotFound());

app.MapPut("/aircraft-v2/{id:guid}", (Guid id, CreateAircraftV2Request req) =>
{
    if (aircraftStore.ContainsKey(id) == false)
        return Results.NotFound();

    var updated = new AircraftV2
    {
        Id = id,
        Model = req.Model,
        Manufacturer = req.Manufacturer,
        SerialNumber = req.SerialNumber,
        YearOfManufacture = req.YearOfManufacture,
        PriceMillions = req.PriceMillions,
        EmptyWeightKg = req.EmptyWeightKg,
        Status = req.Status,
        Role = req.Role,
        Tags = req.Tags.AsReadOnly(),
        FirstFlightDate = req.FirstFlightDate,
        LastMaintenanceTime = req.LastMaintenanceTime,
        BaseLocation = req.BaseLocation,
        Specs = req.Specs,
        Conflicts = req.Conflicts,
        Metadata = req.Metadata,
        EstimatedUnitsProduced = req.EstimatedUnitsProduced,
        EstimatedActiveUnits = req.EstimatedActiveUnits,
        PhotoUrl = req.PhotoUrl,
        ManualArchive = req.ManualArchive
    };

    aircraftStore[id] = updated;
    return Results.Ok(updated);
});

app.MapPost("/aircraft-v2", (CreateAircraftV2Request req) =>
{
    var aircraft = new AircraftV2
    {
        Id = Guid.CreateVersion7(),
        Model = req.Model,
        Manufacturer = req.Manufacturer,
        SerialNumber = req.SerialNumber,
        YearOfManufacture = req.YearOfManufacture,
        PriceMillions = req.PriceMillions,
        EmptyWeightKg = req.EmptyWeightKg,
        Status = req.Status,
        Role = req.Role,
        Tags = req.Tags.AsReadOnly(),
        FirstFlightDate = req.FirstFlightDate,
        LastMaintenanceTime = req.LastMaintenanceTime,
        BaseLocation = req.BaseLocation,
        Specs = req.Specs,
        Conflicts = req.Conflicts,
        Metadata = req.Metadata,
        EstimatedUnitsProduced = req.EstimatedUnitsProduced,
        EstimatedActiveUnits = req.EstimatedActiveUnits,
        PhotoUrl = req.PhotoUrl,
        ManualArchive = req.ManualArchive
    };

    //return Results.Created($"/aircraft-v2/{aircraft.Id}", aircraft);
    aircraftStore[aircraft.Id] = aircraft;
    return Results.Created($"/aircraft-v2/{aircraft.Id}", aircraft);
});

app.Run();

record AircraftV1(Guid Id, string Model, string Manufacturer, int Year);

record AircraftV2
{
    public required Guid Id { get; init; }
    public required string Model { get; init; }
    public required string Manufacturer { get; init; }
    public string? SerialNumber { get; init; }
    public required int YearOfManufacture { get; init; }
    public required decimal PriceMillions { get; init; }
    public required double EmptyWeightKg { get; init; }
    public required AircraftStatus Status { get; init; }
    public required AircraftRole Role { get; init; }
    public required IReadOnlyList<string> Tags { get; init; }
    public required DateOnly FirstFlightDate { get; init; }
    public required DateTimeOffset LastMaintenanceTime { get; init; }
    public required GeoLocation BaseLocation { get; init; }
    public required AircraftSpecs Specs { get; init; }
    public required List<ConflictHistory> Conflicts { get; init; }
    public required Dictionary<string, string> Metadata { get; init; }
    public int? EstimatedUnitsProduced { get; init; }
    public int? EstimatedActiveUnits { get; init; }
    public Uri? PhotoUrl { get; init; }
    public byte[]? ManualArchive { get; init; }

}

record CreateAircraftV2Request
{
    public required string Model { get; init; }
    public required string Manufacturer { get; init; }
    public string? SerialNumber { get; init; }
    public required int YearOfManufacture { get; init; }
    public required decimal PriceMillions { get; init; }
    public required double EmptyWeightKg { get; init; }
    public required AircraftStatus Status { get; init; }
    public required AircraftRole Role { get; init; }
    public required List<string> Tags { get; init; }
    public required DateOnly FirstFlightDate { get; init; }
    public required DateTimeOffset LastMaintenanceTime { get; init; }
    public required GeoLocation BaseLocation { get; init; }
    public required AircraftSpecs Specs { get; init; }
    public required List<ConflictHistory> Conflicts { get; init; }
    public required Dictionary<string, string> Metadata { get; init; }
    public int? EstimatedUnitsProduced { get; init; }
    public int? EstimatedActiveUnits { get; init; }
    public Uri? PhotoUrl { get; init; }
    public byte[]? ManualArchive { get; init; }

}


enum AircraftRole
{
    Fighter,
    Bomber,
    Transport,
    Trainer,
    Drone,
    Reconnaissance
}

enum AircraftStatus
{
    Active,
    Maintenance,
    Retired,
    Stored
}

record GeoLocation(double Latitude, double Longitude);

record AircraftSpecs(int MaxSpeedKmh, double WingspanMeters,
 int RangeKm, int? MaxAltitudeMeters, TimeSpan FlightEndurance);

record ConflictHistory(string Name, int StartYear, int EndYear);

