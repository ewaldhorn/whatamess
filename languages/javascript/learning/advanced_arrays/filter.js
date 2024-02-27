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

const salesReps = personnel.filter(p => p.department === "SALES");
console.log("Sales Reps:", salesReps)

const firstHRPerson = personnel.find(p => p.department === "HR");
console.log("First HR:", firstHRPerson);

const allSalesPeople = personnel.every(p => p.department === "SALES");
const everyoneOver21 = personnel.every(p => p.age >= 21);
const someSafetyPeople = personnel.some(p => p.department === "SAFETY");

console.log("Is everyone in sales?:", allSalesPeople);
console.log("Is everyone over 21?", everyoneOver21);
console.log("Safety reps here?", someSafetyPeople);