const myArray = [1, 2, 3, 4, 5, 6, 7, 8, 9];
const control = 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9;
console.log("We have    : " + myArray);
console.log("This totals: " + control);

let forTotal = 0;
for (let i = 0; i < myArray.length; i++) {
    forTotal += myArray[i];
}
console.log("Sum the hard way: " + forTotal + " (" + (forTotal == control) + ")");

let forEachTotal = 0;
myArray.forEach((e) => forEachTotal += e);
console.log("Sum with forEach: " + forEachTotal + " (" + (forEachTotal == control) + ")");

let reduceTotal = myArray.reduce((a, v) => a + v, 0);
console.log("Sum with reduce : " + reduceTotal + " (" + (reduceTotal == control) + ")");