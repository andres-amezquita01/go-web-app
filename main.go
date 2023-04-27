package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AddRequest struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

type AddResponse struct {
	Result int `json:"result"`
}

func AddNumbers(w http.ResponseWriter, r *http.Request) {
	// Decodifica la solicitud JSON y la almacena en la variable "request"
	var request AddRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := request.Num1 + request.Num2

	response := AddResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func ShowAddNumbers(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Now, you are going to add two numbers!")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/add", AddNumbers).Methods("POST")
	r.HandleFunc("/add", ShowAddNumbers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
