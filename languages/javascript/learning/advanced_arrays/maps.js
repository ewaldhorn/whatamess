const list = [1, 2, false, "data", { entry: "brown", count: 11 }, 4];

const newList = list.map(entry => "yolo");

console.log("List:", list);
console.log("NewList:", newList);

const numbers = [1, 5, 10, 30, 100];
const triples = numbers.map(n => n * 3);

console.log("Numbers:", numbers);
console.log("Triples:", triples);

const personnel = [{ name: "Larry" }, { name: "Sandra" }, { name: "Sal" }, { name: "Rita" }];

const personnelArea = personnel.map(p => { return { name: p.name, area: "Boston" } });

console.log("Personnel:", personnel);
console.log("WithAreas:", personnelArea);