package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Aircraft struct {
	ID           int    `json:"id"`
	Code         string `json:"code"`
	Manufacturer string `json:"manufacturer"`
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

// var aircrafts = []Aircraft{
// 	{
// 		ID:           1,
// 		Code:         "A320",
// 		Manufacturer: "Airbus",
// 	},
// 	{
// 		ID:           2,
// 		Code:         "B737",
// 		Manufacturer: "Boeing",
// 	},
// }

var aircrafts = []AircraftV2{
	{
		ID:           1,
		Model:        "A320",
		Manufacturer: "Airbus",
	},
}

func aircraftV2CollectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listAircraftV2Handler(w, r)
	case http.MethodPost:
		writeError(w, http.StatusNotImplemented, "not implemented", "createAircraftV2Handler pending")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func aircraftV2ItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAircraftV2ByIDHandler(w, r)
	case http.MethodPut:
		writeError(w, http.StatusNotImplemented, "not implemented", "updateAircraftV2Handler pending")
	case http.MethodDelete:
		writeError(w, http.StatusNotImplemented, "not implemented", "deleteAircraftV2Handler pending")
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
	return uuid.Parse(rawID)
}

var aircraftV2Store = NewAircraftV2Store()

type AircraftV2Store struct {
	mu   sync.RWMutex
	data map[uuid.UUID]AircraftV2
}

func NewAircraftV2Store() *AircraftV2Store {
	return &AircraftV2Store{
		data: make(map[uuid.UUID]AircraftV2),
	}
}

func (s *AircraftV2Store) List() []AircraftV2 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]AircraftV2, 0, len(s.data))
	for aircraft := range s.data {
		result = append(result, aircraft)
	}
	return result
}

func (s *AircraftV2Store) Get(id uuid.UUID) (AircraftV2, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	aircraft, ok := s.data[id]
	return aircraft, ok
}

func (s *AircraftV2Store) Create(aircraft AircraftV2) AircraftV2 {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[aircraft.ID] = aircraft
	return aircraft
}

func (s *AircraftV2Store) Update(id uuid.UUID, aircraft AircraftV2) (AircraftV2, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[id]; !exists {
		return AircraftV2{}, false
	}

	aircraft.ID = id
	s.data[id] = aircraft
	return aircraft, true
}

func (s *AircraftV2Store) Delete(id uuid.UUID) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.data[id]; !exists {
		return false
	}

	delete(s.data, id)
	return true
}

var nextAircraftId = 3

func listAircraftV2Handler(w http.ResponseWriter, r *http.Request) {
	items := aircraftV2Store.List()
	writeJSON(w, http.StatusOK, items)
}

func getAircraftV2ByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := parseAircraftV2IDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	item, ok := aircraftV2Store.Get(id)
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
		//http.Error(writer, `{"err": "invalid request body"}`, http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"err": "invalid request body"})
		return
	}

	createRequest.Code = strings.TrimSpace(createRequest.Code)
	createRequest.Manufacturer = strings.TrimSpace(createRequest.Manufacturer)

	if createRequest.Code == "" {
		//http.Error(writer, `{"err": "code is required"}`, http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"err": "code is required"})
		return
	}

	if createRequest.Manufacturer == "" {
		//http.Error(writer, `{"err": "manufacturer is required"}`, http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"err": "manufacturer is required"})
		return
	}

	newAircraft := AircraftV2{
		ID:           nextAircraftId,
		Model:        createRequest.Code,
		Manufacturer: createRequest.Manufacturer,
	}

	aircrafts = append(aircrafts, newAircraft)
	nextAircraftId++

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(newAircraft)
	if err != nil {
		log.Printf("failed to encode created aircraft: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func decolamosHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//_, err := fmt.Fprint(writer, "Decolamos")
	err := json.NewEncoder(writer).Encode(map[string]string{"message": "Decolamos"})
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	//fmt.Println("Hello, World!")
	http.HandleFunc("/decolamos", decolamosHandler)
	http.HandleFunc("/aircraft", aircraftHandler)

	http.HandleFunc("/aircraft-v2", aircraftV2CollectionHandler) // GET, POST
	http.HandleFunc("/aircraft-v2/", aircraftV2ItemHandler)      // GET, PUT, DELETE com id

	log.Println("server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
