package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var (
	config configuration
	db     *sql.DB
)

func initServer() *http.Server {
	router := httprouter.New()

	// Add a default root 200 handler
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {})

	// Add a version endpoint
	router.GET("/ami-version", makeRequestHandlerFile("version info", "/etc/ami_version"))

	addHealthHandlers(router)
	addStatusHandlers(router)
	if config.EnableDebugEndpoints {
		log.Print("Enabling Debug Endpoints")
		addDebugHandlers(router)
	}
	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", config.Port),
	}
}

func initDB() {
	var err error
	db, err = sql.Open("postgres", config.Connstr)
	if err == nil {
		log.Printf("Connected to PGBouncer database")
	} else {
		db = nil
		log.Printf("Could not open database: %s", err)
	}
}

func main() {
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("Could not process configuration: %s", err)
	}
	initDB()
	webserver := initServer()
	log.Printf("Listening on port %d", config.Port)
	webserver.ListenAndServe()
}
