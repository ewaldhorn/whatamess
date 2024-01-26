# create an "out" folder if it does not exist yet
mkdir -p out
# copy the web resources to the out folder
cp web/* out/
zig build-exe src/main.zig -target wasm32-freestanding -fno-entry --export=helloWasm -femit-bin=out/main.wasm