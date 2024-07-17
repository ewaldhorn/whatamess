let buildArgsList;

// `modulePromise` is a promise to the `WebAssembly.module` object to be
//   instantiated.
// `importObjectPromise` is a promise to an object that contains any additional
//   imports needed by the module that aren't provided by the standard runtime.
//   The fields on this object will be merged into the importObject with which
//   the module will be instantiated.
// This function returns a promise to the instantiated module.
export const instantiate = async (modulePromise, importObjectPromise) => {
    let dartInstance;

    function stringFromDartString(string) {
        const totalLength = dartInstance.exports.$stringLength(string);
        let result = '';
        let index = 0;
        while (index < totalLength) {
          let chunkLength = Math.min(totalLength - index, 0xFFFF);
          const array = new Array(chunkLength);
          for (let i = 0; i < chunkLength; i++) {
              array[i] = dartInstance.exports.$stringRead(string, index++);
          }
          result += String.fromCharCode(...array);
        }
        return result;
    }

    function stringToDartString(string) {
        const length = string.length;
        let range = 0;
        for (let i = 0; i < length; i++) {
            range |= string.codePointAt(i);
        }
        if (range < 256) {
            const dartString = dartInstance.exports.$stringAllocate1(length);
            for (let i = 0; i < length; i++) {
                dartInstance.exports.$stringWrite1(dartString, i, string.codePointAt(i));
            }
            return dartString;
        } else {
            const dartString = dartInstance.exports.$stringAllocate2(length);
            for (let i = 0; i < length; i++) {
                dartInstance.exports.$stringWrite2(dartString, i, string.charCodeAt(i));
            }
            return dartString;
        }
    }

    // Prints to the console
    function printToConsole(value) {
      if (typeof dartPrint == "function") {
        dartPrint(value);
        return;
      }
      if (typeof console == "object" && typeof console.log != "undefined") {
        console.log(value);
        return;
      }
      if (typeof print == "function") {
        print(value);
        return;
      }

      throw "Unable to print message: " + js;
    }

    // Converts a Dart List to a JS array. Any Dart objects will be converted, but
    // this will be cheap for JSValues.
    function arrayFromDartList(constructor, list) {
        const length = dartInstance.exports.$listLength(list);
        const array = new constructor(length);
        for (let i = 0; i < length; i++) {
            array[i] = dartInstance.exports.$listRead(list, i);
        }
        return array;
    }

    buildArgsList = function(list) {
        const dartList = dartInstance.exports.$makeStringList();
        for (let i = 0; i < list.length; i++) {
            dartInstance.exports.$listAdd(dartList, stringToDartString(list[i]));
        }
        return dartList;
    }

    // A special symbol attached to functions that wrap Dart functions.
    const jsWrappedDartFunctionSymbol = Symbol("JSWrappedDartFunction");

    function finalizeWrapper(dartFunction, wrapped) {
        wrapped.dartFunction = dartFunction;
        wrapped[jsWrappedDartFunctionSymbol] = true;
        return wrapped;
    }

    // Imports
    const dart2wasm = {

_51: (x0,x1) => x0.createElement(x1),
_66: (x0,x1) => x0.getContext(x1),
_75: (x0,x1,x2,x3,x4) => x0.fillRect(x1,x2,x3,x4),
_76: f => finalizeWrapper(f,x0 => dartInstance.exports._76(f,x0)),
_77: (x0,x1,x2) => x0.addEventListener(x1,x2),
_78: (x0,x1,x2,x3,x4) => x0.getImageData(x1,x2,x3,x4),
_79: (x0,x1) => x0.getElementById(x1),
_80: (x0,x1) => x0.getElementById(x1),
_81: x0 => x0.focus(),
_82: f => finalizeWrapper(f,x0 => dartInstance.exports._82(f,x0)),
_83: (x0,x1,x2,x3,x4) => x0.clearRect(x1,x2,x3,x4),
_84: (x0,x1,x2,x3,x4) => x0.fillRect(x1,x2,x3,x4),
_85: f => finalizeWrapper(f,x0 => dartInstance.exports._85(f,x0)),
_86: (x0,x1) => x0.requestAnimationFrame(x1),
_87: f => finalizeWrapper(f,x0 => dartInstance.exports._87(f,x0)),
_88: (x0,x1) => x0.requestAnimationFrame(x1),
_89: (x0,x1) => x0.getElementById(x1),
_90: (x0,x1) => x0.getElementById(x1),
_91: (x0,x1,x2) => x0.insertBefore(x1,x2),
_99: (x0,x1,x2,x3) => x0.fillText(x1,x2,x3),
_100: (x0,x1,x2,x3) => x0.strokeText(x1,x2,x3),
_110: (x0,x1,x2,x3,x4) => x0.createLinearGradient(x1,x2,x3,x4),
_111: (x0,x1,x2) => x0.addColorStop(x1,x2),
_112: (x0,x1,x2) => x0.addColorStop(x1,x2),
_113: (x0,x1,x2) => x0.addColorStop(x1,x2),
_114: (x0,x1,x2) => x0.addColorStop(x1,x2),
_115: (x0,x1,x2) => x0.addColorStop(x1,x2),
_116: (x0,x1,x2) => x0.addColorStop(x1,x2),
_117: (x0,x1,x2) => x0.addColorStop(x1,x2),
_118: (x0,x1,x2) => x0.addColorStop(x1,x2),
_119: (x0,x1,x2) => x0.addColorStop(x1,x2),
_1251: x0 => x0.value,
_1539: (x0,x1) => x0.width = x1,
_1540: x0 => x0.width,
_1541: (x0,x1) => x0.height = x1,
_1542: x0 => x0.height,
_1612: (x0,x1) => x0.strokeStyle = x1,
_1614: (x0,x1) => x0.fillStyle = x1,
_1626: (x0,x1) => x0.lineWidth = x1,
_1636: (x0,x1) => x0.font = x1,
_1638: (x0,x1) => x0.textAlign = x1,
_1640: (x0,x1) => x0.textBaseline = x1,
_1674: x0 => x0.width,
_1675: x0 => x0.height,
_1676: x0 => x0.data,
_1918: () => globalThis.window,
_6889: (x0,x1) => x0.textContent = x1,
_6894: () => globalThis.document,
_6984: x0 => x0.body,
_7344: x0 => x0.clientWidth,
_7345: x0 => x0.clientHeight,
_7592: x0 => x0.offsetX,
_7593: x0 => x0.offsetY,
_12763: v => stringToDartString(v.toString()),
_12778: () => {
          let stackString = new Error().stack.toString();
          let frames = stackString.split('\n');
          let drop = 2;
          if (frames[0] === 'Error') {
              drop += 1;
          }
          return frames.slice(drop).join('\n');
        },
_12787: s => stringToDartString(JSON.stringify(stringFromDartString(s))),
_12788: s => printToConsole(stringFromDartString(s)),
_12806: (c) =>
              queueMicrotask(() => dartInstance.exports.$invokeCallback(c)),
_12808: (a, i) => a.push(i),
_12819: a => a.length,
_12821: (a, i) => a[i],
_12822: (a, i, v) => a[i] = v,
_12824: a => a.join(''),
_12827: (s, t) => s.split(t),
_12830: s => s.trim(),
_12834: (s, p, i) => s.indexOf(p, i),
_12837: (o, start, length) => new Uint8Array(o.buffer, o.byteOffset + start, length),
_12838: (o, start, length) => new Int8Array(o.buffer, o.byteOffset + start, length),
_12839: (o, start, length) => new Uint8ClampedArray(o.buffer, o.byteOffset + start, length),
_12840: (o, start, length) => new Uint16Array(o.buffer, o.byteOffset + start, length),
_12841: (o, start, length) => new Int16Array(o.buffer, o.byteOffset + start, length),
_12842: (o, start, length) => new Uint32Array(o.buffer, o.byteOffset + start, length),
_12843: (o, start, length) => new Int32Array(o.buffer, o.byteOffset + start, length),
_12846: (o, start, length) => new Float32Array(o.buffer, o.byteOffset + start, length),
_12847: (o, start, length) => new Float64Array(o.buffer, o.byteOffset + start, length),
_12851: (o) => new DataView(o.buffer, o.byteOffset, o.byteLength),
_12855: Function.prototype.call.bind(Object.getOwnPropertyDescriptor(DataView.prototype, 'byteLength').get),
_12856: (b, o) => new DataView(b, o),
_12858: Function.prototype.call.bind(DataView.prototype.getUint8),
_12860: Function.prototype.call.bind(DataView.prototype.getInt8),
_12862: Function.prototype.call.bind(DataView.prototype.getUint16),
_12864: Function.prototype.call.bind(DataView.prototype.getInt16),
_12866: Function.prototype.call.bind(DataView.prototype.getUint32),
_12868: Function.prototype.call.bind(DataView.prototype.getInt32),
_12874: Function.prototype.call.bind(DataView.prototype.getFloat32),
_12876: Function.prototype.call.bind(DataView.prototype.getFloat64),
_12897: o => o === undefined,
_12898: o => typeof o === 'boolean',
_12899: o => typeof o === 'number',
_12901: o => typeof o === 'string',
_12904: o => o instanceof Int8Array,
_12905: o => o instanceof Uint8Array,
_12906: o => o instanceof Uint8ClampedArray,
_12907: o => o instanceof Int16Array,
_12908: o => o instanceof Uint16Array,
_12909: o => o instanceof Int32Array,
_12910: o => o instanceof Uint32Array,
_12911: o => o instanceof Float32Array,
_12912: o => o instanceof Float64Array,
_12913: o => o instanceof ArrayBuffer,
_12914: o => o instanceof DataView,
_12915: o => o instanceof Array,
_12916: o => typeof o === 'function' && o[jsWrappedDartFunctionSymbol] === true,
_12920: (l, r) => l === r,
_12921: o => o,
_12922: o => o,
_12923: o => o,
_12924: b => !!b,
_12925: o => o.length,
_12928: (o, i) => o[i],
_12929: f => f.dartFunction,
_12930: l => arrayFromDartList(Int8Array, l),
_12931: l => arrayFromDartList(Uint8Array, l),
_12932: l => arrayFromDartList(Uint8ClampedArray, l),
_12933: l => arrayFromDartList(Int16Array, l),
_12934: l => arrayFromDartList(Uint16Array, l),
_12935: l => arrayFromDartList(Int32Array, l),
_12936: l => arrayFromDartList(Uint32Array, l),
_12937: l => arrayFromDartList(Float32Array, l),
_12938: l => arrayFromDartList(Float64Array, l),
_12939: (data, length) => {
          const view = new DataView(new ArrayBuffer(length));
          for (let i = 0; i < length; i++) {
              view.setUint8(i, dartInstance.exports.$byteDataGetUint8(data, i));
          }
          return view;
        },
_12940: l => arrayFromDartList(Array, l),
_12941: stringFromDartString,
_12942: stringToDartString,
_12945: l => new Array(l),
_12949: (o, p) => o[p],
_12953: o => String(o)
    };

    const baseImports = {
        dart2wasm: dart2wasm,


        Math: Math,
        Date: Date,
        Object: Object,
        Array: Array,
        Reflect: Reflect,
    };

    const jsStringPolyfill = {
        "charCodeAt": (s, i) => s.charCodeAt(i),
        "compare": (s1, s2) => {
            if (s1 < s2) return -1;
            if (s1 > s2) return 1;
            return 0;
        },
        "concat": (s1, s2) => s1 + s2,
        "equals": (s1, s2) => s1 === s2,
        "fromCharCode": (i) => String.fromCharCode(i),
        "length": (s) => s.length,
        "substring": (s, a, b) => s.substring(a, b),
    };

    dartInstance = await WebAssembly.instantiate(await modulePromise, {
        ...baseImports,
        ...(await importObjectPromise),
        "wasm:js-string": jsStringPolyfill,
    });

    return dartInstance;
}

// Call the main function for the instantiated module
// `moduleInstance` is the instantiated dart2wasm module
// `args` are any arguments that should be passed into the main function.
export const invoke = (moduleInstance, ...args) => {
    const dartMain = moduleInstance.exports.$getMain();
    const dartArgs = buildArgsList(args);
    moduleInstance.exports.$invokeMain(dartMain, dartArgs);
}

