package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// for prod use secure key, net hard-coded
var store = sessions.NewCookieStore([]byte("1234"))

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "custom-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val := session.Values["hello"]
	if val != nil {
		fmt.Fprintln(w, "retrieving existing session: ")
		fmt.Fprintln(w, val)
		return
	}
	session.Values["Hello"] = "world"
	session.Save(r, w)
	fmt.Fprintln(w, "no existing session found, set value for session")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sessionHandler)
	http.ListenAndServe(":3000", mux)
}


