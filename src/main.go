package main

import (
	"fmt"
	"os"

	"github.com/ivancduran/imgserver/src/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	err := godotenv.Load("../env.dev")
	err = godotenv.Load("env")
	if err != nil {
		fmt.Println("Error loading env file")
	}

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	// v1
	a := e.Group("/v1")
	routes.ApiRoutes(a)

	// db init
	// StartDB()

	// start server
	e.Logger.Fatal(e.Start(":8090"))
}

func StartDB() {
	// models.AutoMigrate()
}
