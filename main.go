package main

import (
	"Project/database"
	"Project/router"
	"os"
)

func main() {
	database.ConnectDB()
	var PORT = os.Getenv("PORT")
	r := router.StartApp()
	r.Run(":" + PORT)
}