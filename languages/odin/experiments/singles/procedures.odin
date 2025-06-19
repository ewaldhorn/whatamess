package main

import "core:fmt"

main :: proc() {
    a := value()
    b := value()
    c := value()

    fmt.printfln("the values are %d, %d and %d", a, b, c)
}

value :: proc() -> int {
    return 47
}
