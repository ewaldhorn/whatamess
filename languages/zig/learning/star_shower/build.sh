# with 0.11.x
# exports are automatic
# zig build-lib src/main.zig -target wasm32-freestanding -dynamic -rdynamic -femit-bin=out/main.wasm
# 
# with 0.12.x
# seems I need to specify exports, need to create a proper build.zig file at some stage
zig build-exe src/main.zig -target wasm32-freestanding -fno-entry --export=_draw --export=allocUint8 --export=free -femit-bin=out/main.wasm
