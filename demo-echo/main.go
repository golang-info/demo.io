package main

import (
	"github.com/labstack/echo/v4"

	_ "demo.io/demo-echo/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"demo.io/demo-echo/handler"
	"demo.io/demo-echo/utils"
	"demo.io/demo-echo/common"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// 初始化k8s客户端
	if err := common.InitClient(); err != nil {
		panic(err.Error())
	}

	e := echo.New()

	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	e.Use(utils.CustomeMiddleware)
	e.Use(utils.CustomeResponse)

	e.GET("swagger/*", echoSwagger.WrapHandler)

	g := e.Group("/v3")
	handler.InitRouter(g)

	//e.POST("/postData", postData.PostData)

	e.Logger.Fatal(e.Start(":1321"))
}
