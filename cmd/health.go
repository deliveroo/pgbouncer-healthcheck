package main

import (
	"context"
	"database/sql"
	"os/exec"

	"github.com/pkg/errors"
)

func checkHealth(ctx context.Context, db *sql.DB) error {
	_, err := getUsers(ctx, db)
	return errors.Wrap(err, "PGbouncer health check failed")
}

func getDmesg() ([]byte, error) {
	dmesgCmd := exec.Command("dmesg")
	output, err := dmesgCmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching dmesg data")
	}
	return output, nil
}

func getProcesses() ([]byte, error) {
	psCmd := exec.Command("ps", "-ef")
	output, err := psCmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching process data")
	}
	return output, nil
}
