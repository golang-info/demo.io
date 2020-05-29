package main

import (
	"net/http"
)

type Message struct { HomeData, HeaderData string}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sedTemplate)
	http.ListenAndServe(":3000", mux)
}
