package main

import "core:fmt"
import "core:os"
import "core:strconv"

main :: proc() {
    sa := "foo"
    sb := "13"

    a := parse_int_or_fail(sa)
    b := parse_int_or_fail(sb)

    fmt.println(a + b)
}

parse_int_or_fail :: proc(s: string) -> int {
    n, ok := strconv.parse_int(s)
    if !ok {
        fmt.println("not a valid integer:", s)
        os.exit(-1)
    }

    return n
}
