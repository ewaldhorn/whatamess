console.log("Using Bun and JS to solve AOC2022 Day 1");

let carrying = [];

// first get the input file read
const filename = "input.txt";
const file = Bun.file(filename);
const text = await file.text();

// split the text into elements
const elements = text.split("\n");

// now parse the elements
var total = 0;

for (let el of elements) {
    if (el.length > 1) {
        total += Number.parseInt(el);
    } else {
        carrying.push(total);
        total = 0;
    }
}

// sort our array
carrying.sort((a, b) => b - a);

console.log(`The most calories carried is ${carrying[0]}.`);
console.log(`The top three total calories carried is ${carrying[0] + carrying[1] + carrying[2]}.`);