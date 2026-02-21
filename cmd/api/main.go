package main

import (
	"github.com/Turgho/GoFlowDesk/internal/database"
	"github.com/Turgho/GoFlowDesk/internal/router"
)

// @title GoFlowDesk API
// @version 1.0
// @description This is the API documentation for GoFlowDesk, a workflow management system built with Go and PostgreSQL.
// @contact.name API Support
// @contact.url http://www.goflowdesk.com/support
// @contact.email
func main() {
	// Set up the database connection
	dbConnection := database.SetupDatabase()
	defer dbConnection.Close()

	// Set up the router and start the server
	router := router.SetupRouter(dbConnection)
	router.Run(":8080")
}
