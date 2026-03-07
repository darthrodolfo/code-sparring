package main

import (
	"fmt"
	"log"
	"net/http"
)

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

	log.Println("server running on :8080")

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatalf("failed to start server: %v", error)
	}
}
