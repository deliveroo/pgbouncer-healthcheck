package main

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

// User as returned by PGBouncer's SHOW USERS
type User struct {
	Name     string
	PoolMode *string
}

// Config item as returned by PGBouncer's SHOW CONFIG
type Config struct {
	Key        string
	Value      *string
	Changeable bool
}

// Database as returned by PGBouncer's SHOW DATABASES
type Database struct {
	Name               string
	Host               *string
	Port               string
	Database           string
	ForceUser          *string
	PoolSize           int
	ReservePool        int
	PoolMode           *string
	MaxConnections     int
	CurrentConnections int
}

// Pool as returned by PGBouncer's SHOW POOLS
type Pool struct {
	Database  string
	User      string
	ClActive  int
	ClWaiting int
	SvActive  int
	SvIdle    int
	SvUsed    int
	SvTested  int
	SvLogin   int
	MaxWait   int
	PoolMode  string
}

// Client as returned by PGBouncer's SHOW CLIENTS
type Client struct {
	Type        string
	User        string
	Database    string
	State       string
	Addr        string
	Port        int
	LocalAddr   string
	LocalPort   int
	ConnectTime string
	RequestTime string
	Ptr         string
	Link        string
	RemotePid   int
	TLS         string
}

// Server as returned by PGBouncer's SHOW SERVERS
type Server struct {
	Type        string
	User        string
	Database    string
	State       string
	Addr        string
	Port        int
	LocalAddr   string
	LocalPort   int
	ConnectTime string
	RequestTime string
	Ptr         string
	Link        string
	RemotePid   int
	TLS         string
}

// Mem info record as returned by PGBouncer's SHOW MEM
type Mem struct {
	Name     string
	Size     int
	Used     int
	Free     int
	MemTotal int
}

// Stat  record as returned by PGBouncer's SHOW STAT
type Stat struct {
	Database       string
	TotalRequests  int
	TotalReceived  int
	TotalSent      int
	TotalQueryTime int
	AvgReq         int
	AvgRecv        int
	AvgSent        int
	AvgQuery       int
}

func unwrapNullString(in sql.NullString) *string {
	if in.Valid {
		return &in.String
	}
	return nil
}

func getUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	rows, err := db.QueryContext(ctx, "SHOW USERS")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var rawPoolMode sql.NullString

		if err := rows.Scan(&user.Name, &rawPoolMode); err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		user.PoolMode = unwrapNullString(rawPoolMode)
		users = append(users, user)
	}
	return users, nil
}

func getConfigs(ctx context.Context, db *sql.DB) ([]Config, error) {
	rows, err := db.QueryContext(ctx, "SHOW CONFIG")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var configs []Config
	for rows.Next() {
		var config Config
		var rawValue sql.NullString
		var rawChangeable string

		if err := rows.Scan(&config.Key, &rawValue, &rawChangeable); err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		config.Changeable = rawChangeable == "yes"
		config.Value = unwrapNullString(rawValue)
		configs = append(configs, config)
	}
	return configs, nil
}

func getDatabases(ctx context.Context, db *sql.DB) ([]Database, error) {
	rows, err := db.QueryContext(ctx, "SHOW DATABASES")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var databases []Database
	for rows.Next() {
		var database Database
		var rawHost sql.NullString
		var rawForceUser sql.NullString
		var rawPoolMode sql.NullString

		err := rows.Scan(
			&database.Name,
			&rawHost,
			&database.Port,
			&database.Database,
			&rawForceUser,
			&database.PoolSize,
			&database.ReservePool,
			&rawPoolMode,
			&database.MaxConnections,
			&database.CurrentConnections,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		database.Host = unwrapNullString(rawHost)
		database.ForceUser = unwrapNullString(rawForceUser)
		database.PoolMode = unwrapNullString(rawPoolMode)
		databases = append(databases, database)
	}
	return databases, nil
}

func getPools(ctx context.Context, db *sql.DB) ([]Pool, error) {
	rows, err := db.QueryContext(ctx, "SHOW POOLS")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var pools []Pool
	for rows.Next() {
		var pool Pool

		err := rows.Scan(
			&pool.Database,
			&pool.User,
			&pool.ClActive,
			&pool.ClWaiting,
			&pool.SvActive,
			&pool.SvIdle,
			&pool.SvUsed,
			&pool.SvTested,
			&pool.SvLogin,
			&pool.MaxWait,
			&pool.PoolMode,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}

func getClients(ctx context.Context, db *sql.DB) ([]Client, error) {
	rows, err := db.QueryContext(ctx, "SHOW CLIENTS")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client

		err := rows.Scan(
			&client.Type,
			&client.User,
			&client.Database,
			&client.State,
			&client.Addr,
			&client.Port,
			&client.LocalAddr,
			&client.LocalPort,
			&client.ConnectTime,
			&client.RequestTime,
			&client.Ptr,
			&client.Link,
			&client.RemotePid,
			&client.TLS,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func getServers(ctx context.Context, db *sql.DB) ([]Server, error) {
	rows, err := db.QueryContext(ctx, "SHOW SERVERS")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var servers []Server
	for rows.Next() {
		var server Server

		err := rows.Scan(
			&server.Type,
			&server.User,
			&server.Database,
			&server.State,
			&server.Addr,
			&server.Port,
			&server.LocalAddr,
			&server.LocalPort,
			&server.ConnectTime,
			&server.RequestTime,
			&server.Ptr,
			&server.Link,
			&server.RemotePid,
			&server.TLS,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		servers = append(servers, server)
	}
	return servers, nil
}

func getMems(ctx context.Context, db *sql.DB) ([]Mem, error) {
	rows, err := db.QueryContext(ctx, "SHOW MEM")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var mems []Mem
	for rows.Next() {
		var mem Mem

		err := rows.Scan(
			&mem.Name,
			&mem.Size,
			&mem.Used,
			&mem.Free,
			&mem.MemTotal,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		mems = append(mems, mem)
	}
	return mems, nil
}

func getStats(ctx context.Context, db *sql.DB) ([]Stat, error) {
	rows, err := db.QueryContext(ctx, "SHOW STATS")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query PGBouncer")
	}
	defer rows.Close()

	var stats []Stat
	for rows.Next() {
		var stat Stat

		err := rows.Scan(
			&stat.Database,
			&stat.TotalRequests,
			&stat.TotalReceived,
			&stat.TotalSent,
			&stat.TotalQueryTime,
			&stat.AvgReq,
			&stat.AvgRecv,
			&stat.AvgSent,
			&stat.AvgQuery,
		)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		stats = append(stats, stat)
	}
	return stats, nil
}
