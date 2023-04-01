package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Bonifaceebuka/book-store-RESTful-API-in-GO/config"
	"github.com/Bonifaceebuka/book-store-RESTful-API-in-GO/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var DBConnection *gorm.DB
var BookModel models.Book

func init() {
	DBConnection = config.GetDBConnection()
	config.LoadEnv()
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-METHODS", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal("API service is fully up!")
	w.Write(response)
	fmt.Println(string(response))
}

func StoreBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()
	// r.ParseMultipartForm()
	json.NewDecoder(r.Body).Decode(&BookModel)

	name := r.FormValue("name")
	author := r.FormValue("author")
	publication := r.FormValue("publication")

	book := models.Book{
		Name:        name,
		Author:      author,
		Publication: publication,
	}

	DBConnection.Create(&book)

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal("Book added successfully!")
	w.Write(response)
}

func GetAllBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET")

	books := models.GetAllBooks()

	json.NewEncoder(res).Encode(books)
}

func GetBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Methods", "GET")
	res.Header().Set("Access-Control-Allow-Orign", "*")
	res.Header().Set("Access-Control-Allow-Type", "application/json")

	vars := mux.Vars(req)
	book_id := vars["book_id"]

	id, _ := strconv.Atoi(book_id)

	foundBook := DBConnection.Model(&BookModel).Where("id=?", id).Preload("Book").First(&BookModel)

	fmt.Println(foundBook)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Methods", "PUT")
	res.Header().Set("Access-Control-Allow-Orign", "*")
	res.Header().Set("Access-Control-Allow-Type", "application/json")

	vars := mux.Vars(req)
	book_id := vars["book_id"]

	id, _ := strconv.Atoi(book_id)

	foundBook := DBConnection.Model(&BookModel).Where("id=?", id).Preload("Book").First(&BookModel)
	if foundBook == nil {
		fmt.Print("No matching Book found")
		return
	}

	book := models.Book{
		Id: uint(id),
	}
	req.ParseForm()

	name := req.FormValue("name")
	author := req.FormValue("author")
	publication := req.FormValue("publication")

	DBConnection.Model(&book).Where("id=?", id).Updates(
		map[string]interface{}{
			"name":        name,
			"author":      author,
			"publication": publication,
		})

	fmt.Println("Book updated successfully")
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Methods", "DELETE")
	res.Header().Set("Access-Control-Allow-Orign", "*")
	res.Header().Set("Access-Control-Allow-Type", "application/json")

	vars := mux.Vars(req)
	book_id := vars["book_id"]

	id, _ := strconv.Atoi(book_id)

	foundBook := DBConnection.Model(&BookModel).Where("id=?", id).Preload("Book").First(&BookModel)
	if foundBook == nil {
		fmt.Print("No matching Book found")
		return
	}

	book := models.Book{
		Id: uint(id),
	}
	DBConnection.Model(&book).Where("id=?", id).Delete(&book)
	fmt.Println("Book deleted")
}
