package main

import "core:fmt"
import "core:os"

main :: proc() {
    fmt.println("Number of command-line arguments:", len(os.args))
    fmt.println("Value of first-command-line argument:", os.args[0])

    for a in os.args[1:] {
        fmt.println(a)
    }
}
