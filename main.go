package main

import (
	"github.com/namelew/gin-api/database"
	"github.com/namelew/gin-api/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
