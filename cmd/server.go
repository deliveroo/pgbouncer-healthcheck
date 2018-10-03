package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func checkHealth(ctx context.Context, db *sql.DB) error {
	_, err := getUsers(ctx, db)
	return errors.Wrap(err, "PGbouncer health check failed")
}

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
