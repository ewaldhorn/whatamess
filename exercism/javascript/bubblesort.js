// Bubblesort in JS

const myArray = [9, 0, 1, 4, 3, 2, 7, 5, 6, 8];
const my2ndArray = JSON.parse(JSON.stringify(myArray));

console.log("Current status 1: " + myArray);
console.log("Current status 2: " + my2ndArray);

const bubbleSort = (arr) => {
    let swapped;

    do {
        swapped = false;
        for (let i = 0; i < arr.length - 1; i++) {
            if (arr[i] > arr[i + 1]) {
                let temp = arr[i];
                arr[i] = arr[i + 1];
                arr[i + 1] = temp;
                swapped = true;
            }
        }
    } while (swapped);

    return arr;
};

const bubbleSortImproved = (arr) => {
    let swapped;

    do {
        swapped = false;
        for (let i = 0; i < arr.length - 1; i++) {
            if (arr[i] > arr[i + 1]) {
                [arr[i], arr[i + 1]] = [arr[i + 1], arr[i]];
                swapped = true;
            }
        }
    } while (swapped);

    return arr;
};

console.log();
console.log("Sorting.");
console.log();

bubbleSort(myArray);
bubbleSortImproved(my2ndArray);
console.log("Current status 1: " + myArray);
console.log("Current status 2: " + my2ndArray);

