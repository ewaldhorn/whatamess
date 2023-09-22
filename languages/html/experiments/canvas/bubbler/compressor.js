// Import Terser so we can use it
const { minify } = require('terser');

// Import fs so we can read/write files
const fs = require('fs');

// Define the config for how Terser should minify the code
// This is set to how you currently have this web tool configured
const config = {
    compress: {
        dead_code: true,
        drop_console: true,
        drop_debugger: true,
        keep_classnames: false,
        keep_fargs: false,
        keep_fnames: false,
        keep_infinity: false
    },
    mangle: {
        eval: true,
        keep_classnames: false,
        keep_fnames: false,
        toplevel: true,
        safari10: false
    },
    module: false,
    sourceMap: false,
    output: {
        comments: false
    }
};

// Load in your code to minify
const code = fs.readFileSync('script.js', 'utf8');

// Minify the code with Terser
const minified = await minify(code, config);

// Save the code!
fs.writeFileSync('script.min.js', minified.code);