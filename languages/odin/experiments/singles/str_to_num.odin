package main

import "core:fmt"
import "core:strconv"

main :: proc() {
    sa := "12"
    sb := "47"

    a, _ := strconv.parse_int(sa)
    b, _ := strconv.parse_int(sb)

    fmt.println(a + b)
}
