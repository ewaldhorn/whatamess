let mem;

const importObject = {
  dom_env: {
    dom_now: () => performance.now(),
  },
  app_env: {
    app_init: (w, h) => {
      const cv = document.getElementById("app").appendChild(document.createElement("canvas"));
      cv.width = w;
      cv.height = h;
      window._ctx = cv.getContext("2d");
      window._ctx.fillStyle="#339933";
      window._ctx.fillRect(0,0,w,h);
    },
    update_update: (a) => {
      console.log("Got back:", a);
    },
  },
};

WebAssembly.instantiateStreaming(fetch("application.wasm"), importObject).then((wasm) => {
  mem = wasm.instance.exports.memory;
  wasm.instance.exports.bootup();   // JS -> Odin
});
