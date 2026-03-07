package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"strings"
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

var aircrafts = []Aircraft{
	{
		ID:           1,
		Code:         "A320",
		Manufacturer: "Airbus",
	},
	{
		ID:           2,
		Code:         "B737",
		Manufacturer: "Boeing",
	},
}

var nextAircraftId = 3

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

	newAircraft := Aircraft{
		ID:           nextAircraftId,
		Code:         createRequest.Code,
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

	log.Println("server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
