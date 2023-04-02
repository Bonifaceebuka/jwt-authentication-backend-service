package routes

import (
	controllers "github.com/Bonifaceebuka/jwt-authentication-backend-service/controllers"
	"github.com/Bonifaceebuka/jwt-authentication-backend-service/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Router(router *fiber.App) {

	router.Get("/", controllers.Home)
	router.Post("/signup", controllers.Signup)
	router.Post("/signin", controllers.Signin)
	router.Use(middlewares.IsAuthenticate)
}
