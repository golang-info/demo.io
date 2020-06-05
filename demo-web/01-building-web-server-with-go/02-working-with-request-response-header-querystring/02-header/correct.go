package main

import "net/http"

func setHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad request!\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", setHeader)

	http.ListenAndServe(":3000", mux)
}
