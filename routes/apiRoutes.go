package routes

import (
	controllers "github.com/Bonifaceebuka/book-store-RESTful-API-in-GO/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/book/store", controllers.StoreBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/book/{book_id}", controllers.GetBook).Methods("GET", "OPTIONS")
	router.HandleFunc("/book/{book_id}", controllers.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods("DELETE", "OPTIONS")

	return router
}
