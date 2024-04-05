// main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct untuk data yang akan digunakan
type Data struct {
	Message string `json:"message"`
}

// Handler untuk endpoint pertama
func endpointOneHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Message: "Hello from Endpoint One!",
	}

	// Mengubah struct menjadi JSON dan mengirimkannya sebagai response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Handler untuk endpoint kedua
func endpointTwoHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Message: "Hello from Endpoint Two!",
	}

	// Mengubah struct menjadi JSON dan mengirimkannya sebagai response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Menghubungkan handler dengan endpoint
	http.HandleFunc("/endpoint1", endpointOneHandler)
	http.HandleFunc("/endpoint2", endpointTwoHandler)

	// Menjalankan server di port 8080
	port := 8080
	fmt.Printf("Server running on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
