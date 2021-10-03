package main

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/sonadztux/tgtc-project-gogql/backend/database"
	"github.com/sonadztux/tgtc-project-gogql/backend/server"
)

func main(){
	// Init database connection
	database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http
	// router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}