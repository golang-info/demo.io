package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	forms := &formHandler{
		surveys: []survey{},
	}
	router := httprouter.New()
	router.GET("/", forms.handleForm)
	router.POST("/create/survey", forms.createSurvey)
	router.GET("/survey/:id", forms.showSurvey)
	router.GET("/surveys", forms.allSurveys)

	http.ListenAndServe(":3000", router)
}
