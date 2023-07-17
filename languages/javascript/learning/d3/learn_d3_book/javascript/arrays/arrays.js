const colors = ["red", "blue", "green"];
const geocoords = [27.2345, 34.9937];
const numbers = [1, 2, 3, 4, 5, 6];
const empty = [];

for (let i = 0; i < colors.length; i++) {
    console.log(colors[i]);
}

for (let color of colors) {
    console.log(color);
}

colors.forEach(function (color, idx, colors) {
    console.log("(" + color + ") " + (idx + 1) + ": " + idx + " [" + colors + "]");
});