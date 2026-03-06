using System.Collections.Concurrent;
using System.Globalization;
using System.Runtime.InteropServices;
using System.Text.Json;
using System.Transactions;
using Microsoft.AspNetCore.Mvc.Filters;
using Microsoft.AspNetCore.Razor.TagHelpers;
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

app.MapGet("/aircraft-v2", () =>
{
    using var connection = CreateConnection();
    using var listCommand = connection.CreateCommand();
    listCommand.CommandText = "SELECT * FROM aircraft_v2";
    using var reader = listCommand.ExecuteReader();

    var aircraftList = new List<AircraftV2>();
    while (reader.Read())
    {
        var id = Guid.Parse(reader.GetString(reader.GetOrdinal("id")));

        using var tagCommand = connection.CreateCommand();
        tagCommand.CommandText = "SELECT tag FROM aircraft_tags WHERE aircraft_id = @id";
        tagCommand.Parameters.AddWithValue("@id", id.ToString());

        var tagList = new List<string>();
        using var tagReader = tagCommand.ExecuteReader();
        while (tagReader.Read()) tagList.Add(tagReader.GetString(0));


        using var conflictsCommand = connection.CreateCommand();
        conflictsCommand.CommandText = @"SELECT name, start_year, end_year FROM aircraft_conflicts
         where aircraft_id = @id";
        conflictsCommand.Parameters.AddWithValue("@id", id.ToString());

        var conflictList = new List<ConflictHistory>();
        using var conflictReader = conflictsCommand.ExecuteReader();
        while (conflictReader.Read())
            conflictList.Add(new ConflictHistory(conflictReader.GetString(0),
            conflictReader.GetInt32(1),
            conflictReader.GetInt32(2)));

        aircraftList.Add(new AircraftV2
        {
            Id = id,
            Model = reader.GetString(reader.GetOrdinal("model")),
            Manufacturer = reader.GetString(reader.GetOrdinal("manufacturer")),
            SerialNumber = reader.IsDBNull(reader.GetOrdinal("serial_number")) ? null : reader.GetString(reader.GetOrdinal("serial_number")),
            YearOfManufacture = reader.GetInt32(reader.GetOrdinal("year_of_manufacture")),
            PriceMillions = decimal.Parse(reader.GetString(reader.GetOrdinal("price_millions")), CultureInfo.InvariantCulture),
            EmptyWeightKg = reader.GetDouble(reader.GetOrdinal("empty_weight_kg")),
            Status = Enum.Parse<AircraftStatus>(reader.GetString(reader.GetOrdinal("status"))),
            Role = Enum.Parse<AircraftRole>(reader.GetString(reader.GetOrdinal("role"))),
            Tags = tagList.AsReadOnly(),
            FirstFlightDate = DateOnly.Parse(reader.GetString(reader.GetOrdinal("first_flight_date"))),
            LastMaintenanceTime = DateTimeOffset.Parse(reader.GetString(reader.GetOrdinal("last_maintenance_time"))),
            BaseLocation = new GeoLocation(
                reader.GetDouble(reader.GetOrdinal("base_latitude")),
                reader.GetDouble(reader.GetOrdinal("base_longitude"))),
            Specs = new AircraftSpecs(
                reader.GetInt32(reader.GetOrdinal("max_speed_kmh")),
                reader.GetDouble(reader.GetOrdinal("wingspan_meters")),
                reader.GetInt32(reader.GetOrdinal("range_km")),
                reader.IsDBNull(reader.GetOrdinal("max_altitude_meters")) ? null : reader.GetInt32(reader.GetOrdinal("max_altitude_meters")),
                TimeSpan.Parse(reader.GetString(reader.GetOrdinal("flight_endurance")))),
            Conflicts = conflictList,
            Metadata = JsonSerializer.Deserialize<Dictionary<string, string>>(reader.GetString(reader.GetOrdinal("metadata")))!,
            EstimatedUnitsProduced = reader.IsDBNull(reader.GetOrdinal("estimated_units_produced")) ? null : reader.GetInt32(reader.GetOrdinal("estimated_units_produced")),
            EstimatedActiveUnits = reader.IsDBNull(reader.GetOrdinal("estimated_active_units")) ? null : reader.GetInt32(reader.GetOrdinal("estimated_active_units")),
            PhotoUrl = reader.IsDBNull(reader.GetOrdinal("photo_url")) ? null : new Uri(reader.GetString(reader.GetOrdinal("photo_url"))),
            ManualArchive = reader.IsDBNull(reader.GetOrdinal("manual_archive")) ? null : (byte[])reader["manual_archive"]
        });
    }

    return Results.Ok(aircraftList);
});

app.MapGet("/aircraft-v2/{id:guid}", (Guid id) =>
{
    using var connection = CreateConnection();
    using var queryCommand = connection.CreateCommand();
    queryCommand.CommandText = "SELECT * FROM aircraft_v2 WHERE id = @id";
    queryCommand.Parameters.AddWithValue("@id", id.ToString());

    using var reader = queryCommand.ExecuteReader();
    if (!reader.Read())
    {
        return Results.NotFound();
    }

    using var tagCommand = connection.CreateCommand();
    tagCommand.CommandText = "SELECT tag FROM aircraft_tags WHERE aircraft_id = @id";
    tagCommand.Parameters.AddWithValue("@id", id.ToString());

    var tagList = new List<string>();
    using var tagReader = tagCommand.ExecuteReader();
    while (tagReader.Read()) tagList.Add(tagReader.GetString(0));


    using var conflictsCommand = connection.CreateCommand();
    conflictsCommand.CommandText = @"SELECT name, start_year, end_year FROM aircraft_conflicts WHERE aircraft_id = @id";
    conflictsCommand.Parameters.AddWithValue("@id", id.ToString());

    var conflictList = new List<ConflictHistory>();
    using var conflictReader = conflictsCommand.ExecuteReader();
    while (conflictReader.Read())
        conflictList.Add(new ConflictHistory(conflictReader.GetString(0),
        conflictReader.GetInt32(1),
        conflictReader.GetInt32(2)));

    var aircraft = new AircraftV2
    {
        Id = id,
        Model = reader.GetString(reader.GetOrdinal("model")),
        Manufacturer = reader.GetString(reader.GetOrdinal("manufacturer")),
        SerialNumber = reader.IsDBNull(reader.GetOrdinal("serial_number")) ? null : reader.GetString(reader.GetOrdinal("serial_number")),
        YearOfManufacture = reader.GetInt32(reader.GetOrdinal("year_of_manufacture")),
        PriceMillions = decimal.Parse(reader.GetString(reader.GetOrdinal("price_millions")), CultureInfo.InvariantCulture),
        EmptyWeightKg = reader.GetDouble(reader.GetOrdinal("empty_weight_kg")),
        Status = Enum.Parse<AircraftStatus>(reader.GetString(reader.GetOrdinal("status"))),
        Role = Enum.Parse<AircraftRole>(reader.GetString(reader.GetOrdinal("role"))),
        Tags = tagList.AsReadOnly(),
        FirstFlightDate = DateOnly.Parse(reader.GetString(reader.GetOrdinal("first_flight_date"))),
        LastMaintenanceTime = DateTimeOffset.Parse(reader.GetString(reader.GetOrdinal("last_maintenance_time"))),
        BaseLocation = new GeoLocation(
            reader.GetDouble(reader.GetOrdinal("base_latitude")),
            reader.GetDouble(reader.GetOrdinal("base_longitude"))),
        Specs = new AircraftSpecs(
            reader.GetInt32(reader.GetOrdinal("max_speed_kmh")),
            reader.GetDouble(reader.GetOrdinal("wingspan_meters")),
            reader.GetInt32(reader.GetOrdinal("range_km")),
            reader.IsDBNull(reader.GetOrdinal("max_altitude_meters")) ? null : reader.GetInt32(reader.GetOrdinal("max_altitude_meters")),
            TimeSpan.Parse(reader.GetString(reader.GetOrdinal("flight_endurance")))),
        Conflicts = conflictList,
        Metadata = JsonSerializer.Deserialize<Dictionary<string, string>>(reader.GetString(reader.GetOrdinal("metadata")))!,
        EstimatedUnitsProduced = reader.IsDBNull(reader.GetOrdinal("estimated_units_produced")) ? null : reader.GetInt32(reader.GetOrdinal("estimated_units_produced")),
        EstimatedActiveUnits = reader.IsDBNull(reader.GetOrdinal("estimated_active_units")) ? null : reader.GetInt32(reader.GetOrdinal("estimated_active_units")),
        PhotoUrl = reader.IsDBNull(reader.GetOrdinal("photo_url")) ? null : new Uri(reader.GetString(reader.GetOrdinal("photo_url"))),
        ManualArchive = reader.IsDBNull(reader.GetOrdinal("manual_archive")) ? null : (byte[])reader["manual_archive"]
    };

    return Results.Ok(aircraft);
});

app.MapDelete("/aircraft-v2/{id:guid}", (Guid id) =>
{
    using var connection = CreateConnection();
    using var deleteCommand = connection.CreateCommand();
    deleteCommand.CommandText = "DELETE FROM aircraft_v2 WHERE id = @id";
    deleteCommand.Parameters.AddWithValue("@id", id.ToString());

    int rowsAffected = deleteCommand.ExecuteNonQuery();
    return rowsAffected > 0 ? Results.NoContent() : Results.NotFound();
});

app.MapPut("/aircraft-v2/{id:guid}", (Guid id, CreateAircraftV2Request req) =>
{
    using var connection = CreateConnection();
    using var transaction = connection.BeginTransaction();

    using var command = connection.CreateCommand();
    command.CommandText = @"
        UPDATE aircraft_v2 SET 
            model = @model,
            manufacturer = @manufacturer,
            serial_number = @serialNumber,
            year_of_manufacture = @yearOfManufacture,
            price_millions = @priceMillions,
            empty_weight_kg = @emptyWeightKg,
            status = @status,
            role = @role,
            first_flight_date = @firstFlightDate,
            last_maintenance_time = @lastMaintenanceTime,
            base_latitude = @baseLatitude,
            base_longitude = @baseLongitude,
            max_speed_kmh = @maxSpeedKmh,
            wingspan_meters = @wingspanMeters,
            range_km = @rangeKm,
            max_altitude_meters = @maxAltitudeMeters,
            flight_endurance = @flightEndurance,
            metadata = @metadata,
            estimated_units_produced = @estimatedUnitsProduced,
            estimated_active_units = @estimatedActiveUnits,
            photo_url = @photoUrl,
            manual_archive = @manualArchive
        WHERE id = @id";

    command.Parameters.AddWithValue("@id", id.ToString());
    command.Parameters.AddWithValue("@model", req.Model);
    command.Parameters.AddWithValue("@manufacturer", req.Manufacturer);
    command.Parameters.AddWithValue("@serialNumber", (object?)req.SerialNumber ?? DBNull.Value);
    command.Parameters.AddWithValue("@yearOfManufacture", req.YearOfManufacture);
    command.Parameters.AddWithValue("@priceMillions", req.PriceMillions.ToString(CultureInfo.InvariantCulture));
    command.Parameters.AddWithValue("@emptyWeightKg", req.EmptyWeightKg);
    command.Parameters.AddWithValue("@status", req.Status.ToString());
    command.Parameters.AddWithValue("@role", req.Role.ToString());
    command.Parameters.AddWithValue("@firstFlightDate", req.FirstFlightDate.ToString("yyyy-MM-dd"));
    command.Parameters.AddWithValue("@lastMaintenanceTime", req.LastMaintenanceTime.ToString("o"));
    command.Parameters.AddWithValue("@baseLatitude", req.BaseLocation.Latitude);
    command.Parameters.AddWithValue("@baseLongitude", req.BaseLocation.Longitude);
    command.Parameters.AddWithValue("@maxSpeedKmh", req.Specs.MaxSpeedKmh);
    command.Parameters.AddWithValue("@wingspanMeters", req.Specs.WingspanMeters);
    command.Parameters.AddWithValue("@rangeKm", req.Specs.RangeKm);
    command.Parameters.AddWithValue("@maxAltitudeMeters", (object?)req.Specs.MaxAltitudeMeters ?? DBNull.Value);
    command.Parameters.AddWithValue("@flightEndurance", req.Specs.FlightEndurance.ToString());
    command.Parameters.AddWithValue("@metadata", JsonSerializer.Serialize(req.Metadata));
    command.Parameters.AddWithValue("@estimatedUnitsProduced", (object?)req.EstimatedUnitsProduced ?? DBNull.Value);
    command.Parameters.AddWithValue("@estimatedActiveUnits", (object?)req.EstimatedActiveUnits ?? DBNull.Value);
    command.Parameters.AddWithValue("@photoUrl", (object?)req.PhotoUrl?.ToString() ?? DBNull.Value);
    command.Parameters.AddWithValue("@manualArchive", (object?)req.ManualArchive ?? DBNull.Value);

    int rowsAffected = command.ExecuteNonQuery();
    if (rowsAffected == 0)
    {
        return Results.NotFound();
    }

    // Replace tags
    using var deleteTagsCommand = connection.CreateCommand();
    deleteTagsCommand.CommandText = "DELETE FROM aircraft_tags WHERE aircraft_id = @id";
    deleteTagsCommand.Parameters.AddWithValue("@id", id.ToString());
    deleteTagsCommand.ExecuteNonQuery();

    using var tagCommand = connection.CreateCommand();
    tagCommand.CommandText = "INSERT INTO aircraft_tags (aircraft_id, tag) VALUES (@aircraftId, @tag)";
    tagCommand.Parameters.Add(new SqliteParameter("@aircraftId", id.ToString()));
    tagCommand.Parameters.Add(new SqliteParameter("@tag", ""));
    foreach (var tag in req.Tags)
    {
        tagCommand.Parameters["@tag"].Value = tag;
        tagCommand.ExecuteNonQuery();
    }

    // Replace conflicts
    using var deleteConflictsCommand = connection.CreateCommand();
    deleteConflictsCommand.CommandText = "DELETE FROM aircraft_conflicts WHERE aircraft_id = @id";
    deleteConflictsCommand.Parameters.AddWithValue("@id", id.ToString());
    deleteConflictsCommand.ExecuteNonQuery();

    using var conflictCommand = connection.CreateCommand();
    conflictCommand.CommandText = @"INSERT INTO aircraft_conflicts (aircraft_id, name, start_year, end_year)
     VALUES (@aircraftId, @name, @startYear, @endYear)";
    conflictCommand.Parameters.Add(new SqliteParameter("@aircraftId", id.ToString()));
    conflictCommand.Parameters.Add(new SqliteParameter("@name", String.Empty));
    conflictCommand.Parameters.Add(new SqliteParameter("@startYear", 0));
    conflictCommand.Parameters.Add(new SqliteParameter("@endYear", 0));
    foreach (var conflict in req.Conflicts)
    {
        conflictCommand.Parameters["@name"].Value = conflict.Name;
        conflictCommand.Parameters["@startYear"].Value = conflict.StartYear;
        conflictCommand.Parameters["@endYear"].Value = conflict.EndYear;
        conflictCommand.ExecuteNonQuery();
    }

    transaction.Commit();

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

    //aircraftStore[aircraft.Id] = aircraft;


    using var connection = CreateConnection();
    using var transaction = connection.BeginTransaction();

    using var command = connection.CreateCommand();

    command.CommandText = @"
    INSERT INTO aircraft_v2 (id, model,
     manufacturer, serial_number,
    year_of_manufacture, price_millions,
     empty_weight_kg, status,
      role, first_flight_date,
       last_maintenance_time,base_latitude,
        base_longitude, max_speed_kmh
    ,wingspan_meters, range_km, 
    max_altitude_meters, flight_endurance,
    metadata, estimated_units_produced,
     estimated_active_units, photo_url,manual_archive 
    ) VALUES (@id, @model,
     @manufacturer, @serialNumber,
      @yearOfManufacture, @priceMillions,
       @emptyWeightKg, @status,
        @role, @firstFlightDate, 
        @lastMaintenanceTime, @baseLatitude, 
            @baseLongitude, @maxSpeedKmh, 
            @wingspanMeters, @rangeKm,
            @maxAltitudeMeters, @flightEndurance,
            @metadata, @estimatedUnitsProduced, 
            @estimatedActiveUnits, @photoUrl,
             @manualArchive
            )
    ";

    command.Parameters.AddWithValue("@id", aircraft.Id.ToString());
    command.Parameters.AddWithValue("@model", aircraft.Model);
    command.Parameters.AddWithValue("@manufacturer", aircraft.Manufacturer);
    command.Parameters.AddWithValue("@serialNumber", (object?)aircraft.SerialNumber ?? DBNull.Value);
    command.Parameters.AddWithValue("@yearOfManufacture", aircraft.YearOfManufacture);
    command.Parameters.AddWithValue("@priceMillions", aircraft.PriceMillions.ToString(CultureInfo.InvariantCulture));
    command.Parameters.AddWithValue("@emptyWeightKg", aircraft.EmptyWeightKg);
    command.Parameters.AddWithValue("@status", aircraft.Status.ToString());
    command.Parameters.AddWithValue("@role", aircraft.Role.ToString());
    command.Parameters.AddWithValue("@firstFlightDate", aircraft.FirstFlightDate.ToString("yyyy-MM-dd"));
    command.Parameters.AddWithValue("@lastMaintenanceTime", aircraft.LastMaintenanceTime.ToString("o"));
    command.Parameters.AddWithValue("@baseLatitude", aircraft.BaseLocation.Latitude);
    command.Parameters.AddWithValue("@baseLongitude", aircraft.BaseLocation.Longitude);
    command.Parameters.AddWithValue("@maxSpeedKmh", aircraft.Specs.MaxSpeedKmh);
    command.Parameters.AddWithValue("@wingspanMeters", aircraft.Specs.WingspanMeters);
    command.Parameters.AddWithValue("@rangeKm", aircraft.Specs.RangeKm);
    command.Parameters.AddWithValue("@maxAltitudeMeters", (object?)aircraft.Specs.MaxAltitudeMeters ?? DBNull.Value);
    command.Parameters.AddWithValue("@flightEndurance", aircraft.Specs.FlightEndurance.ToString());
    command.Parameters.AddWithValue("@metadata", JsonSerializer.Serialize(aircraft.Metadata));
    command.Parameters.AddWithValue("@estimatedUnitsProduced", (object?)aircraft.EstimatedUnitsProduced ?? DBNull.Value);
    command.Parameters.AddWithValue("@estimatedActiveUnits", (object?)aircraft.EstimatedActiveUnits ?? DBNull.Value);
    command.Parameters.AddWithValue("@photoUrl", (object?)aircraft.PhotoUrl?.ToString() ?? DBNull.Value);
    command.Parameters.AddWithValue("@manualArchive", (object?)aircraft.ManualArchive ?? DBNull.Value);

    command.ExecuteNonQuery();
    using var tagCommand = connection.CreateCommand();
    tagCommand.CommandText = "INSERT INTO aircraft_tags (aircraft_id, tag) VALUES (@aircraftId, @tag)";
    tagCommand.Parameters.Add(new SqliteParameter("@aircraftId", aircraft.Id.ToString()));
    tagCommand.Parameters.Add(new SqliteParameter("@tag", ""));

    foreach (var tag in aircraft.Tags)
    {
        tagCommand.Parameters["@tag"].Value = tag;
        tagCommand.ExecuteNonQuery();
    }

    using var conflictCommand = connection.CreateCommand();
    conflictCommand.CommandText = @"INSERT INTO aircraft_conflicts (aircraft_id, name, start_year, end_year)
     VALUES (@aircraftId, @name, @startYear, @endYear)";
    conflictCommand.Parameters.Add(new SqliteParameter("@aircraftId", aircraft.Id.ToString()));
    conflictCommand.Parameters.Add(new SqliteParameter("@name", String.Empty));
    conflictCommand.Parameters.Add(new SqliteParameter("@startYear", 0));
    conflictCommand.Parameters.Add(new SqliteParameter("@endYear", 0));

    foreach (var conflict in aircraft.Conflicts)
    {
        conflictCommand.Parameters["@name"].Value = conflict.Name;
        conflictCommand.Parameters["@startYear"].Value = conflict.StartYear;
        conflictCommand.Parameters["@endYear"].Value = conflict.EndYear;
        conflictCommand.ExecuteNonQuery();
    }


    transaction.Commit();
    return Results.Created($"/aircraft-v2/{aircraft.Id}", aircraft);


});

SqliteConnection CreateConnection()
{
    var conn = new SqliteConnection(connectionString);
    conn.Open();
    using var pragmaFlag = conn.CreateCommand();
    pragmaFlag.CommandText = "PRAGMA foreign_keys = ON";
    pragmaFlag.ExecuteNonQuery();
    return conn;
}

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

