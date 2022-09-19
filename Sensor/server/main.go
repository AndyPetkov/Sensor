package main

import (
	"fmt"
	"server/controller"
	"server/logger"

	_ "server/database"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	logger.GetInstance().InfoLogger.Println("Starting the application...")
	controller.HandleRequests()
}
