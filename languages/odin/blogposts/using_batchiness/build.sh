odin build src/ \
  -target:js_wasm32 \
  -out:web/example.wasm \
  -o:size \
  -no-entry-point
