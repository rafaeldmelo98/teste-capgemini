package main

import (
	"teste-capgemini/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/stats", handlers.Stats)
	e.POST("/sequence", handlers.CheckSequence)
	e.Logger.Fatal(e.Start(":8080"))
}
