#!/bin/bash
set -uo pipefail

source common.sh

# Enable enhanced check
export ENHANCED_CHECK=1
# Enable datadog agent check
export CHECK_DDAGENT=1

init_creds
make_dd_stub_pass
start_healthcheck
start_pgbouncer

test_ok_healthcheck
test_notok_debug
test_ok_status

stop_pgbouncer
test_notok_healthcheck
test_notok_status

make_dd_stub_fail
test_notok_healthcheck

start_pgbouncer
test_notok_healthcheck

end_tests
summary
