package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name string
}

func main() {
	loopText := `{{range .}} Found {{.Name}}{{"\n"}}{{end}}`
	allUsers := []User{
		{Name: "sumit"},
		{Name: "deepti"},
		{Name: "amit"},
	}
	loopTemplate, err := template.New("logp").Parse(loopText)
	if err != nil {
		log.Fatal("error parsing template")
	}
	loopTemplate.Execute(os.Stdout, allUsers)
}
