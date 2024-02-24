module ReverseString

let reverse (input: string): string = 
    let mutable ans = ""
    for i in input.Length .. -1 .. 0 do
        ans <- ans + input[i..i]
    ans
