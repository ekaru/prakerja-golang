package main

import (
	"log"
	"sesi6/configs"
	"sesi6/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	configs.InitDB()
	e := echo.New()

	e = routes.Routes(e)
	e.Start("localhost:8000") // adding localhost for prevent windows firewall popup

}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
