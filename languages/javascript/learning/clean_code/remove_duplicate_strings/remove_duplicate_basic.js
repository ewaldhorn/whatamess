function removeDuplicateStrings(array) {
    const outputArray = [];

    array.forEach((element, index) => {
        // check if the same item exists ahead of where we are
        // if it does, then don't add it to the output
        if (array.indexOf(element, index + 1) === -1) {
            outputArray.push(element);
        }
    });

    return outputArray.sort();
}

function removeDuplicateStringsSimpler(array) {
    const outputArray = [];

    array.forEach(element => {
        // Check if the same item exists already in the output array.
        // If it doesn't, then we can add it:
        if (outputArray.indexOf(element) === -1) {
            outputArray.push(element);
        }
    });

    return outputArray.sort();
}

const myArray = ["this", "is", "an", "array", "of", "strings", "and", "this", "array", "has", "a", "duplicate", "of", "some", "kind"];
console.log("Algorithm One")
console.log(removeDuplicateStrings(myArray));
console.log("Algorithm Two")
console.log(removeDuplicateStringsSimpler(myArray));