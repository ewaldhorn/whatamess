package main

import "core:fmt"
import "core:os"

main :: proc() {
    fmt.println(len(os.args))
}
