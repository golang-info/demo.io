package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	hello := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("This is version V1 !!!")
		fmt.Println("ok")
		fmt.Fprintln(w, "This is version V1 \n")
	}

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":31003", nil))
}
