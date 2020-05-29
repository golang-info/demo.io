package main

import (
	"html/template"
	"net/http"
)

func sedTemplate(w http.ResponseWriter, r *http.Request) {
	homeTemplate, err := template.ParseFiles("home.html", "header.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	msg := Message{HomeData: "Gophers", HeaderData: "Some header content"}
	// execute template
	w.Header().Set("Content-type", "text/html")
	err = homeTemplate.Execute(w, msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
