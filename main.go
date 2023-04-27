package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define la estructura para la solicitud de sumar dos números
type AddRequest struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

// Define la estructura para la respuesta de la suma de dos números
type AddResponse struct {
	Result int `json:"result"`
}

// Función para sumar dos números
func AddNumbers(w http.ResponseWriter, r *http.Request) {
	// Decodifica la solicitud JSON y la almacena en la variable "request"
	var request AddRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Realiza la suma de los dos números
	result := request.Num1 + request.Num2

	// Codifica la respuesta en JSON y la envía como respuesta
	response := AddResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func ShowAddNumbers(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "You are going to add two numbers!")

	// Decodifica la solicitud JSON y la almacena en la variable "request"
}
func main() {
	// Crea un nuevo enrutador "mux"
	r := mux.NewRouter()

	// Configura la ruta y el controlador para la función AddNumbers
	r.HandleFunc("/add", AddNumbers).Methods("POST")
	r.HandleFunc("/add", ShowAddNumbers).Methods("GET")
	// Inicia el servidor en el puerto 8000
	log.Fatal(http.ListenAndServe(":8000", r))
}
