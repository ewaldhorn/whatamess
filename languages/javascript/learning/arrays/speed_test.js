// Specific to Bun, Node performance is different.

let bigFatArray = [];

// Populate our big array with random data.
const populateBigFatArray = (count) => {
    for (let i = 0; i < count; i++) {
        let tmp = { number: Math.random() * 100000, anotherNumber: Math.random() * 100000 };
        bigFatArray.push(tmp);
    }
    console.log("BigFatArray now contains " + bigFatArray.length + " elements.");
};

const newHeader = (msg) => {
    console.log();
    console.log(msg);
};

console.log("Rough and ready performance tests for various loop strategies.");
newHeader("Populate array");
populateBigFatArray(10000000);


newHeader("For loop:");
let startTime = performance.now();
for (let i = 0; i < bigFatArray.length; i++) {
    if (!bigFatArray[i]) {
        console.log("Unexpected.");
    }
}
let endTime = performance.now();
let elapsedTime = endTime - startTime;
console.log("Total time taken: " + elapsedTime);

newHeader("While loop:");
startTime = performance.now();
let i = 0;
while (i < bigFatArray.length) {
    if (!bigFatArray[i]) {
        console.log("Unexpected.");
    }

    i++;
}
endTime = performance.now();
elapsedTime = endTime - startTime;
console.log("Total time taken: " + elapsedTime);

newHeader("For in loop:");
startTime = performance.now();
for (let i in bigFatArray) {
    if (!bigFatArray[i]) {
        console.log("Unexpected.");
    }
}
endTime = performance.now();
elapsedTime = endTime - startTime;
console.log("Total time taken: " + elapsedTime);

newHeader("For of loop:");
startTime = performance.now();
for (let element of bigFatArray) {
    if (!element) {
        console.log("Unexpected.");
    }
}
endTime = performance.now();
elapsedTime = endTime - startTime;
console.log("Total time taken: " + elapsedTime);

newHeader("For each loop:");
startTime = performance.now();
bigFatArray.forEach((e) => { if (!e) { console.log("Unexpected.") } });
endTime = performance.now();
elapsedTime = endTime - startTime;
console.log("Total time taken: " + elapsedTime);