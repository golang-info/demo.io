package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (f *formHandler) handleForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	file := "form.html"
	res, err := template.ParseFiles(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = res.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (f *formHandler) createSurvey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	name := r.PostForm.Get("firstName")
	feedback := r.PostForm.Get("feedback")
	id := f.addFeedback(name, feedback)

	http.Redirect(w, r, fmt.Sprintf("/survey/%d", id), http.StatusSeeOther)
}

func (f *formHandler) showSurvey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	survey, err := f.getFeedback(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "Survey with id= %d created successfully\n", survey.id)
	fmt.Fprintf(w, "username=%s, feedback=%s", survey.name, survey.feedback)

}

func (f *formHandler) allSurveys(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, f.surveys)
}


