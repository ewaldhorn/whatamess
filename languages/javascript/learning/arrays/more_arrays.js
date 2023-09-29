console.log("Some JavaScript Array Refreshers.");

console.log();
console.log("Simple array creation and element access:");
let myArray = [12, "Dylan", 34, "Sally"];
let myFancyArray = Array(12, "Dylan", 34, "Sally");

const reporter = () => {
    console.log("myArray now contains: " + myArray);
    console.log("myFancyArray now contains: " + myFancyArray);
};

const newHeader = (msg) => {
    console.log();
    console.log(msg);
};

console.log(myArray[0] + " : " + myFancyArray[0]);
console.log(myArray[1] + " : " + myFancyArray[1]);
console.log(myArray[2] + " : " + myFancyArray[2]);
console.log(myArray[3] + " : " + myFancyArray[3]);

console.log("myArray contains " + myArray.length + " items.");
console.log("myFancyArray also contains " + myFancyArray.length + " items.");

console.log("myArray contains: " + myArray);
console.log("myFancyArray contains: " + myFancyArray);

newHeader("Adding contents to arrays:");
myArray.push(56, "Bob", 78, "Ted"); // append to the end
myFancyArray.unshift("Kate"); // insert in the front
reporter();

newHeader("Removing elements from an array:");
myArray.shift(); // remove first element
myFancyArray.pop(); // remove last element
reporter();

newHeader("Splicing and dicing:");
myArray.splice(3, 2); // remove two elements (at positions 3 and 4) starting at position 3
myFancyArray.splice(2, 0, 'Derrick', 19); // add two elements, delete 0 elements
reporter();

newHeader("Check if it's an array:");
console.log("myArray is an array: " + Array.isArray(myArray));
console.log("reporter is an array: " + Array.isArray(reporter));

newHeader("Loops");
console.log("For loop:");
for (let i = 0; i < myArray.length; i++) {
    console.log(myArray[i]);
}
newHeader("While loop:");
let i = 0;
while (i < myArray.length) {
    console.log(myArray[i]);
    i++;
}
newHeader("For in loop:");
for (let i in myArray) {
    console.log(myArray[i]);
}
newHeader("For of loop:");
for (let element of myArray) {
    console.log(element);
}
newHeader("For each loop:");
myArray.forEach((e) => console.log(e));

newHeader("Joining up:");
console.log("Just .toString()   : " + myArray.toString());
console.log("Custom join instead: " + myArray.join("|"));

newHeader("Comparisons:");
const tmpArraySame = ["Dylan", 34, "Sally", 78, "Ted"];
const tmpArrayDiff = ["Dylan", 34, "Sally", 88, "Ted"];
const tmpArrayShorter = ["Dylan", 34, "Sally", 78];
const tmpArrayDiffOrder = ["Dylan", 78, "Sally", 34, "Ted"];

let isEqual = (myArray.length === tmpArraySame.length) && myArray.every((el, idx) => el == tmpArraySame[idx]);
console.log("myArray === tmpArraySame: " + isEqual);

isEqual = (myArray.length === tmpArrayDiff.length) && myArray.every((el, idx) => el == tmpArrayDiff[idx]);
console.log("myArray === tmpArrayDiff: " + isEqual);

isEqual = (myArray.length === tmpArrayShorter.length) && myArray.every((el, idx) => el == tmpArrayShorter[idx]);
console.log("myArray === tmpArrayShorter: " + isEqual);

isEqual = (myArray.length === tmpArrayDiffOrder.length) && myArray.every((el) => tmpArrayDiffOrder.includes(el));
console.log("myArray === tmpArrayDiffOrder: " + isEqual);

newHeader("Copying:");
console.log("With slice, it's ok for primitives: " + myArray.slice());
const newArray = JSON.parse(JSON.stringify(myArray));
console.log("Deep copies are required for objects: " + newArray);
console.log();
console.log("Now modify the copy.");
newArray[0] = "Frank";
console.log("Original: " + myArray);
console.log("Copy    : " + newArray);

newHeader("Merging arrays:");
console.log("Merge myArray and myFancyArray: " + myArray.concat(myFancyArray));
console.log("Merge myArray and myFancyArray and myArray: " + myArray.concat(myFancyArray, myArray));
console.log();
console.log("Now with the spread (...) operator:");
console.log("Merge myArray and myFancyArray: " + [...myArray, ...myFancyArray]);
console.log("Merge myArray and myFancyArray and myArray: " + [...myArray, ...myFancyArray, ...myArray]);

newHeader("Searching:");