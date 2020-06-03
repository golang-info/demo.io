package main

import (
	"encoding/json"
	"net/http"
)

type Address struct {
	Street string `json:"street, emitempty"`
	Landmark string `json:"landmark, omitempty"`
	Pincode int `json:"pincode, omitempty"`
}

func sendJSON(w http.ResponseWriter, r *http.Request) {
	a := Address{
		Street: "Viman Nagar",
		Landmark: "Nexa",
		Pincode: 411014,
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(a)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sendJSON)
	http.ListenAndServe(":3000", mux)
}
