# 0.11.x
# zig build-lib src/main.zig -target wasm32-freestanding -dynamic -rdynamic -femit-bin=out/main.wasm
# 
# 0.12.x
zig build-exe src/main.zig -target wasm32-freestanding -fno-entry --export=_draw --export=allocUint8 --export=free -femit-bin=out/main.wasm
