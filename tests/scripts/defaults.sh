#!/bin/bash
set -uo pipefail

source common.sh

title "Testing wih default settings"

# All healthcheck settings at default
start_healthcheck
start_pgbouncer

desc "PGBouncer is running, healthcheck should return OK"
test_ok_healthcheck
desc "Debug endpoints disabled by default, should return Error"
test_notok_debug
desc "No credentials set, so status endpoints should return Error"
test_notok_status

stop_pgbouncer
desc "PGBouncer is stopped, healthcheck should return Error"
test_notok_healthcheck
desc "PGbouncer is stopped, status endpints should return Error"
test_notok_status

start_pgbouncer
desc "PGBouncer is running again, healthcheck should return OK"
test_ok_healthcheck

end_tests
summary
