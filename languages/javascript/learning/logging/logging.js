// ======================================================================================= STANDARD
const person = "Bob", age = 29;
console.log("Name:", person, "Age:", age);

// ========================================================================================== TABLE
const users = [
    { name: "Chris", age: 25 },
    { name: "Dennis", age: 15 },
    { name: "Victor", age: 17 }
];

console.table(users);

// ========================================================================================== GROUP
// Start a new console group named "Group 1"
console.group("Group 1");

// Log messages inside "Group 1"
console.log("Message 1");
console.log("Message 2");

// End "Group 1"
console.groupEnd();

// Start a new collapsed console group named "Group 2"
console.groupCollapsed("Group 2");

// Log messages inside "Group 2"
console.log("Message 3");
console.log("Message 4");

// End "Group 2"
console.groupEnd();

// ========================================================================================= ASSERT
const x = 5;

// Check if the condition x === 10 is true, if not, log the error message
console.assert(x === 10, "x is not equal to 10");
console.assert(x === 5, "x is equal to 5, so this won't be logged");

// ========================================================================================== COUNT
function greet() {
    // Log and count the number of times "greet" is called
    console.count("greet");

    // Return a greeting message
    return "Hello!";
}

// Call greet() two times
greet();
greet();

// Reset the counter for "greet"
console.countReset("greet");

// Call greet() again
greet();

// ========================================================================================== TRACE
function foo() {
    // Call the bar function
    bar();
}

function bar() {
    // Log a trace of the call stack
    console.trace("Trace:");
}

// Call the foo function
foo();

// ==================================================================================== INTERACTIVE
const obj = { name: "Chris", age: 25 };

// Display an interactive listing of the properties of the object
console.dir(obj);

// ========================================================================================= TIMING
console.time("timer");

console.log("\nWorking...");
for (let i = 0; i < 250_000_000; i++) {
    for (let j = 0; j < 30; j++) {
        // Some time-consuming operation
    }
}

console.timeEnd("timer");

// ========================================================================================== CLEAR
console.clear();
console.log("All gone!");