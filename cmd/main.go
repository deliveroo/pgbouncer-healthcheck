package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const port = ":8000"
const connStr = "host=localhost port=6543 dbname=pgbouncer sslmode=disable"

type requestHandlerWithDB func(context.Context, http.ResponseWriter, *sql.DB) error
type requestHandlerSimple func(http.ResponseWriter) error

func makeHandlerWithDB(db *sql.DB, h requestHandlerWithDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		err := h(ctx, w, db)
		if err != nil {
			errMsg := fmt.Sprintf("An error occured: %s", err)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
	}
}

func makeHandlerSimple(h requestHandlerSimple) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w)
		if err != nil {
			errMsg := fmt.Sprintf("An error occured: %s", err)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
	}
}

func initServer() *http.Server {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", makeHandlerWithDB(db, handleHealth))
	mux.HandleFunc("/users", makeHandlerWithDB(db, handleUsers))
	mux.HandleFunc("/configs", makeHandlerWithDB(db, handleConfigs))
	mux.HandleFunc("/databases", makeHandlerWithDB(db, handleDatabases))
	mux.HandleFunc("/pools", makeHandlerWithDB(db, handlePools))
	mux.HandleFunc("/clients", makeHandlerWithDB(db, handleClients))
	mux.HandleFunc("/servers", makeHandlerWithDB(db, handleServers))
	mux.HandleFunc("/mems", makeHandlerWithDB(db, handleMems))

	mux.HandleFunc("/dmesg", makeHandlerSimple(handleDmesg))
	mux.HandleFunc("/processes", makeHandlerSimple(handleProcesses))

	return &http.Server{
		Handler: mux,
		Addr:    port,
	}
}

func main() {
	webserver := initServer()

	fmt.Println("Listening on port 8000")
	webserver.ListenAndServe()
}
