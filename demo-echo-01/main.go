package main

/*
https://www.restapiexample.com/golang-tutorial/consume-restful-apis-using-echo-golang/
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Employee []struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

func main() {
	e := echo.New()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/employee", func(c echo.Context) error {
		// Build the request
		req, err := http.NewRequest("GET", "http://dummy.restapiexample.com/api/v1/employees", nil)
		if err != nil {
			fmt.Println("Error is req: ", err)
		}

		// create a Client
		client := &http.Client{}

		// Do sends an HTTP request and
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error in send req: ", err)
		}

		// Defer the closing of the body
		defer resp.Body.Close()

		// Fill the data with the data from the JSON
		var data Employee

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			fmt.Println(err)
		}

		return c.JSON(http.StatusOK, data)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
