package main

import (
	"task/controller"
	"task/database"
	"task/service"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

	db, _ := database.Connect()
	
	service.Connect(db)

}
func main() {
	controller.Routes().Run(":4949")

}
