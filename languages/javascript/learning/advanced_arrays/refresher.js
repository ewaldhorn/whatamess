/*
    Some basic JavaScript array refreshers.
 */
console.log("\n");


const items = ["cat", "dog"];
console.log("Current:", items, " (" + items.length + ")");
items.push("parrot");
console.log("Updated:", items, " (" + items.length + ")");

console.log("\nItems:");
for (let i = 0; i < items.length; i++) {
    console.log((i + 1), ":", items[i]);
}


console.log("\n");