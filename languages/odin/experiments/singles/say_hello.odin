package main

import "core:fmt"
import "core:os"

main :: proc() {
    for name in os.args[1:] {
        fmt.printfln("Hi there, %s!", name)
    }
}
