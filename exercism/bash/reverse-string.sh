#!/usr/bin/env bash

# read the string into a variable
theString="$*"

# determine the length, set up variables
length="${#theString}"   
i=0   
result=""

# loop through string, appended characters in reverse
while (( i<length )); do 
    result="${theString:i++:1}$result"
done

# print result
echo "$result"
