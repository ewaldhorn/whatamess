console.log("Monads!");

const square = (num) => { return { result: num.result * num.result, logs: num.logs.concat(`Squared ${num.result} to get ${num.result * num.result}.`) } };
const addOne = (num) => { return { result: num.result + 1, logs: num.logs.concat(`Added 1 to ${num.result} to get ${num.result + 1}.`) } };

const wrapWithLogs = (num) => { return { result: num, logs: [] } };

const result = square(addOne(square(wrapWithLogs(2))));
console.log("Square adding one to square two gives 25: " + JSON.stringify(result));