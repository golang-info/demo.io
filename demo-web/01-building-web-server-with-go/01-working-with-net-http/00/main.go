package main

import "net/http"

func main() {
	registerRoutes()
	httpServer := http.Server{
		Addr: ":3000",
		Handler: mux,
	}
	httpServer.ListenAndServe()
}
