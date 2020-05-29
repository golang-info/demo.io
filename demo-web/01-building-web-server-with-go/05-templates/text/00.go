package main

import (
	"os"
	"text/template"
)

type Informateion struct {
	Username string
	Useremail string
	Seat int
}

func main() {
	texttemplate := `{{.Username}} is allocated seat {{.Seat}}.
					{{.Username}} contected at {{.Useremail}}\n`
	info := Informateion{
		Username: "sumit",
		Useremail: "gophersumit@gmail.com",
		Seat: 21,
	}
	InfoTemplate, err := template.New("information").Parse(texttemplate)
	if err != nil {
		panic(err)
	}
	InfoTemplate.Execute(os.Stdout, info)
}