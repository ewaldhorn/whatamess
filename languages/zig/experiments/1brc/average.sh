#!/bin/bash

# Input file path
input_file="measurements.txt"

# Check if the file exists
if [ ! -f "$input_file" ]; then
    echo "Error: File not found: $input_file"
    exit 1
fi

# Use awk to calculate average temperature per city
awk -F';' '{
    city = $1
    temperature = $2
    sum[city] += temperature
    count[city]++
}
END {
    for (city in sum) {
        average = sum[city] / count[city]
        printf "City: %s, Average Temperature: %.2f\n", city, average
    }
}' "$input_file"
