#!/usr/bin/env bash

isAskingQuestion() {
  [[ $1 =~ \?\ *$ ]]
}

isAllUppercase() {
  [[ $1 =~ ^[^[:lower:]]+$ && $1 =~ [[:upper:]] ]]
}

main() {
    # first check if it's empty, or a bunch of whitespace 
    if [[ $1 == "" || $1 =~ ^[[:space:]]*$ ]]; then
        echo "Fine. Be that way!"
    # now check for uppercase and a question mark
    elif isAskingQuestion "$1" && isAllUppercase "$1"; then
        echo "Calm down, I know what I'm doing!"
    elif isAskingQuestion "$1"; then
        echo "Sure."
    elif isAllUppercase "$1"; then
        echo "Whoa, chill out!"
    else
        echo "Whatever."
    fi
}

main "$*" || echo "Fine. Be that way!"
