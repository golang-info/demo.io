package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	hed := r.Header
	fmt.Println(w, hed)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", headers)

	http.ListenAndServe(":3000", mux)
}
