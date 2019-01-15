#!/bin/bash
set -uo pipefail

source common.sh

# Enabled enhanced check
export ENHANCED_CHECK=1
init_creds
start_healthcheck
start_pgbouncer

test_ok_healthcheck
test_notok_debug
test_ok_status

stop_pgbouncer
test_notok_healthcheck
test_notok_status

end_tests
summary
