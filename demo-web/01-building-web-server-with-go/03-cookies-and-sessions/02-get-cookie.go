package main

import (
	"fmt"
	"net/http"
)

func getCookies(w http.ResponseWriter, r *http.Request) {
	// get all cookies
	cookies := r.Cookies()
	for _, cookie := range cookies {
		fmt.Fprintln(w, cookie)
	}
	// get named cookie
	cookie, err := r.Cookie("cookie-1")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
	fmt.Fprintln(w, cookie)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getCookies)
	http.ListenAndServe(":3000", mux)
}
