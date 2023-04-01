package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "github.com/Bonifaceebuka/jwt-authentication-backend-service/routes"
)

func main() {
	router := router.Router()
	port := os.Getenv("PORT")
	port = ":" + port
	fmt.Println("Starting the server on port %s", port)

	log.Fatal(http.ListenAndServe(port, router))

	fmt.Println("Server started on port %s", port)

}
