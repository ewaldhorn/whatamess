#!/usr/bin/env bash

grains() {
    if [ "$1" -lt 1 ] || [ "$1" -gt 64 ]; then
        echo "Error: invalid input"
        exit 1
    else
        printf "%llu\n" $(( 2 ** ($1 - 1) )) # need to power to one less than the block, thanks wikipedia!
    fi
}

total() {
     printf "%llu\n" $(( 2 ** 64 - 1 )) # need an unsigned long and to subtract 1 because of overflow
}

main() {
    if [[ $1 == "total" ]]; then
        total
    else
        grains "$1"
    fi
}

# remember to call main
main "$*" || exit 1
