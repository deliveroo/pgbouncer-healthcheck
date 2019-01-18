#!/bin/bash
set -uo pipefail

source common.sh

title "Testing wih enhanced check enabled"

# Enabled enhanced check
export ENHANCED_CHECK=1
init_creds
start_healthcheck
start_pgbouncer

desc "PGBouncer is running, healthcheck should return OK"
test_ok_healthcheck
desc "Debug endpoints disabled by default, should return Error"
test_notok_debug
desc "Status endpoints should return OK"
test_ok_status

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
