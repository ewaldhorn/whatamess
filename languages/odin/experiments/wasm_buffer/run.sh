#!/bin/bash

set -euo pipefail

echo "=> Building..."
./build.sh

echo "=> Dev Server Starting on http://localhost:9000"
http-server ./web/ -p 9000 -c-1
