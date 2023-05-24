console.log("JavaScript Arrays")

const inputArray = [9, 2, 3, 8, 1, 5, 4, 7, 6]

console.log("For Each:")
inputArray.forEach((val) => { console.log(val, "is even:", val % 2 == 0) })
console.log()

console.log("Map")
let evenArray = inputArray.map((val)=>{
    if (val %2==0) {
        return val
    }
})

console.log("Even values are", evenArray)