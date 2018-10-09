package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"os/exec"

	"github.com/pkg/errors"
)

func checkHealth(ctx context.Context, db *sql.DB) error {
	_, err := getUsers(ctx, db)
	if err != nil {
		return errors.Wrap(err, "PGbouncer health check failed")
	}
	err = getAgentHealth()
	return err
}

func getAgentHealth() error {
	psCmd := exec.Command("sudo", "datadog-agent", "health")
	err := psCmd.Run()
	if err != nil {
		return errors.Wrap(err, "Datadog agent health command failed")
	}
	return nil
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
	psCmd := exec.Command("ps", "-eo", "user,pid,ppid,c,stime,tty,%cpu,%mem,vsz,rsz,cmd")
	output, err := psCmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching process data")
	}
	return output, nil
}

func getMeminfo() ([]byte, error) {
	info, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching memory data")
	}
	return info, nil
}

func getLogs() ([]byte, error) {
	psCmd := exec.Command("journalctl", "--reverse", "-b", "--no-pager", "-n", "50")
	output, err := psCmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "Error fetching logs")
	}
	return output, nil
}
