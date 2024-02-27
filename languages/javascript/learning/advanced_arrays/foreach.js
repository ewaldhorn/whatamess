console.log("\n");


const items = ["cat", "dog"];
items.push("parrot");

console.log("\nItems:");
items.forEach((item, idx) => { console.log((idx + 1), ":", item) });


const numbers = [1, 5, 10, 30];
const triples = [];

numbers.forEach(n => triples.push(n * 3));

console.log("Numbers:", numbers);
console.log("Triples:", triples);

console.log("\n");