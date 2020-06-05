package main

import (
	"net/http"
)

func showQuery(w http.ResponseWriter, r *http.Request) {
	querystring := r.URL.Query()
	w.Write([]byte("query strings are\n"))
	w.Write([]byte("Name: " + querystring.Get("name") + "\n"))
	w.Write([]byte("Email: " + querystring.Get("email") + "\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", showQuery)

	http.ListenAndServe(":3000", mux)
}