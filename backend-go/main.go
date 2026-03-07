package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"strings"
	"time"
)

var db *sql.DB

type Aircraft struct {
	ID           uuid.UUID `json:"id"`
	Code         string    `json:"code"`
	Manufacturer string    `json:"manufacturer"`
}

type CreateAircraftRequest struct {
	Code         string `json:"code"`
	Manufacturer string `json:"manufacturer"`
}

type AircraftRole string

const (
	RoleFighter        AircraftRole = "Fighter"
	RoleBomber         AircraftRole = "Bomber"
	RoleTransport      AircraftRole = "Transport"
	RoleTrainer        AircraftRole = "Trainer"
	RoleDrone          AircraftRole = "Drone"
	RoleReconnaissance AircraftRole = "Reconnaissance"
)

type AircraftStatus string

const (
	StatusActive      AircraftStatus = "Active"
	StatusMaintenance AircraftStatus = "Maintenance"
	StatusRetired     AircraftStatus = "Retired"
	StatusStored      AircraftStatus = "Stored"
)

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AircraftSpecs struct {
	MaxSpeedKmh       int           `json:"max_speed_kmh"`
	WingspanMeters    float64       `json:"wingspan_meters"`
	RangeKm           int           `json:"range_km"`
	MaxAltitudeMeters *int          `json:"max_altitude_meters"`
	FlightEndurance   time.Duration `json:"flight_endurance"`
}

type ConflictHistory struct {
	Name      string `json:"name"`
	StartYear int    `json:"start_year"`
	EndYear   int    `json:"end_year"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

type AircraftV2 struct {
	ID                     uuid.UUID         `json:"id"`
	Model                  string            `json:"model"`
	Manufacturer           string            `json:"manufacturer"`
	SerialNumber           *string           `json:"serial_number"`
	YearOfManufacture      int               `json:"year_of_manufacture"`
	PriceMillions          decimal.Decimal   `json:"price_millions"`
	EmptyWeightKg          float64           `json:"empty_weight_kg"`
	Status                 AircraftStatus    `json:"status"`
	Role                   AircraftRole      `json:"role"`
	Tags                   []string          `json:"tags"`
	FirstFlightDate        time.Time         `json:"first_flight_date"`
	LastMaintenanceTime    time.Time         `json:"last_maintenance_time"`
	BaseLocation           GeoLocation       `json:"base_location"`
	Specs                  AircraftSpecs     `json:"specs"`
	Conflicts              []ConflictHistory `json:"conflicts"`
	Metadata               map[string]string `json:"metadata"`
	EstimatedUnitsProduced *int              `json:"estimated_units_produced"`
	EstimatedActiveUnits   *int              `json:"estimated_active_units"`
	PhotoUrl               *string           `json:"photo_url"`
	ManualArchive          []byte            `json:"manual_archive"`
}

type CreateAircraftV2Request struct {
	Model                  string            `json:"model"`
	Manufacturer           string            `json:"manufacturer"`
	SerialNumber           *string           `json:"serial_number"`
	YearOfManufacture      int               `json:"year_of_manufacture"`
	PriceMillions          decimal.Decimal   `json:"price_millions"`
	EmptyWeightKg          float64           `json:"empty_weight_kg"`
	Status                 AircraftStatus    `json:"status"`
	Role                   AircraftRole      `json:"role"`
	Tags                   []string          `json:"tags"`
	FirstFlightDate        time.Time         `json:"first_flight_date"`
	LastMaintenanceTime    time.Time         `json:"last_maintenance_time"`
	BaseLocation           GeoLocation       `json:"base_location"`
	Specs                  AircraftSpecs     `json:"specs"`
	Conflicts              []ConflictHistory `json:"conflicts"`
	Metadata               map[string]string `json:"metadata"`
	EstimatedUnitsProduced *int              `json:"estimated_units_produced"`
	EstimatedActiveUnits   *int              `json:"estimated_active_units"`
	PhotoUrl               *string           `json:"photo_url"`
	ManualArchive          []byte            `json:"manual_archive"`
}

func aircraftV2CollectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listAircraftV2Handler(w, r)
	case http.MethodPost:
		createAircraftV2Handler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func aircraftV2ItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAircraftV2ByIDHandler(w, r)
	case http.MethodPut:
		updateAircraftV2Handler(w, r)
	case http.MethodDelete:
		deleteAircraftV2Handler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to write json response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string, details string) {
	resp := ErrorResponse{
		Error:   message,
		Details: details,
	}
	writeJSON(w, status, resp)
}

func decodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}

func parseAircraftV2IDFromPath(path string) (uuid.UUID, error) {
	const prefix = "/aircraft-v2/"
	rawID := strings.TrimPrefix(path, prefix)
	if rawID == "" || rawID == path {
		return uuid.Nil, fmt.Errorf("missing aircraft id in path")
	}
	return uuid.Parse(rawID)
}

func listAircraftV2Handler(w http.ResponseWriter, r *http.Request) {
	items, err := listAircraftV2FromDB(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list aircraft", err.Error())
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func getAircraftV2ByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseAircraftV2IDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	item, ok, err := getAircraftV2ByIDFromDB(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch aircraft", err.Error())
		return
	}
	if !ok {
		writeError(w, http.StatusNotFound, "aircraft not found", "")
		return
	}

	writeJSON(w, http.StatusOK, item)
}

func aircraftHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getAircraftHandler(writer, request)
	case http.MethodPost:
		createAircraftHandler(writer, request)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getAircraftHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(aircrafts)
	if err != nil {
		log.Printf("failed to encode aircraft list: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func createAircraftHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var createRequest CreateAircraftRequest

	err := json.NewDecoder(request.Body).Decode(&createRequest)
	if err != nil {
		json.NewEncoder(writer).Encode(map[string]string{"err": "invalid request body"})
		return
	}

	createRequest.Code = strings.TrimSpace(createRequest.Code)
	createRequest.Manufacturer = strings.TrimSpace(createRequest.Manufacturer)

	if createRequest.Code == "" {
		json.NewEncoder(writer).Encode(map[string]string{"err": "code is required"})
		return
	}

	if createRequest.Manufacturer == "" {
		json.NewEncoder(writer).Encode(map[string]string{"err": "manufacturer is required"})
		return
	}

	newAircraft := Aircraft{
		ID:           uuid.New(),
		Code:         createRequest.Code,
		Manufacturer: createRequest.Manufacturer,
	}

	aircrafts = append(aircrafts, newAircraft)

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(newAircraft)
	if err != nil {
		log.Printf("failed to encode created aircraft: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func mapCreateRequestToAircraftV2(req CreateAircraftV2Request) AircraftV2 {
	return AircraftV2{
		ID:                     uuid.New(),
		Model:                  strings.TrimSpace(req.Model),
		Manufacturer:           strings.TrimSpace(req.Manufacturer),
		SerialNumber:           req.SerialNumber,
		YearOfManufacture:      req.YearOfManufacture,
		PriceMillions:          req.PriceMillions,
		EmptyWeightKg:          req.EmptyWeightKg,
		Status:                 req.Status,
		Role:                   req.Role,
		Tags:                   req.Tags,
		FirstFlightDate:        req.FirstFlightDate,
		LastMaintenanceTime:    req.LastMaintenanceTime,
		BaseLocation:           req.BaseLocation,
		Specs:                  req.Specs,
		Conflicts:              req.Conflicts,
		Metadata:               req.Metadata,
		EstimatedUnitsProduced: req.EstimatedUnitsProduced,
		EstimatedActiveUnits:   req.EstimatedActiveUnits,
		PhotoUrl:               req.PhotoUrl,
		ManualArchive:          req.ManualArchive,
	}
}

func createAircraftV2Handler(w http.ResponseWriter, r *http.Request) {
	var req CreateAircraftV2Request

	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	req.Model = strings.TrimSpace(req.Model)
	req.Manufacturer = strings.TrimSpace(req.Manufacturer)

	if err := validateCreateAircraftV2Request(&req); err != nil {
		writeError(w, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	entity := mapCreateRequestToAircraftV2(req)
	if err := createAircraftV2InDB(r.Context(), entity); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create aircraft", err.Error())
		return
	}

	w.Header().Set("Location", "/aircraft-v2/"+entity.ID.String())
	writeJSON(w, http.StatusCreated, entity)
}

func decolamosHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := json.NewEncoder(writer).Encode(map[string]string{"message": "Decolamos"})
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func updateAircraftV2Handler(w http.ResponseWriter, r *http.Request) {
	id, err := parseAircraftV2IDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	var req CreateAircraftV2Request

	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	req.Model = strings.TrimSpace(req.Model)
	req.Manufacturer = strings.TrimSpace(req.Manufacturer)

	if err := validateCreateAircraftV2Request(&req); err != nil {
		writeError(w, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	entity := mapCreateRequestToAircraftV2(req)
	updated, ok, err := updateAircraftV2InDB(r.Context(), id, entity)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update aircraft", err.Error())
		return
	}
	if !ok {
		writeError(w, http.StatusNotFound, "aircraft not found", "")
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func deleteAircraftV2Handler(w http.ResponseWriter, r *http.Request) {
	id, err := parseAircraftV2IDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	ok, err := deleteAircraftV2InDB(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete aircraft", err.Error())
		return
	}
	if !ok {
		writeError(w, http.StatusNotFound, "aircraft not found", "")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func isValidAircraftStatus(v AircraftStatus) bool {
	switch v {
	case StatusActive, StatusMaintenance, StatusRetired, StatusStored:
		return true
	default:
		return false
	}
}

func isValidAircraftRole(v AircraftRole) bool {
	switch v {
	case RoleFighter, RoleBomber, RoleTransport, RoleTrainer, RoleDrone, RoleReconnaissance:
		return true
	default:
		return false
	}
}

func validateCreateAircraftV2Request(req *CreateAircraftV2Request) error {
	req.Model = strings.TrimSpace(req.Model)
	req.Manufacturer = strings.TrimSpace(req.Manufacturer)

	if strings.TrimSpace(req.Model) == "" {
		return fmt.Errorf("model is required")
	}

	if len(req.Model) > 80 {
		return fmt.Errorf("model must be <= 80 chars")
	}

	if req.Manufacturer == "" {
		return fmt.Errorf("manufacturer is required")
	}
	if len(req.Manufacturer) > 80 {
		return fmt.Errorf("manufacturer must be <= 80 chars")
	}

	currentYear := time.Now().Year()
	if req.YearOfManufacture < 1903 || req.YearOfManufacture > currentYear+1 {
		return fmt.Errorf("year_of_manufacture must be between 1903 and %d", currentYear+1)
	}

	if isValidAircraftStatus(req.Status) == false {
		return fmt.Errorf("invalid status")
	}
	if isValidAircraftRole(req.Role) == false {
		return fmt.Errorf("invalid role")
	}

	if req.EmptyWeightKg <= 0 {
		return fmt.Errorf("empty_weight_kg must be > 0")
	}

	if req.Specs.MaxSpeedKmh <= 0 {
		return fmt.Errorf("specs.max_speed_kmh must be > 0")
	}
	if req.Specs.WingspanMeters <= 0 {
		return fmt.Errorf("specs.wingspan_meters must be > 0")
	}
	if req.Specs.RangeKm <= 0 {
		return fmt.Errorf("specs.range_km must be > 0")
	}
	if req.Specs.MaxAltitudeMeters != nil && *req.Specs.MaxAltitudeMeters <= 0 {
		return fmt.Errorf("specs.max_altitude_meters must be > 0 when provided")
	}
	if req.Specs.FlightEndurance <= 0 {
		return fmt.Errorf("specs.flight_endurance must be > 0")
	}

	normalizedTags, err := normalizeTags(req.Tags)
	if err != nil {
		return err
	}
	req.Tags = normalizedTags

	return nil
}

func normalizeTags(tags []string) ([]string, error) {
	if len(tags) == 0 {
		return []string{}, nil
	}

	seen := make(map[string]struct{})
	result := make([]string, 0, len(tags))

	for _, t := range tags {
		trimmed := strings.TrimSpace(t)
		if trimmed == "" {
			continue
		}
		if len(trimmed) > 24 {
			return nil, fmt.Errorf("tag '%s' exceeds 24 chars", trimmed)
		}

		key := strings.ToLower(trimmed)
		if _, exists := seen[key]; exists {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, trimmed)
	}

	if len(result) > 10 {
		return nil, fmt.Errorf("tags cannot exceed 10 items")
	}

	return result, nil
}

func initSQLite(dataSource string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, err
	}

	if _, err := database.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		_ = database.Close()
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS aircraft_v2 (
    id TEXT PRIMARY KEY,
    model TEXT NOT NULL,
    manufacturer TEXT NOT NULL,
    serial_number TEXT,
    year_of_manufacture INTEGER NOT NULL,
    price_millions TEXT NOT NULL,
    empty_weight_kg REAL NOT NULL,
    status TEXT NOT NULL,
    role TEXT NOT NULL,
    first_flight_date TEXT NOT NULL,
    last_maintenance_time TEXT NOT NULL,
    base_latitude REAL NOT NULL,
    base_longitude REAL NOT NULL,
    max_speed_kmh INTEGER NOT NULL,
    wingspan_meters REAL NOT NULL,
    range_km INTEGER NOT NULL,
    max_altitude_meters INTEGER,
    flight_endurance TEXT NOT NULL,
    metadata TEXT NOT NULL,
    estimated_units_produced INTEGER,
    estimated_active_units INTEGER,
    photo_url TEXT,
    manual_archive BLOB
);

CREATE TABLE IF NOT EXISTS aircraft_tags (
    aircraft_id TEXT NOT NULL,
    tag TEXT NOT NULL,
    PRIMARY KEY (aircraft_id, tag),
    FOREIGN KEY (aircraft_id) REFERENCES aircraft_v2(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS aircraft_conflicts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    aircraft_id TEXT NOT NULL,
    name TEXT NOT NULL,
    start_year INTEGER NOT NULL,
    end_year INTEGER NOT NULL,
    FOREIGN KEY (aircraft_id) REFERENCES aircraft_v2(id) ON DELETE CASCADE
);
`
	if _, err := database.Exec(schema); err != nil {
		_ = database.Close()
		return nil, err
	}

	return database, nil
}

func listAircraftV2FromDB(ctx context.Context) ([]AircraftV2, error) {
	rows, err := db.QueryContext(ctx, `
SELECT id, model, manufacturer, serial_number, year_of_manufacture, price_millions, empty_weight_kg,
       status, role, first_flight_date, last_maintenance_time, base_latitude, base_longitude,
       max_speed_kmh, wingspan_meters, range_km, max_altitude_meters, flight_endurance,
       metadata, estimated_units_produced, estimated_active_units, photo_url, manual_archive
FROM aircraft_v2`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]AircraftV2, 0)
	for rows.Next() {
		aircraft, err := scanAircraftV2Row(rows)
		if err != nil {
			return nil, err
		}
		aircraft.Tags, err = getTagsByAircraftID(ctx, aircraft.ID)
		if err != nil {
			return nil, err
		}
		aircraft.Conflicts, err = getConflictsByAircraftID(ctx, aircraft.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, aircraft)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func getAircraftV2ByIDFromDB(ctx context.Context, id uuid.UUID) (AircraftV2, bool, error) {
	row := db.QueryRowContext(ctx, `
SELECT id, model, manufacturer, serial_number, year_of_manufacture, price_millions, empty_weight_kg,
       status, role, first_flight_date, last_maintenance_time, base_latitude, base_longitude,
       max_speed_kmh, wingspan_meters, range_km, max_altitude_meters, flight_endurance,
       metadata, estimated_units_produced, estimated_active_units, photo_url, manual_archive
FROM aircraft_v2 WHERE id = ?`, id.String())

	aircraft, err := scanAircraftV2Row(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return AircraftV2{}, false, nil
		}
		return AircraftV2{}, false, err
	}

	aircraft.Tags, err = getTagsByAircraftID(ctx, id)
	if err != nil {
		return AircraftV2{}, false, err
	}
	aircraft.Conflicts, err = getConflictsByAircraftID(ctx, id)
	if err != nil {
		return AircraftV2{}, false, err
	}

	return aircraft, true, nil
}

func createAircraftV2InDB(ctx context.Context, aircraft AircraftV2) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	metadataJSON, err := json.Marshal(aircraft.Metadata)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
INSERT INTO aircraft_v2 (
    id, model, manufacturer, serial_number, year_of_manufacture, price_millions, empty_weight_kg,
    status, role, first_flight_date, last_maintenance_time, base_latitude, base_longitude,
    max_speed_kmh, wingspan_meters, range_km, max_altitude_meters, flight_endurance, metadata,
    estimated_units_produced, estimated_active_units, photo_url, manual_archive
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		aircraft.ID.String(),
		aircraft.Model,
		aircraft.Manufacturer,
		aircraft.SerialNumber,
		aircraft.YearOfManufacture,
		aircraft.PriceMillions.String(),
		aircraft.EmptyWeightKg,
		string(aircraft.Status),
		string(aircraft.Role),
		aircraft.FirstFlightDate.Format(time.RFC3339),
		aircraft.LastMaintenanceTime.Format(time.RFC3339),
		aircraft.BaseLocation.Latitude,
		aircraft.BaseLocation.Longitude,
		aircraft.Specs.MaxSpeedKmh,
		aircraft.Specs.WingspanMeters,
		aircraft.Specs.RangeKm,
		aircraft.Specs.MaxAltitudeMeters,
		aircraft.Specs.FlightEndurance.String(),
		string(metadataJSON),
		aircraft.EstimatedUnitsProduced,
		aircraft.EstimatedActiveUnits,
		aircraft.PhotoUrl,
		aircraft.ManualArchive,
	)
	if err != nil {
		return err
	}

	if err := replaceTagsTx(ctx, tx, aircraft.ID, aircraft.Tags); err != nil {
		return err
	}
	if err := replaceConflictsTx(ctx, tx, aircraft.ID, aircraft.Conflicts); err != nil {
		return err
	}

	return tx.Commit()
}

func updateAircraftV2InDB(ctx context.Context, id uuid.UUID, aircraft AircraftV2) (AircraftV2, bool, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return AircraftV2{}, false, err
	}
	defer tx.Rollback()

	metadataJSON, err := json.Marshal(aircraft.Metadata)
	if err != nil {
		return AircraftV2{}, false, err
	}

	result, err := tx.ExecContext(ctx, `
UPDATE aircraft_v2 SET
    model = ?, manufacturer = ?, serial_number = ?, year_of_manufacture = ?, price_millions = ?, empty_weight_kg = ?,
    status = ?, role = ?, first_flight_date = ?, last_maintenance_time = ?, base_latitude = ?, base_longitude = ?,
    max_speed_kmh = ?, wingspan_meters = ?, range_km = ?, max_altitude_meters = ?, flight_endurance = ?, metadata = ?,
    estimated_units_produced = ?, estimated_active_units = ?, photo_url = ?, manual_archive = ?
WHERE id = ?`,
		aircraft.Model,
		aircraft.Manufacturer,
		aircraft.SerialNumber,
		aircraft.YearOfManufacture,
		aircraft.PriceMillions.String(),
		aircraft.EmptyWeightKg,
		string(aircraft.Status),
		string(aircraft.Role),
		aircraft.FirstFlightDate.Format(time.RFC3339),
		aircraft.LastMaintenanceTime.Format(time.RFC3339),
		aircraft.BaseLocation.Latitude,
		aircraft.BaseLocation.Longitude,
		aircraft.Specs.MaxSpeedKmh,
		aircraft.Specs.WingspanMeters,
		aircraft.Specs.RangeKm,
		aircraft.Specs.MaxAltitudeMeters,
		aircraft.Specs.FlightEndurance.String(),
		string(metadataJSON),
		aircraft.EstimatedUnitsProduced,
		aircraft.EstimatedActiveUnits,
		aircraft.PhotoUrl,
		aircraft.ManualArchive,
		id.String(),
	)
	if err != nil {
		return AircraftV2{}, false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return AircraftV2{}, false, err
	}
	if rowsAffected == 0 {
		return AircraftV2{}, false, nil
	}

	if err := replaceTagsTx(ctx, tx, id, aircraft.Tags); err != nil {
		return AircraftV2{}, false, err
	}
	if err := replaceConflictsTx(ctx, tx, id, aircraft.Conflicts); err != nil {
		return AircraftV2{}, false, err
	}

	if err := tx.Commit(); err != nil {
		return AircraftV2{}, false, err
	}

	aircraft.ID = id
	return aircraft, true, nil
}

func deleteAircraftV2InDB(ctx context.Context, id uuid.UUID) (bool, error) {
	result, err := db.ExecContext(ctx, `DELETE FROM aircraft_v2 WHERE id = ?`, id.String())
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func replaceTagsTx(ctx context.Context, tx *sql.Tx, aircraftID uuid.UUID, tags []string) error {
	if _, err := tx.ExecContext(ctx, `DELETE FROM aircraft_tags WHERE aircraft_id = ?`, aircraftID.String()); err != nil {
		return err
	}
	for _, tag := range tags {
		if _, err := tx.ExecContext(ctx, `INSERT INTO aircraft_tags (aircraft_id, tag) VALUES (?, ?)`, aircraftID.String(), tag); err != nil {
			return err
		}
	}
	return nil
}

func replaceConflictsTx(ctx context.Context, tx *sql.Tx, aircraftID uuid.UUID, conflicts []ConflictHistory) error {
	if _, err := tx.ExecContext(ctx, `DELETE FROM aircraft_conflicts WHERE aircraft_id = ?`, aircraftID.String()); err != nil {
		return err
	}
	for _, c := range conflicts {
		if _, err := tx.ExecContext(ctx, `INSERT INTO aircraft_conflicts (aircraft_id, name, start_year, end_year) VALUES (?, ?, ?, ?)`,
			aircraftID.String(), c.Name, c.StartYear, c.EndYear); err != nil {
			return err
		}
	}
	return nil
}

func getTagsByAircraftID(ctx context.Context, id uuid.UUID) ([]string, error) {
	rows, err := db.QueryContext(ctx, `SELECT tag FROM aircraft_tags WHERE aircraft_id = ?`, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]string, 0)
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, rows.Err()
}

func getConflictsByAircraftID(ctx context.Context, id uuid.UUID) ([]ConflictHistory, error) {
	rows, err := db.QueryContext(ctx, `SELECT name, start_year, end_year FROM aircraft_conflicts WHERE aircraft_id = ?`, id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	conflicts := make([]ConflictHistory, 0)
	for rows.Next() {
		var c ConflictHistory
		if err := rows.Scan(&c.Name, &c.StartYear, &c.EndYear); err != nil {
			return nil, err
		}
		conflicts = append(conflicts, c)
	}
	return conflicts, rows.Err()
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanAircraftV2Row(scanner rowScanner) (AircraftV2, error) {
	var (
		idStr                  string
		model                  string
		manufacturer           string
		serialNumber           sql.NullString
		yearOfManufacture      int
		priceMillionsStr       string
		emptyWeightKg          float64
		statusStr              string
		roleStr                string
		firstFlightRaw         string
		lastMaintenanceRaw     string
		baseLatitude           float64
		baseLongitude          float64
		maxSpeedKmh            int
		wingspanMeters         float64
		rangeKm                int
		maxAltitudeMeters      sql.NullInt64
		flightEnduranceRaw     string
		metadataRaw            string
		estimatedUnitsProduced sql.NullInt64
		estimatedActiveUnits   sql.NullInt64
		photoURL               sql.NullString
		manualArchive          []byte
	)

	if err := scanner.Scan(
		&idStr,
		&model,
		&manufacturer,
		&serialNumber,
		&yearOfManufacture,
		&priceMillionsStr,
		&emptyWeightKg,
		&statusStr,
		&roleStr,
		&firstFlightRaw,
		&lastMaintenanceRaw,
		&baseLatitude,
		&baseLongitude,
		&maxSpeedKmh,
		&wingspanMeters,
		&rangeKm,
		&maxAltitudeMeters,
		&flightEnduranceRaw,
		&metadataRaw,
		&estimatedUnitsProduced,
		&estimatedActiveUnits,
		&photoURL,
		&manualArchive,
	); err != nil {
		return AircraftV2{}, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return AircraftV2{}, err
	}
	price, err := decimal.NewFromString(priceMillionsStr)
	if err != nil {
		return AircraftV2{}, err
	}
	firstFlight, err := parseFlexibleDateTime(firstFlightRaw)
	if err != nil {
		return AircraftV2{}, err
	}
	lastMaintenance, err := parseFlexibleDateTime(lastMaintenanceRaw)
	if err != nil {
		return AircraftV2{}, err
	}
	flightEndurance, err := parseFlexibleDuration(flightEnduranceRaw)
	if err != nil {
		return AircraftV2{}, err
	}

	metadata := map[string]string{}
	if strings.TrimSpace(metadataRaw) != "" {
		if err := json.Unmarshal([]byte(metadataRaw), &metadata); err != nil {
			return AircraftV2{}, err
		}
	}

	var serialNumberPtr *string
	if serialNumber.Valid {
		serialNumberValue := serialNumber.String
		serialNumberPtr = &serialNumberValue
	}

	var maxAltitudePtr *int
	if maxAltitudeMeters.Valid {
		v := int(maxAltitudeMeters.Int64)
		maxAltitudePtr = &v
	}

	var estimatedUnitsProducedPtr *int
	if estimatedUnitsProduced.Valid {
		v := int(estimatedUnitsProduced.Int64)
		estimatedUnitsProducedPtr = &v
	}

	var estimatedActiveUnitsPtr *int
	if estimatedActiveUnits.Valid {
		v := int(estimatedActiveUnits.Int64)
		estimatedActiveUnitsPtr = &v
	}

	var photoURLPtr *string
	if photoURL.Valid {
		v := photoURL.String
		photoURLPtr = &v
	}

	return AircraftV2{
		ID:           id,
		Model:        model,
		Manufacturer: manufacturer,
		SerialNumber: serialNumberPtr,
		YearOfManufacture: yearOfManufacture,
		PriceMillions:     price,
		EmptyWeightKg:     emptyWeightKg,
		Status:            AircraftStatus(statusStr),
		Role:              AircraftRole(roleStr),
		Tags:              []string{},
		FirstFlightDate:   firstFlight,
		LastMaintenanceTime: lastMaintenance,
		BaseLocation: GeoLocation{
			Latitude:  baseLatitude,
			Longitude: baseLongitude,
		},
		Specs: AircraftSpecs{
			MaxSpeedKmh:       maxSpeedKmh,
			WingspanMeters:    wingspanMeters,
			RangeKm:           rangeKm,
			MaxAltitudeMeters: maxAltitudePtr,
			FlightEndurance:   flightEndurance,
		},
		Conflicts:              []ConflictHistory{},
		Metadata:               metadata,
		EstimatedUnitsProduced: estimatedUnitsProducedPtr,
		EstimatedActiveUnits:   estimatedActiveUnitsPtr,
		PhotoUrl:               photoURLPtr,
		ManualArchive:          manualArchive,
	}, nil
}

func parseFlexibleDateTime(raw string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339, raw); err == nil {
		return t, nil
	}
	if t, err := time.Parse("2006-01-02", raw); err == nil {
		return t, nil
	}
	return time.Time{}, fmt.Errorf("unsupported datetime format: %s", raw)
}

func parseFlexibleDuration(raw string) (time.Duration, error) {
	if d, err := time.ParseDuration(raw); err == nil {
		return d, nil
	}

	parts := strings.Split(raw, ":")
	if len(parts) == 3 {
		var hh, mm, ss int
		if _, err := fmt.Sscanf(raw, "%d:%d:%d", &hh, &mm, &ss); err == nil {
			return time.Duration(hh)*time.Hour + time.Duration(mm)*time.Minute + time.Duration(ss)*time.Second, nil
		}
	}

	return 0, fmt.Errorf("unsupported duration format: %s", raw)
}

var aircrafts = []Aircraft{
	{
		ID:           uuid.New(),
		Code:         "A320",
		Manufacturer: "Airbus",
	},
	{
		ID:           uuid.New(),
		Code:         "B737",
		Manufacturer: "Boeing",
	},
}

func main() {

	var err error
	db, err = initSQLite("aerostack.db")
	if err != nil {
		log.Fatalf("failed to init sqlite: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/decolamos", decolamosHandler)
	http.HandleFunc("/aircraft", aircraftHandler)

	http.HandleFunc("/aircraft-v2", aircraftV2CollectionHandler)
	http.HandleFunc("/aircraft-v2/", aircraftV2ItemHandler)

	log.Println("server running on :8080")

	httpErr := http.ListenAndServe(":8080", nil)
	if httpErr != nil {
		log.Fatalf("failed to start server: %v", httpErr)
	}
}
