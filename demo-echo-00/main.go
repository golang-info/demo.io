package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/labstack/echo"
)

type Album struct {
	ID   int    `jsonapi:"primary,albums"`
	Name string `jsonapi:"attr,name"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var b bytes.Buffer
		err := jsonapi.MarshalPayload(&b, albumList())
		if err != nil {
			fmt.Println(err)
		}
		return c.JSONBlob(http.StatusOK, b.Bytes())
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func albumList() []*Album {
	a1 := Album{123, "allbum1"}
	a2 := Album{456, "allbum2"}
	albums := []*Album{&a1, &a2}
	return albums
}
