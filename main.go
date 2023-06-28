package main

import (
	"os"
	"sesi6/configs"
	"sesi6/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// loadEnv()
	configs.InitDB()
	e := echo.New()

	e = routes.Routes(e)
	e.Start(":" + getPort()) // adding localhost for prevent windows firewall popup

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }
