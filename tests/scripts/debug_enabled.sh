#!/bin/bash
set -uo pipefail

source common.sh

title "Testing with default settings plus Debug endpoints enabled"

export ENABLE_DEBUG_ENDPOINTS=1
start_healthcheck
start_pgbouncer

desc "PGBouncer is running, healthcheck should return OK"
test_ok_healthcheck
desc "Debug endpoints enabled, should return OK"
test_ok_debug
desc "No credentials set, so status endpoints should return Error"
test_notok_status

stop_pgbouncer
desc "PGBouncer is stopped, but debug endpoints should still work"
test_ok_debug
desc "PGBouncer is stopped, healthcheck should return Error"
test_notok_healthcheck

end_tests
summary
