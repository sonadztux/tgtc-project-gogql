package main

import (
	"time"

	"github.com/sonadztux/tgtc-project-gogql/backend/controller"
	"github.com/sonadztux/tgtc-project-gogql/backend/database"
	"github.com/sonadztux/tgtc-project-gogql/backend/server"
)

func main() {
	// Init database connection
	database.InitDB()

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, controller.Register())
}
