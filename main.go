package main

import (
	"log"
	"os"
	"project-pertama/database"
	"project-pertama/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")		
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":"+port)
}