package handler

import (
	"github.com/labstack/echo/v4"
	"demo.io/demo-echo/handler/helloWorld"
	"demo.io/demo-echo/handler/pods"
)

func InitRouter(e *echo.Group) {
	helloWorld.InitRouter(e.Group("/demo"))
	pods.InitRouter(e.Group("/pods"))
}