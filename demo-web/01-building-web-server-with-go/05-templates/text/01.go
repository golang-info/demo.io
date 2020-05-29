package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Result struct {
	Ok bool
	Data string
}

func main() {
	templatetext := "{{if .Ok}}Results = {{.Data}} {{else}} No Results found\n {{end}}"
	r1 := Result{Ok: true, Data: "Some data from server\n"}
	r2 := Result{Ok: false, Data: ""}
	conditionalTemplate, err := template.New("conditional").Parse(templatetext)
	if err != nil {
		log.Fatal("template parse error")
	}
	fmt.Printf("Executing first template r1\n")
	conditionalTemplate.Execute(os.Stdout, r1)
	fmt.Printf("Executing second template r2\n")
	conditionalTemplate.Execute(os.Stdout, r2)

}
