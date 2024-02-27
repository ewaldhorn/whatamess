const personnel = [
    { name: "Ronel", surname: "Smith", age: 31, department: "SALES" },
    { name: "Willem", surname: "Brink", age: 33, department: "SALES" },
    { name: "Vee", surname: "Dale", age: 45, department: "PROCUREMENT" },
    { name: "Valerie", surname: "Simms", age: 48, department: "HR" },
    { name: "Astrid", surname: "Bool", age: 27, department: "HR" },
    { name: "Nic", surname: "Firth", age: 24, department: "MARKETING" },
    { name: "Nelly", surname: "Scruu", age: 55, department: "TRAINING" },
    { name: "Curtis", surname: "Limto", age: 48, department: "SAFETY" },
    { name: "Dale", surname: "Everrt", age: 23, department: "SALES" },
    { name: "Danielle", surname: "Witsche", age: 39, department: "SAFETY" }];

const numbers = [1, 4, 5, 6, 2, 3, 8, 9, 4, 3, 2, 4, 5, 6];
const sorted = [...numbers].sort((left, right) => left >= right); // copy because sort works in-place

console.log("Numbers (raw)   :", numbers);
console.log("Numbers (sorted):", sorted);

console.log(personnel.map((p) => `${p.name} ${p.surname}`));

console.log(personnel.map(x => x).sort((left, right) => left.surname >= right.surname).map((p) => `${p.name} ${p.surname}`));