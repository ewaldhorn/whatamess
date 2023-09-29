// Shuffle an array of items randomly.

let myArray = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15];

const report = () => myArray.toString();
const reportFixed = () => myArray.map((e) => e.toString().padStart(3));

const shuffleArrayWithFor = (arr) => {
    for (let pos = arr.length - 1; pos > 0; pos--) {
        const j = Math.floor(Math.random() * (pos + 1));
        [arr[pos], arr[j]] = [arr[j], arr[pos]];
    }
};

const shuffleArrayWithSort = (arr) => arr.sort(() => Math.random() - 0.5); // don't care about parameters


const shuffleArrayWithMap = (arr) => arr.map((e) => ({ sort: Math.random(), value: e }))
    .sort((a, b) => a.sort - b.sort)
    .map((a) => a.value);

console.log("Starting with : " + report());
console.log();

console.log("Starting fixed: " + reportFixed());
shuffleArrayWithFor(myArray);
console.log("Shuffled to   : " + reportFixed());
shuffleArrayWithSort(myArray);
console.log("Shuffled to   : " + reportFixed());
myArray = shuffleArrayWithMap(myArray);
console.log("Shuffled to   : " + reportFixed());