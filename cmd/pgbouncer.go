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
		var poolMode sql.NullString

		if err := rows.Scan(&user.Name, &poolMode); err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		user.PoolMode = unwrapNullString(poolMode)
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
			&database.CurrentConnections)
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
			&pool.PoolMode)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to fetch row from results")
		}
		pools = append(pools, pool)
	}
	return pools, nil
}
