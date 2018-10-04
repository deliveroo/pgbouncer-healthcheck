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

func handleUsers(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	users, err := getUsers(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching users from PGBouncer")
	}
	resp, err := json.Marshal(users)
	if err != nil {
		return errors.Wrap(err, "Error converting users to JSON")
	}
	w.Write(resp)
	return nil
}

func handleConfigs(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	configs, err := getConfigs(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching configs from PGBouncer")
	}
	resp, err := json.Marshal(configs)
	if err != nil {
		return errors.Wrap(err, "Error converting configs to JSON")
	}
	w.Write(resp)
	return nil
}

func handleDatabases(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	databases, err := getDatabases(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching databases from PGBouncer")
	}
	resp, err := json.Marshal(databases)
	if err != nil {
		return errors.Wrap(err, "Error converting databases to JSON")
	}
	w.Write(resp)
	return nil
}

func handlePools(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	pools, err := getPools(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching pools from PGBouncer")
	}
	resp, err := json.Marshal(pools)
	if err != nil {
		return errors.Wrap(err, "Error converting pools to JSON")
	}
	w.Write(resp)
	return nil
}

func handleClients(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	clients, err := getClients(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching clients from PGBouncer")
	}
	resp, err := json.Marshal(clients)
	if err != nil {
		return errors.Wrap(err, "Error converting clients to JSON")
	}
	w.Write(resp)
	return nil
}

func handleServers(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	servers, err := getServers(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching servers from PGBouncer")
	}
	resp, err := json.Marshal(servers)
	if err != nil {
		return errors.Wrap(err, "Error converting servers to JSON")
	}
	w.Write(resp)
	return nil
}

func handleMems(ctx context.Context, w http.ResponseWriter, db *sql.DB) error {
	mems, err := getMems(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching mems from PGBouncer")
	}
	resp, err := json.Marshal(mems)
	if err != nil {
		return errors.Wrap(err, "Error converting mems to JSON")
	}
	w.Write(resp)
	return nil
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
