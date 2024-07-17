(async function () {
  let dart2wasm_runtime;
  let moduleInstance;
  try {
    const dartModulePromise = WebAssembly.compileStreaming(fetch("main.wasm"));
    const imports = {};
    dart2wasm_runtime = await import("./main.mjs");
    moduleInstance = await dart2wasm_runtime.instantiate(
      dartModulePromise,
      imports,
    );
  } catch (exception) {
    console.error(`Failed to fetch and instantiate wasm module: ${exception}`);
    console.error("See https://dart.dev/web/wasm for more information.");
    document.getElementById("heading").innerText =
      "\n\n\n\n\n\n\nSorry, your browser does not support WASM GC!\n\n\n\nSee https://developer.chrome.com/blog/wasmgc\n\n\n\n\n";
    document.getElementById("output").style.display = "none";
  }

  if (moduleInstance) {
    try {
      await dart2wasm_runtime.invoke(moduleInstance);
    } catch (exception) {
      console.error(`Exception while invoking test: ${exception}`);
    }
  }
})();
