#!/usr/bin/env bash

main() {
    # first remove unwanted characters
    cleaned="$*"
    cleaned="${cleaned//_}"
    cleaned="${cleaned//"*"}"

    # convert dashes into spaces
    cleaned="${cleaned/-/" "}"

    # now make it all uppercase
    cleaned="${cleaned^^}"

    # create an array for the letters
    declare -a letterArray
    # now split on space or dash
    # and find the first character of each word
    set -- junk $cleaned
    shift
    for word; do
      letterArray+=${word:0:1}
    done
    echo $letterArray
}

main "$*"

