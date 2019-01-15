#!/bin/bash
set -uo pipefail

source common.sh

# All healthcheck settings at default
start_healthcheck
start_pgbouncer

test_ok_healthcheck
test_notok_debug
test_notok_status

stop_pgbouncer
test_notok_healthcheck
test_notok_status

end_tests
summary
