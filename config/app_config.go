package config

import (
	"fmt"
	"log"

	// "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	databaseConnection *gorm.DB
)

const (
	username = "root"
	password = ""
	hostname = "localhost:3306"
	dbname   = "go_book_store_db"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
func connectDB() *gorm.DB {

	dbConnection, err := gorm.Open("mysql", dsn(dbname))

	// handle error
	if err != nil {
		panic(err)
	}

	// err = dbConnection.Ping()
	databaseConnection = dbConnection

	// err := dbConnection.AutoMigrate(&Book{})

	// if err != nil {
	// 	fmt.Print("Unable to migrate DB tables")
	// }
	// defer dbConnection.Close()
	if err != nil {
		panic("Unable to open database")
	}
	return databaseConnection
}

func LoadEnv() {
	err := godotenv.Load(".env")
	// err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: unable to load the .env file")
	}
}

func GetDBConnection() *gorm.DB {
	connectDB()
	return databaseConnection
}
