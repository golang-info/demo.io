package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sendResponse)
	server := http.Server{
		Addr: ":3000",
		Handler: customMiddleware(mux),
	}
	server.ListenAndServe()
}
