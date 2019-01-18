package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

func addStatusHandlers(router *httprouter.Router) {
	router.GET("/status/users", makeRequestHandlerWithContext(handleUsers))
	router.GET("/status/configs", makeRequestHandlerWithContext(handleConfigs))
	router.GET("/status/databases", makeRequestHandlerWithContext(handleDatabases))
	router.GET("/status/pools", makeRequestHandlerWithContext(handlePools))
	router.GET("/status/clients", makeRequestHandlerWithContext(handleClients))
	router.GET("/status/servers", makeRequestHandlerWithContext(handleServers))
	router.GET("/status/memory", makeRequestHandlerWithContext(handleMems))
	router.GET("/status/stats", makeRequestHandlerWithContext(handleStats))
}

func handleUsers(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	users, err := getUsers(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching users from PGBouncer")
	}
	return returnJSON(w, users)
}

func handleConfigs(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	configs, err := getConfigs(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching configs from PGBouncer")
	}
	return returnJSON(w, configs)
}

func handleDatabases(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	databases, err := getDatabases(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching databases from PGBouncer")
	}
	return returnJSON(w, databases)
}

func handlePools(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	pools, err := getPools(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching pools from PGBouncer")
	}
	return returnJSON(w, pools)
}

func handleClients(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	clients, err := getClients(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching clients from PGBouncer")
	}
	return returnJSON(w, clients)
}

func handleServers(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	servers, err := getServers(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching servers from PGBouncer")
	}
	return returnJSON(w, servers)
}

func handleMems(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	mems, err := getMems(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching mems from PGBouncer")
	}
	return returnJSON(w, mems)
}

func handleStats(ctx context.Context, w http.ResponseWriter, _ httprouter.Params) error {
	if db == nil {
		return errors.New("PGBouncer Database is not connected")
	}
	stats, err := getStats(ctx, db)
	if err != nil {
		return errors.Wrap(err, "Error fetching stats from PGBouncer")
	}
	return returnJSON(w, stats)
}
