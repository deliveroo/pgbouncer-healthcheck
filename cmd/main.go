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

type requestHandler func(context.Context, http.ResponseWriter, *sql.DB) error

func makeHandler(db *sql.DB, h requestHandler) http.HandlerFunc {
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

func initServer() *http.Server {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", makeHandler(db, handleHealth))
	mux.HandleFunc("/users", makeHandler(db, handleUsers))
	mux.HandleFunc("/configs", makeHandler(db, handleConfigs))
	mux.HandleFunc("/databases", makeHandler(db, handleDatabases))
	mux.HandleFunc("/pools", makeHandler(db, handlePools))

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
