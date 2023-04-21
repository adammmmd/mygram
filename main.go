package main

import (
	"Project/database"
	"Project/router"
	"os"
)

func main() {
	database.ConnectDB()
	r := router.StartApp()
	r.Run(":8080")
}