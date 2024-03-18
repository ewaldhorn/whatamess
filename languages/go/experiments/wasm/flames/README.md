# Flames in the Browser via Go

## Building

Task is used for everything, just run `task` to get a list of options.

### Wasm_Exec

`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .` to copy the current wasm_exec.js to the current folder.





/** flames */
(function () {
    var scr;
    let pallette = [];
    let flamePixelBuffer = [];
    let flameHeight = 0.050;
    let flameResolution = 1;
    let screenBuffer;
    let backBuffer;
    function setup() {
        scr = dcl.setupScreen(800, 600, 0, 0, document.getElementById('raytube'));
        backBuffer = dcl.createBuffer(400, 160);
        scr.setBgColor('black');
        document.body.style.backgroundColor = 'black';

        pallette = dcl.palette.fire;
        // Initialize color buffer
        for (let y = 0; y < backBuffer.canvas.height; y += flameResolution) {
            let row = [];
            for (let x = 0; x < backBuffer.canvas.width; x += flameResolution) {
                row.push(y === backBuffer.canvas.height - 1 ? dcl.randomi(0, 255) : 0);
            }
            flamePixelBuffer.push(row);
        }
        // initialize screen buffer
        screenBuffer = new ImageData(backBuffer.canvas.width, backBuffer.canvas.height);
    }

    function draw() {
        randomizeLastRowOfPixels();
        calculateAveragePixelValues();
        drawFlamePixelBuffer();
        backBuffer.paint(screenBuffer);
        dcl.image(backBuffer.canvas, 0, 0, 320, 240, scr.width / 320);
        requestAnimationFrame(draw);
    }

    setup();
    draw();

    function drawFlamePixelBuffer() {
        // Draw everything on the screen (push to a imagedata buffer first
        // Drawing each pixel directly greatly reduces performance!);
        for (let y = 0; y < screenBuffer.height; y += flameResolution) {
            for (let x = 0; x < screenBuffer.width; x += flameResolution) {
                let fy = (y / flameResolution) % flamePixelBuffer.length;
                let fx = (x / flameResolution) % flamePixelBuffer[fy].length;
                let i = Math.floor(flamePixelBuffer[fy][fx]);
                let cf = pallette[i];
                if (!cf) {
                    continue;
                }
                putPixel(y, x, cf);
            }
        }
    }

    function randomizeLastRowOfPixels() {
        //Set last row to random
        for (let x = 0; x < flamePixelBuffer[flamePixelBuffer.length - 1].length; x++) {
            flamePixelBuffer[flamePixelBuffer.length - 1][x] = dcl.randomi(0, 255);
        }
    }

    function calculateAveragePixelValues() {
        //Calculate sum for adjecent pixles:
        // |0|1|0| <-- Pixel (1 = in sum, 0 = excluded)
        // |1|1|1| 
        // |0|1|0| 
        for (let y = 0; y < flamePixelBuffer.length - 1; y++) {
            for (let x = 0; x < flamePixelBuffer[y].length; x++) {
                let y1 = (y + 1) % flamePixelBuffer.length;
                let y2 = (y + 2) % flamePixelBuffer.length;
                let x1 = (x - 1 + flamePixelBuffer[y].length) % flamePixelBuffer[y].length;
                let x2 = x % flamePixelBuffer[y].length;
                let x3 = (x + 1 + flamePixelBuffer[y].length) % flamePixelBuffer[y].length;
                let sum = ((
                    flamePixelBuffer[y1][x1] +
                    flamePixelBuffer[y1][x2] +
                    flamePixelBuffer[y1][x3] +
                    flamePixelBuffer[y2][x2]
                    // this average controls how high the fire can rise
                )) / (4 + flameHeight);
                flamePixelBuffer[y][x] = sum;
            }
        }
    }



    function putPixel(y, x, cf) {
        let idx = 4 * (y * screenBuffer.width + x);

        screenBuffer.data[idx] = cf.r;
        screenBuffer.data[idx + 1] = cf.g;
        screenBuffer.data[idx + 2] = cf.b;
        screenBuffer.data[idx + 3] = 255;
    }
})();
