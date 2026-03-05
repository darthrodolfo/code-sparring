var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", () => "Hello AeroStack!");

app.MapGet("/decolamos", () => "Decolamos!");

app.MapGet("/aircraft", () => new List<AircraftV1>());


app.Run();

record AircraftV1(Guid Id, string Model, string Manufacturer, int Year);

record AircraftV2
{
    public required Guid Id { get; init; }
    public required string Model { get; init; }
    public required string Manufacturer { get; init; }
    public required int Year { get; init; }
}


