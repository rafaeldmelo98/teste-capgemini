package main

import (
	"database/sql"
	"fmt"
	"log"
	"teste-capgemini/database"
	"teste-capgemini/handlers"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Starting API")
	database.SetUpDatabase()
	db, err := sql.Open("sqlite3", "sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	handler := handlers.Handler{
		DB: db,
	}
	e := echo.New()
	e.GET("/stats", handler.Stats)
	e.POST("/sequence", handler.CheckSequence)
	e.Logger.Fatal(e.Start(":8080"))
}
