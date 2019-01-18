#!/bin/bash

cd "$(dirname "$0")" || exit

source scripts/common.sh

run_test() {
    echo ""
    echo "==================================="
    echo "Running test: $1"
    echo "==================================="
    echo ""
    if docker run -it --rm -e VERBOSE -w="/tests" "$APPNAME" "/tests/$1"; then
        pass "-* $1 script passed *-"
    else
        fail "-* $1 script failed *-"
    fi
}

APPNAME=${1:-pgbouncer-healthcheck}
shift

export VERBOSE=${VERBOSE:-0}

if [[ $# -eq 0 ]]; then
    mapfile -t tests < <(find scripts -name '*.sh' ! -name 'common.sh' -printf '%f\n')
else
    tests=("$@")
fi

for test in "${tests[@]}"; do
    run_test "$test"
done

echo -e "${TC_LBLU}
==========================
== META TESTING SUMMARY ==
==========================
${TC_RST}"

echo -ne "${TC_BLD}${TC_GRN}Scripts Passed: $passed "
if (( failed >0 )); then
    echo -ne "${TC_RED}"
fi
echo -e "Failed: $failed${TC_RST}"

if [[ $failed -gt 0 ]]; then
    exit 1
fi
