package main

import (
	"rest-api/database"
	"rest-api/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
