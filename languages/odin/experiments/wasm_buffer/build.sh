#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"

FLAGS=(-target:js_wasm32 -o:size -no-entry-point)

odin build ./src -out:./web/application.wasm "${FLAGS[@]}"
