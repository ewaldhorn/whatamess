#!/usr/bin/env bash

# first check that we received a valid year, if any
# need to use $* because $1 only check first argument, there could be too many
if ! [[ "$*" =~ ^[0-9]+$ ]]; then
    echo "Usage: leap.sh <year>"
    exit 1
fi

# ok, now we have a year, lets do the divisible checks
if [[ ( $(("$1" % 4)) -eq 0 && $(("$1" % 100)) -ne 0) || ( $(("$1" % 400)) -eq 0 ) ]]; then
    echo true
else
    echo false
fi
