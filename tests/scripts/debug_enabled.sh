#!/bin/bash
set -uo pipefail

source common.sh

export ENABLE_DEBUG_ENDPOINTS=1
start_healthcheck
start_pgbouncer

test_ok_healthcheck
test_ok_debug
test_notok_status

stop_pgbouncer
test_ok_debug
test_notok_healthcheck

end_tests
summary
