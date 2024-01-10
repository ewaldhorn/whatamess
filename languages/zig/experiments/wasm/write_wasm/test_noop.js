const fs = require('fs');
const source = fs.readFileSync("./noop.wasm");
const typedArray = new Uint8Array(source);

WebAssembly.instantiate(typedArray, {}).then(result => {
    console.log("It loaded!");
});