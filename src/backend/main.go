package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person struct to hold the JSON data
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register a handler function for the path
	r.HandleFunc("/", PersonHandler).Methods("POST", "OPTIONS")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8081", r))
}

// PersonHandler handles requests to the /person path
func PersonHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	// Handle preflight request (CORS)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Decode the JSON request body into a Person struct
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Print the result to the console
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received name and age"))
}
