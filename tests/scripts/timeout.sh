#!/bin/bash
set -uo pipefail

source common.sh

title "Testing wih default settings for working timeout on commands "

# Enable datadog agent check
export CHECK_DDAGENT=1

make_dd_stub_pass

start_healthcheck
start_pgbouncer

desc "PGbouncer running and DD Agent check passes, healthcheck should return OK"
test_ok_healthcheck

make_dd_stub_hang

desc "PGbouncer running and DD Agent check will hang. Should timeout and fail healthcheck"
test_notok_healthcheck

make_dd_stub_pass

desc "PGbouncer running and DD Agent check passes, healthcheck should return OK"
test_ok_healthcheck

end_tests
summary
