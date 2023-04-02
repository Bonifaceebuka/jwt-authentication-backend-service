package main

import (
	"os"

	"github.com/Bonifaceebuka/jwt-authentication-backend-service/config"
	"github.com/Bonifaceebuka/jwt-authentication-backend-service/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()

	config.LoadEnv()
	port := os.Getenv("PORT")
	routes.Router(router)
	router.Listen(":" + port)

}
