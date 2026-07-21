#!/usr/bin/env bash
# Build Jasteroids and stage a static, host-anywhere dist/ folder
# (index.html, synth-worklet.js, teavm/ — no server/servlet bits).
set -euo pipefail

cd "$(dirname "$0")"

echo "==> Building with Maven..."
mvn -q clean package

WEBAPP_DIR="target/jasteroids-0.0.1"
DIST_DIR="dist"

if [ ! -d "$WEBAPP_DIR" ]; then
  echo "error: expected build output at $WEBAPP_DIR, but it doesn't exist" >&2
  exit 1
fi

echo "==> Staging $DIST_DIR..."
rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

cp "$WEBAPP_DIR/index.html" "$DIST_DIR/"
cp "$WEBAPP_DIR/synth-worklet.js" "$DIST_DIR/"
mkdir -p "$DIST_DIR/teavm"
cp "$WEBAPP_DIR/teavm/classes.wasm" "$DIST_DIR/teavm/"
cp "$WEBAPP_DIR/teavm/classes.wasm-runtime.js" "$DIST_DIR/teavm/"
# Deliberately not copied: classes.wasm.map / .teadbg / src/ (debug info + Java
# sources) — useful locally, but not something to publish on a public deploy.

echo "==> Done. Static site ready in $DIST_DIR/"
echo "    Serve it with e.g.: python3 -m http.server -d $DIST_DIR 8080"
