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
reporter();

