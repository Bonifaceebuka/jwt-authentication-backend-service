package routes

import (
	controllers "github.com/Bonifaceebuka/jwt-authentication-backend-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(router *fiber.App) {

	router.Get("/", controllers.Home)
	router.Post("/signup", controllers.Signup)
	router.Post("/signin", controllers.Signin)
	// router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET", "OPTIONS")
	// router.HandleFunc("/book/{book_id}", controllers.GetBook).Methods("GET", "OPTIONS")
	// router.HandleFunc("/book/{book_id}", controllers.UpdateBook).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods("DELETE", "OPTIONS")

	// return router
}
