module TwoFer

let twoFer (input: string option): string = 
    match input with
    | Some(input) -> "One for " + input + ", one for me." 
    | None -> "One for you, one for me."
