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


let totalAge = 0;
personnel.forEach(p => totalAge += p.age);
console.log("Average age (for)   :", totalAge / personnel.length);

console.log("Average age (reduce):", (personnel.reduce((sum, p) => sum += p.age, 0) / personnel.length));