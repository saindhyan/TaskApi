package main

import (
	"task/controller"
	"task/database"
	"task/service"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//Database connection
	db, _ := database.Connect()
	// service connection
	service.Connect(db)

}

func main() {
	// running api endpints on port
	controller.Routes().Run(":4949")

}
