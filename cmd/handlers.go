package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func handleHealth(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	// This is a blind yes/no response.
	err := checkHealth(ctx, db)
	return err
}

func returnJSON(w http.ResponseWriter, item interface{}) error {

	resp, err := json.Marshal(item)
	if err != nil {
		return errors.Wrap(err, "Error converting response to JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return nil
}

func handleUsers(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	users, err := getUsers(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching users from PGBouncer")
	}
	return returnJSON(w, users)
}

func handleConfigs(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	configs, err := getConfigs(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching configs from PGBouncer")
	}
	return returnJSON(w, configs)
}

func handleDatabases(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	databases, err := getDatabases(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching databases from PGBouncer")
	}
	return returnJSON(w, databases)
}

func handlePools(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	pools, err := getPools(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching pools from PGBouncer")
	}
	return returnJSON(w, pools)
}

func handleClients(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	clients, err := getClients(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching clients from PGBouncer")
	}
	return returnJSON(w, clients)
}

func handleServers(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	servers, err := getServers(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching servers from PGBouncer")
	}
	return returnJSON(w, servers)
}

func handleMems(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	mems, err := getMems(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching mems from PGBouncer")
	}
	return returnJSON(w, mems)
}

func handleDmesg(w http.ResponseWriter) error {
	output, err := getDmesg()
	if err != nil {
		return errors.Wrap(err, "Error fetching kernel logs")
	}
	w.Write(output)
	return nil
}

func handleProcesses(w http.ResponseWriter) error {
	output, err := getProcesses()
	if err != nil {
		return errors.Wrap(err, "Error fetching process list")
	}
	w.Write(output)
	return nil
}

func handleMeminfo(w http.ResponseWriter) error {
	output, err := getMeminfo()
	if err != nil {
		return errors.Wrap(err, "Error fetching memory info")
	}
	w.Write(output)
	return nil
}

func handleLogs(w http.ResponseWriter) error {
	output, err := getLogs()
	if err != nil {
		return errors.Wrap(err, "Error fetching logs")
	}
	w.Write(output)
	return nil
}
