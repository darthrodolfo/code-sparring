package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func aicraftHandler(writer http.ResponseWriter, request *http.Request) {
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

	error := json.NewEncoder(writer).Encode(aircrafts)
	if error != nil {
		log.Printf("failed to encode aircraft list: %v", error)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func createAircraftHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var createRequest CreateAircraftRequest

	error := json.NewDecoder(request.Body).Decode(&createRequest)
	if error != nil {
		http.Error(writer, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}

	createRequest.Code = strings.TrimSpace(createRequest.Code)
	createRequest.Manufacturer = strings.TrimSpace(createRequest.Manufacturer)

	if createRequest.Code == "" {
		http.Error(writer, `{"error": "code is required"}`, http.StatusBadRequest)
		return
	}

	if createRequest.Manufacturer == "" {
		http.Error(writer, `{"error": "manufacturer is required"}`, http.StatusBadRequest)
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

	error = json.NewEncoder(writer).Encode(newAircraft)
	if error != nil {
		log.Printf("failed to encode created aircraft: %v", error)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func decolamosHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_, error := fmt.Fprint(writer, "Decolamos")
	if error != nil {
		log.Printf("failed to write response: %v", error)
	}
}

func main() {
	//fmt.Println("Hello, World!")
	http.HandleFunc("/decolamos", decolamosHandler)
	http.HandleFunc("/aircraft", aicraftHandler)

	log.Println("server running on :8080")

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatalf("failed to start server: %v", error)
	}
}
