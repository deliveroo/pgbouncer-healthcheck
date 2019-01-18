#!/bin/bash

if [[ VERBOSE -gt 1 ]]; then
    set -x
fi

if [[ -t 2 ]]; then
    # Output is a terminal
    TC_RED="\e[00;31m"
    TC_GRN="\e[00;32m"
    TC_CYN="\e[00;36m"
    TC_YEL="\e[01;33m"
    TC_LBLU="\e[01;34m"
    TC_LPUR="\e[01;35m"

    TC_BLD="\e[1m"
    TC_UND="\e[4m"
    TC_RST="\e[0m"
else
    TC_RED=""
    TC_GRN=""
    TC_CYN=""
    TC_YEL=""
    TC_LBLU=""
    TC_LPUR=""
    TC_BLD=""
    TC_UND=""
    TC_RST=""
fi

# Squelch Shellcheck Warning
export TC_LBLU

passed=0
failed=0

FAIL_HDR="${TC_RED}FAIL: ${TC_RST}"
PASS_HDR="${TC_GRN}PASS: ${TC_RST}"
INFO_HDR="${TC_CYN}INFO: ${TC_RST}"
FATL_HDR="${TC_RED}FATAL: ${TC_RST}"
TITLE_HDR="${TC_UND}${TC_LPUR}**: "
TITLE_END=" :**${TC_RST}"
DESC_HDR="${TC_YEL}>>>> ${TC_RST}"

if [[ $VERBOSE -gt 0 ]]; then
   curl_opts=("--silent" "--output" "/dev/stderr" "--write-out" "%{http_code}")
else
   curl_opts=("--silent" "--output" "/dev/null" "--write-out" "%{http_code}")
fi

# Location to put fake datadog-agent binary
DDAGENT_FAKE=/usr/bin/datadog-agent

_curl() {
    url=$1
    if STATUSCODE=$(curl "${curl_opts[@]}" "$url"); then
        if [[ $STATUSCODE -gt 400 ]]; then
            info "Request returned $STATUSCODE"
            return 1
        fi
    fi
}

init_creds() {
    export PGUSER="dd-agent"
    export PGPASSWORD="3c5c04a7-c21a-480f-bc42-30a1e0550fbc"
}

start_healthcheck() {
    info "Starting Healthcheck process"
    "${GOPATH}/src/github.com/deliveroo/pgbouncer-healthcheck/pgbouncer-healthcheck" &
    health_pid=$!
}

start_pgbouncer() {
    info "Starting PGBouncer"
    pgbouncer -u pgbouncer /etc/pgbouncer/pgbouncer.ini >&/dev/null &
    pgbouncer_pid=$!
    # Needs some time become ready
    sleep 1
}

stop_pgbouncer() {
    info "Stoppping PGBouncer (PID $pgbouncer_pid)"
    kill -INT "$pgbouncer_pid"
    retry=5
    while kill -0 "$pgbouncer_pid" 2>/dev/null && ((retry--)); do
        info "Waiting for PGBouncer to stop"
        sleep 1
    done
    if kill -0 "$pgbouncer_pid" 2>/dev/null; then
        info "Sending SIGTERM to PGBouncer"
        kill -TERM "$pgbouncer_pid"
    fi
}

end_tests() {
    kill -0 $health_pid 2>/dev/null || fail "Health-check process ended early"
}

fail() {
    echo -e "$FAIL_HDR$*">&2
    (( failed++ ))
}

pass() {
    echo -e "$PASS_HDR$*">&2
    (( passed++ ))
}

info() {
    echo -e "$INFO_HDR$*">&2
}

fatal() {
    echo -e "$FATL_HDR$*">&2
    exit 1
}

title() {
    echo -e "$TITLE_HDR$*$TITLE_END\n">&2
}

desc() {
    echo -e "\n\n$DESC_HDR$*">&2
}

test_ok_endpoint() {
    path=$1
    name=$2
    info "Testing for sucessful $name"
    if _curl "localhost:8000$path"; then
        pass "$name returned success"
    else
        fail "$name return error"
    fi
}

test_notok_endpoint() {
    path=$1
    name=$2
    info "Testing for error $name"
    if _curl "localhost:8000$path"; then
        fail "$name returned success"
    else
        pass "$name return error"
    fi
}

test_ok_healthcheck() {
    test_ok_endpoint "/health" "Health-check"
}

test_notok_healthcheck() {
    test_notok_endpoint "/health" "Health-check"
}

test_ok_debug() {
    test_ok_endpoint "/debug/meminfo" "Debug Endpoint"
}

test_notok_debug() {
    test_notok_endpoint "/debug/meminfo" "Debug Endpoint"
}

test_ok_status() {
    test_ok_endpoint "/status/users" "Status Endpoint"
}

test_notok_status() {
    test_notok_endpoint "/status/users" "Status Endpoint"
}

make_dd_stub_hang() {
    printf '#!/bin/sh\n' >"$DDAGENT_FAKE"
    printf 'sleep 1000000\n' >"$DDAGENT_FAKE"
    chmod 755 "$DDAGENT_FAKE"
}

make_dd_stub_pass() {
    printf '#!/bin/true\n' >"$DDAGENT_FAKE"
    chmod 755 "$DDAGENT_FAKE"
}

make_dd_stub_fail() {
    printf '#!/bin/false\n' >"$DDAGENT_FAKE"
    chmod 755 "$DDAGENT_FAKE"
}

summary() {
    echo
    echo -e "${TC_BLD}== TESTING SUMMARY ==${TC_RST}"
    if (( failed >0 )); then
        echo -e "${TC_BLD}${TC_GRN}Passed: $passed ${TC_RED}Failed: $failed${TC_RST}"
    else
        echo -e "${TC_BLD}${TC_GRN}Passed: $passed Failed: $failed${TC_RST}"
    fi
    if [[ $failed -gt 0 ]]; then
        exit 1
    fi
}
