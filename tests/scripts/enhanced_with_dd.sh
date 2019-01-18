#!/bin/bash
set -uo pipefail

source common.sh

title "Testing wih enhanced check plus Datadog check enabled"

# Enable enhanced check
export ENHANCED_CHECK=1
# Enable datadog agent check
export CHECK_DDAGENT=1

init_creds
make_dd_stub_pass
start_healthcheck
start_pgbouncer

desc "PGbouncer running and DD Agent check passes, healthcheck should return OK"
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

make_dd_stub_fail
desc "PGBouncer is stopped, DD Agent Check fails, healthcheck should return Error"
test_notok_healthcheck

start_pgbouncer
desc "PGBouncer is started, DD Agent Check fails, healthcheck should return Error"
test_notok_healthcheck

make_dd_stub_pass

desc "PGbouncer running and DD Agent check passes, healthcheck should return OK"
test_ok_healthcheck

end_tests
summary
