package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func RequestFuture(url string) <- chan []byte {
	c := make(chan []byte, 1)
	go func() {
		var body []byte
		defer func() {
			c <- body
		}()

		res, err := http.Get(url)
		if err != nil {
			return
		}
		defer res.Body.Close()

		body, _ = ioutil.ReadAll(res.Body)
	}()

	return c
}

func main() {
	future := RequestFuture("https://api.github.com/users/octocat/orgs")
	body := <-future
	log.Printf("response longth: %d", len(body))
}