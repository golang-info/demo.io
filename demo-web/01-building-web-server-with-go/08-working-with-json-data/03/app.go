package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Address struct {
	Street string `json:"street,omitempty"`
	Landmark string `json:"landmark,omitempty"`
	Pincode int `json:"pincode,omitempty"`
}

func acceptJSON(w http.ResponseWriter, r *http.Request) {
	a := Address{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&a)
	if err != nil {
		fmt.Fprintf(w, "Error parsing json")
		return
	}
	fmt.Fprintln(w, "received :", &a)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", acceptJSON)
	http.ListenAndServe(":3000", mux)
}
