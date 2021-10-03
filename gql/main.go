package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sonadztux/tgtc-project-gogql/backend/server"
	"github.com/sonadztux/tgtc-project-gogql/gql/gqlserver"
)

func main() {
	router := mux.NewRouter()

	couponResolver := gqlserver.NewResolver()
	schemaWrapper := gqlserver.NewSchemaWrapper().
		WithCouponResolver(couponResolver)

	if err := schemaWrapper.Init(); err != nil {
		log.Fatal("unable to parse schema, err: ", err.Error())
	}

	// GraphQL Handler
	router.Path("/graphql").Handler(gqlserver.NewHandler(schemaWrapper).Handle())

	// this assumes main.go is called from root project,
	// change this accordingly, if it's called elsewhere
	fs := http.FileServer(http.Dir("gql/webstatic"))
	router.PathPrefix("/").Handler(fs)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         9000,
	}
	server.Serve(serverConfig, router)
}
