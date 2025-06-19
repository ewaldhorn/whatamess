package main

import "core:fmt"
import "core:os"
import "core:strconv"

main :: proc() {
    sa := "foo"
    sb := "13"

    a, ok := strconv.parse_int(sa)
    if !ok {
        fmt.println("not a valid integer:", sa)
        os.exit(-1)
    }

    b, ok := strconv.parse_int(sb)
    if !ok {
        fmt.println("not a valid ingeter:", sb)
        os.exit(-1)
    }

    fmt.println(a + b)
}
